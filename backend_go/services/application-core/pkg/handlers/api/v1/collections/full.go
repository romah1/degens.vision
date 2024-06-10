package collections

import (
	"context"
	"net/http"
	"solana_project/services/application-core/pkg/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type HandleApiV1CollectionsFullResponse struct {
	Collection        HandleApiV1CollectionsFullResponseCollection          `json:"collection"`
	MintsDistribution []HandleApiV1CollectionsFullResponseMintsDistribution `json:"mints_distribution"`
}

type HandleApiV1CollectionsFullResponseCollection struct {
	CollectionKey       string                                             `json:"collection_key"`
	CandyMachine        *string                                            `json:"candy_machine"`
	MintsCount          int                                                `json:"mints_count"`
	Launchpad           string                                             `json:"launchpad"`
	Name                string                                             `json:"name"`
	Symbol              string                                             `json:"symbol"`
	Image               string                                             `json:"image"`
	Size                int                                                `json:"size"`
	MintPrice           float64                                            `json:"mint_price"`
	Date                time.Time                                          `json:"date"`
	AttributesAmount    int                                                `json:"attributes_amount"`
	UniqueMintersAmount int                                                `json:"unique_minters_amount"`
	Type                string                                             `json:"type"`
	Links               []HandleApiV1CollectionsFullResponseCollectionLink `json:"links"`
	MarketLinks         []HandleApiV1CollectionsFullResponseCollectionLink `json:"market_links"`
}

type HandleApiV1CollectionsFullResponseCollectionLink struct {
	Type string `json:"type"`
	Uri  string `json:"uri"`
}

type HandleApiV1CollectionsFullResponseMintsDistribution struct {
	MintsAmount   int `json:"mints_amount"`
	HoldersAmount int `json:"holders_amount"`
}

func HandleApiV1CollectionsFull(
	ctx context.Context,
	logger *zap.SugaredLogger,
	pg *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		collectionKey := c.Request.URL.Query().Get("collection_key")
		if collectionKey == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing collection_key param"})
			return
		}
		collection, err := selectCollection(ctx, pg, collectionKey)
		if err != nil {
			logger.Errorf("failed to select collection from db; error=%s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		links, err := selectCollectionLinks(ctx, pg, collectionKey)
		if err != nil {
			logger.Errorf("failed to select collection links from db; err=%s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		collection.Links = links
		collection.MarketLinks = prepareMarketLinks(collectionKey, collection.CandyMachine)

		distribution, err := selectMintsDistribution(ctx, pg, collectionKey)
		if err != nil {
			logger.Errorf("failed to select mints distribution from db; error=%s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		response := HandleApiV1CollectionsFullResponse{
			Collection:        *collection,
			MintsDistribution: distribution,
		}

		c.JSON(http.StatusOK, response)
	}
	return gin.HandlerFunc(fn)
}

func selectCollection(
	ctx context.Context,
	pg *pgxpool.Pool,
	collectionKey string,
) (*HandleApiV1CollectionsFullResponseCollection, error) {
	var collection HandleApiV1CollectionsFullResponseCollection
	row := pg.QueryRow(ctx, sql.SelectCollection, collectionKey)
	err := row.Scan(
		&collection.CollectionKey,
		&collection.CandyMachine,
		&collection.Name,
		&collection.Symbol,
		&collection.Image,
		&collection.Size,
		&collection.MintPrice,
		&collection.MintsCount,
		&collection.Launchpad,
		&collection.AttributesAmount,
		&collection.UniqueMintersAmount,
		&collection.Date,
		&collection.Type,
	)
	if err != nil {
		return nil, err
	}
	return &collection, nil
}

func selectCollectionLinks(
	ctx context.Context,
	pg *pgxpool.Pool,
	collectionKey string,
) ([]HandleApiV1CollectionsFullResponseCollectionLink, error) {
	rows, err := pg.Query(
		ctx,
		sql.SelectCollectionLinks,
		collectionKey,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var links []HandleApiV1CollectionsFullResponseCollectionLink
	for rows.Next() {
		var link HandleApiV1CollectionsFullResponseCollectionLink
		err = rows.Scan(&link.Type, &link.Uri)
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}

	return links, nil
}

func prepareMarketLinks(collectionKey string, candyMachine *string) []HandleApiV1CollectionsFullResponseCollectionLink {
	var res []HandleApiV1CollectionsFullResponseCollectionLink

	res = append(
		res,
		HandleApiV1CollectionsFullResponseCollectionLink{
			Type: "magiceden",
			Uri:  "https://magiceden.io/marketplace/" + collectionKey,
		},
	)

	if candyMachine != nil {
		res = append(
			res,
			HandleApiV1CollectionsFullResponseCollectionLink{
				Type: "tensor",
				Uri:  "https://tensor.trade/trade/" + *candyMachine,
			},
		)

		res = append(
			res,
			HandleApiV1CollectionsFullResponseCollectionLink{
				Type: "sniper",
				Uri:  "https://www.sniper.xyz/collection/" + *candyMachine,
			},
		)

		res = append(
			res,
			HandleApiV1CollectionsFullResponseCollectionLink{
				Type: "hyperspace",
				Uri:  "https://hyperspace.xyz/collection/" + *candyMachine,
			},
		)
	}

	return res
}

func selectMintsDistribution(
	ctx context.Context,
	pg *pgxpool.Pool,
	collectionKey string,
) ([]HandleApiV1CollectionsFullResponseMintsDistribution, error) {
	rows, err := pg.Query(ctx, sql.SelectMintsDistribution, collectionKey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	distribution := []HandleApiV1CollectionsFullResponseMintsDistribution{}
	for rows.Next() {
		var md HandleApiV1CollectionsFullResponseMintsDistribution
		err = rows.Scan(
			&md.MintsAmount,
			&md.HoldersAmount,
		)
		if err != nil {
			return nil, err
		}
		distribution = append(distribution, md)
	}
	return distribution, nil
}
