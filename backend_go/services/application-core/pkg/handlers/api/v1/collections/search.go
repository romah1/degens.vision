package collections

import (
	"context"
	"net/http"
	"solana_project/services/application-core/pkg/sql"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type HandleApiV1CollectionsSearchQuery struct {
	Query string `form:"query" binding:"required"`
}

type HandleApiV1CollectionsSearchResponseBody struct {
	Collections []HandleApiV1CollectionsSearchResponseBodyCollection `json:"collections"`
}

type HandleApiV1CollectionsSearchResponseBodyCollection struct {
	CollectionKey string `json:"collection_key"`
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	Image         string `json:"image"`
}

func HandleApiV1CollectionsSearch(
	ctx context.Context,
	logger *zap.SugaredLogger,
	pg *pgxpool.Pool,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var query HandleApiV1CollectionsSearchQuery
		if err := c.ShouldBind(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		rows, err := pg.Query(
			ctx,
			sql.SearchCollections,
			"%"+query.Query+"%",
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		defer rows.Close()
		response := HandleApiV1CollectionsSearchResponseBody{}
		for rows.Next() {
			var collection HandleApiV1CollectionsSearchResponseBodyCollection
			err = rows.Scan(
				&collection.CollectionKey,
				&collection.Name,
				&collection.Symbol,
				&collection.Image,
			)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
			}
			response.Collections = append(response.Collections, collection)
		}
		c.JSON(http.StatusOK, response)
	}
}
