package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	clients "solana_project/libraries/clients/lmnft"
	"solana_project/libraries/queues/mint_tracker_mints"
	"solana_project/services/mints-tracker/pkg/collection_updater"
	"solana_project/services/mints-tracker/pkg/ipfs_client"
	"solana_project/services/mints-tracker/pkg/queues/mint_tracker_mints_processor"
	"solana_project/services/mints-tracker/pkg/transaction_parser"
	"time"

	"github.com/IBM/sarama"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Postgres struct {
		ConnString string `yaml:"conn_string"`
	} `yaml:"postgres"`

	Kafka struct {
		Address string `yaml:"address"`
	} `yaml:"kafka"`

	Amqp struct {
		Url string `yaml:"url"`
	} `yaml:"amqp"`

	Solana struct {
		Rpc string `yaml:"rpc"`
	} `yaml:"solana"`

	MintTrackerMintsConsumer struct {
		Threads int `yaml:"threads"`
	} `yaml:"mint_tracker_mints_consumer"`

	CollectionUpdater struct {
		Threads         int `yaml:"threads"`
		IntervalSeconds int `yaml:"interval_seconds"`
	} `yaml:"collection_updater"`
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

	producer, err := sarama.NewSyncProducer([]string{cfg.Kafka.Address}, nil)
	if err != nil {
		panic(err)
	}

	lmnftClient := clients.NewLmnftClient(&http.Client{})
	// } dependencies

	err = runMintTrackerMintsConsumers(
		ctx,
		cfg,
		logger,
		rpc.CommitmentFinalized,
		pg,
		producer)
	if err != nil {
		panic(err)
	}

	runCollectionUpdaters(ctx, cfg, logger, pg, lmnftClient)

	logger.Info("service started")

	forever := make(chan bool)
	<-forever
}

func runCollectionUpdaters(
	ctx context.Context,
	cfg *Config,
	logger *zap.SugaredLogger,
	pg *pgxpool.Pool,
	lmnftClient *clients.LmnftClient,
) {
	for i := 0; i < cfg.CollectionUpdater.Threads; i++ {
		rpcClient := rpc.New(cfg.Solana.Rpc)
		updater := collection_updater.NewCollectionUpdater(
			logger, rpcClient, pg, lmnftClient,
		)
		go func() {
			updater.Start(
				ctx,
				time.Duration(cfg.CollectionUpdater.IntervalSeconds)*time.Second,
			)
		}()
	}
}

func runMintTrackerMintsConsumers(
	ctx context.Context,
	cfg *Config,
	logger *zap.SugaredLogger,
	commitment rpc.CommitmentType,
	pg *pgxpool.Pool,
	kafkaProducer sarama.SyncProducer,
) error {

	for i := 0; i < cfg.MintTrackerMintsConsumer.Threads; i++ {
		rpcClient := rpc.New(cfg.Solana.Rpc)
		conn, err := amqp091.Dial(cfg.Amqp.Url)
		if err != nil {
			log.Fatalf("unable to open connect to RabbitMQ server. Error: %s", err)
		}

		ch, err := conn.Channel()
		if err != nil {
			panic(err)
		}
		ipfsClient := ipfs_client.NewIpfsClient([]*http.Client{{}})
		transactionParser := transaction_parser.NewTransactionParser(logger, rpcClient, commitment, ipfsClient)
		processor := mint_tracker_mints_processor.NewMintTrackerMintsProcessor(
			logger,
			rpc.CommitmentFinalized,
			pg,
			transactionParser,
			kafkaProducer,
		)
		consumer, err := mint_tracker_mints.NewMintTrackerMintsConsumer(logger, ch, processor)
		if err != nil {
			return err
		}
		go func() {
			defer conn.Close()
			defer ch.Close()
			consumer.Start(ctx)
		}()
	}
	return nil
}
