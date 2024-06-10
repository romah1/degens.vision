package mints

import (
	"context"
	"encoding/json"
	"net/http"
	"solana_project/services/application-core/pkg/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type HandleApiV1MintsInitialFeedResponse = []HandleApiV1MintsInitialFeedResponseItem

type HandleApiV1MintsInitialFeedResponseItem struct {
	Mint       InitialFeedMint       `json:"mint"`
	Collection InitialFeedCollection `json:"collection"`
}

type InitialFeedMint struct {
	TxSignature      string            `json:"tx_signature"`
	ProgramKey       string            `json:"program_key"`
	BlockTime        time.Time         `json:"block_time"`
	CollectionKey    string            `json:"collection_key"`
	AssetKey         string            `json:"asset_key"`
	AssetOwnerKey    string            `json:"asset_owner_key"`
	AssetName        string            `json:"asset_name"`
	AssetSymbol      string            `json:"asset_symbol"`
	AssetImage       string            `json:"asset_image"`
	AssetDescription string            `json:"asset_description"`
	AssetAttributes  []json.RawMessage `json:"asset_attributes"`
	MintPrice        float64           `json:"mint_price"`
}

type InitialFeedCollection struct {
	CollectionKey string `json:"collection_key"`
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	Image         string `json:"image"`
	Size          uint64 `json:"size"`
	Launchpad     string `json:"launchpad"`
}

func HandleApiV1MintsInitialFeed(
	ctx context.Context,
	logger *zap.SugaredLogger,
	pg *pgxpool.Pool,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := pg.Query(
			ctx,
			sql.SelectInitialMintsFeed,
		)
		if err != nil {
			logger.Errorf("failed to select initial mints feed; error=%s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		defer rows.Close()

		var response HandleApiV1MintsInitialFeedResponse
		for rows.Next() {
			var mint InitialFeedMint
			var collection InitialFeedCollection
			err = rows.Scan(
				&mint.TxSignature,
				&mint.ProgramKey,
				&mint.BlockTime,
				&mint.CollectionKey,
				&mint.AssetKey,
				&mint.AssetOwnerKey,
				&mint.AssetName,
				&mint.AssetSymbol,
				&mint.AssetImage,
				&mint.MintPrice,
				&mint.AssetDescription,
				&mint.AssetAttributes,
				&collection.CollectionKey,
				&collection.Name,
				&collection.Symbol,
				&collection.Image,
				&collection.Size,
				&collection.Launchpad,
			)
			if err != nil {
				logger.Errorf("failed to read initial feed row; error=%s", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
			}
			response = append(response, HandleApiV1MintsInitialFeedResponseItem{
				Mint:       mint,
				Collection: collection,
			})
		}

		c.JSON(http.StatusOK, response)
	}
}
