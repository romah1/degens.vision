package collections

import (
	"context"
	"net/http"
	"solana_project/services/application-core/pkg/sql"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type HandleTopCollectionsHoldersMintedRequest struct {
	CollectionKey string `json:"collection_key" binding:"required"`
}

type HandleTopCollectionsHoldersMintedResponse struct {
	Collections []HandleTopCollectionsHoldersMintedCollection `json:"collections"`
}

type HandleTopCollectionsHoldersMintedCollection struct {
	CollectionKey string `json:"collection_key"`
	MintsCount    int    `json:"mints_count"`
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	Image         string `json:"image"`
	Size          int    `json:"size"`
}

func HandleTopCollectionsHoldersMinted(
	ctx context.Context,
	logger *zap.SugaredLogger,
	pg *pgxpool.Pool,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request HandleTopCollectionsHoldersMintedRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}

		rows, err := pg.Query(
			ctx,
			sql.SelectTopCollectionsHoldersMinted,
			request.CollectionKey,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		defer rows.Close()

		collections := []HandleTopCollectionsHoldersMintedCollection{}
		for rows.Next() {
			var collection HandleTopCollectionsHoldersMintedCollection
			err = rows.Scan(
				&collection.CollectionKey,
				&collection.MintsCount,
				&collection.Name,
				&collection.Symbol,
				&collection.Image,
				&collection.Size,
			)
			if err != nil {
				logger.Errorf("failed to scan row while selecting collections from db; error=%s", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
			}
			collections = append(collections, collection)
		}
		response := HandleTopCollectionsHoldersMintedResponse{
			Collections: collections,
		}

		c.JSON(http.StatusOK, response)
	}
}
