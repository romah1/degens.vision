package collections

import (
	"context"
	"net/http"
	"solana_project/services/application-core/pkg/sql"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type HandleTopHoldersResponse struct {
	Holders []HandleTopHoldersResponseHolder `json:"holders"`
}

type HandleTopHoldersResponseHolder struct {
	Key    string `json:"key"`
	Amount int    `json:"amount"`
}

func HandleTopHolders(
	ctx context.Context,
	logger *zap.SugaredLogger,
	pgxpool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		collectionKey := c.Request.URL.Query().Get("collection_key")
		if collectionKey == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing collection_key param"})
			return
		}

		rows, err := pgxpool.Query(ctx, sql.SelectTopHolders, collectionKey)
		if err != nil {
			logger.Errorf("failed to select top holders from db; error=%s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		defer rows.Close()

		holders := []HandleTopHoldersResponseHolder{}
		for rows.Next() {
			var holder HandleTopHoldersResponseHolder
			err = rows.Scan(
				&holder.Key,
				&holder.Amount,
			)
			if err != nil {
				logger.Errorf("failed to scan row while selecting top holders from db; error=%s", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
			}
			holders = append(holders, holder)
		}
		response := HandleTopHoldersResponse{
			Holders: holders,
		}

		c.JSON(http.StatusOK, response)
	}
	return gin.HandlerFunc(fn)
}
