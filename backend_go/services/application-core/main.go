package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"solana_project/libraries/queues/mint_tracker_mints"
	"solana_project/services/application-core/pkg/auth_middleware"
	"solana_project/services/application-core/pkg/handlers/admin/refcodes"
	"solana_project/services/application-core/pkg/handlers/api/v1/auth"
	"solana_project/services/application-core/pkg/handlers/api/v1/collections"
	"solana_project/services/application-core/pkg/handlers/api/v1/mints"
	"solana_project/services/application-core/pkg/handlers/api/v1/users"
	"solana_project/services/application-core/pkg/handlers/health"
	"solana_project/services/application-core/pkg/handlers/helius"
	"solana_project/services/application-core/pkg/handlers/ws/v1/live_feed"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Postgres struct {
		ConnString string `yaml:"conn_string"`
	} `yaml:"postgres"`
	Kafka struct {
		Address string `yaml:"address"`
	} `yaml:"kafka"`
	Amqp struct {
		Url string `yaml:"url"`
	} `yaml:"amqp"`
}

func NewConfig(path string) (*Config, error) {
	config := Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	configPath := os.Args[1]
	cfg, err := NewConfig(configPath)
	if err != nil {
		panic(err)
	}
	// dependencies {
	ctx := context.Background()
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer func(logger *zap.SugaredLogger) {
		err := logger.Sync()
		if err != nil {
			fmt.Println("Failed to sync logger; error=" + err.Error())
		}
	}(logger)

	pg, err := pgxpool.New(ctx, cfg.Postgres.ConnString)
	if err != nil {
		panic(err)
	}
	liveFeedHub := live_feed.NewLiveFeedHub()

	conn, err := amqp091.Dial(cfg.Amqp.Url)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	producer, err := mint_tracker_mints.NewMintTrackerMintsProducer(logger, ch)
	if err != nil {
		panic(err)
	}
	// } dependencies

	go runLiveFeedConsumer(ctx, cfg, logger, liveFeedHub)
	runServer(ctx, cfg, logger, pg, liveFeedHub, producer)
}

func runLiveFeedConsumer(
	ctx context.Context,
	cfg *Config,
	logger *zap.SugaredLogger,
	liveFeedHub *live_feed.LiveFeedHub,
) {
	sarama.Logger = log.Default()
	consumer, err := sarama.NewConsumer([]string{cfg.Kafka.Address}, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	partConsumer, err := consumer.ConsumePartition("live_feed", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partConsumer.Close()

	for {
		select {
		// Чтение сообщения из Kafka
		case msg, ok := <-partConsumer.Messages():
			if !ok {
				logger.Info("Channel closed, exiting goroutine")
				return
			}
			liveFeedHub.Broadcast(msg.Value)
			logger.Info("Broadcasted message")
		case <-ctx.Done():
			return
		}
	}
}

func runServer(
	ctx context.Context,
	cfg *Config,
	logger *zap.SugaredLogger,
	pg *pgxpool.Pool,
	hub *live_feed.LiveFeedHub,
	producer *mint_tracker_mints.MintTrackerMintsQueueProducer) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authMiddleware, err := auth_middleware.ProvideAuthMiddleware(ctx, logger, pg)
	if err != nil {
		logger.Fatal("Failed to provide auth middleware: " + err.Error())
	}
	err = authMiddleware.MiddlewareInit()
	if err != nil {
		logger.Fatal("Failed to init auth middleware: " + err.Error())
	}
	r.GET("/health", health.HandleHealth)
	r.POST("/auth/init", auth.HandleApiV1AuthInit(ctx, logger, pg))
	r.POST("/auth/verify", authMiddleware.LoginHandler)
	r.POST("/admin/refcodes/create", refcodes.HandleAdminRefcodesCreate(ctx, logger, pg))
	r.GET("/admin/refcodes/view", refcodes.HandleAdminRefcodesView(ctx, logger, pg))
	authRequiredRouter := r.Group("/api/v1")
	authRequiredRouter.Use(authMiddleware.MiddlewareFunc())

	{
		authRequiredRouter.GET("/auth/refresh_token", authMiddleware.RefreshHandler)
		authRequiredRouter.GET("/users/user/:address", users.HandleApiV1UsersUser(ctx, logger, pg))
		authRequiredRouter.POST("/users/update", users.HandleApiV1UpdateUser(ctx, logger, pg))
		authRequiredRouter.GET("/collections/top", collections.HandleApiV1CollectionsTop(ctx, logger, pg))
		authRequiredRouter.GET("/collections/top_holders", collections.HandleTopHolders(ctx, logger, pg))
		authRequiredRouter.GET("/collections/full", collections.HandleApiV1CollectionsFull(ctx, logger, pg))
		authRequiredRouter.GET("/collections/search", collections.HandleApiV1CollectionsSearch(ctx, logger, pg))
		authRequiredRouter.GET("/collections/top_collections_holders_minted", collections.HandleTopCollectionsHoldersMinted(ctx, logger, pg))
		authRequiredRouter.GET("/mints/initial_feed", mints.HandleApiV1MintsInitialFeed(ctx, logger, pg))
	}

	{
		ws := r.Group("/ws/v1")
		ws.Use(authMiddleware.MiddlewareFunc())
		liveFeedHandler := live_feed.NewLiveFeedHandler(hub)
		ws.GET("/live_feed", func(c *gin.Context) {
			liveFeedHandler.ServeHTTP(c.Writer, c.Request)
		})
	}

	{
		r.POST("/helius/mints", helius.HandleHeliusMints(ctx, producer))
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
