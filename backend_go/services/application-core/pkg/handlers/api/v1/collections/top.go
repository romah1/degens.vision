package collections

import (
	"context"
	"net/http"
	"solana_project/services/application-core/pkg/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type HandleTopCollectionsResponse struct {
	Collections []HandleTopCollectionsResponseCollection `json:"collections"`
}

type HandleTopCollectionsResponseCollection struct {
	CollectionKey    string  `json:"collection_key"`
	MintsCountPeriod int     `json:"mints_count_period"`
	MintsCountTotal  int     `json:"mints_count_total"`
	Launchpad        string  `json:"launchpad"`
	Name             string  `json:"name"`
	Symbol           string  `json:"symbol"`
	Image            string  `json:"image"`
	Size             int     `json:"size"`
	MintPrice        float64 `json:"mint_price"`
}

func HandleApiV1CollectionsTop(
	ctx context.Context,
	logger *zap.SugaredLogger,
	pgxpool *pgxpool.Pool) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		minutesGapRaw := c.Request.URL.Query().Get("minutes_gap")
		if minutesGapRaw == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing minutes_gap param"})
			return
		}

		minutesGap, err := strconv.ParseInt(minutesGapRaw, 10, 0)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "minutes_gap param expected to be int"})
			return
		}

		rows, err := pgxpool.Query(ctx, sql.SelectTopCollections, minutesGap)
		if err != nil {
			logger.Errorf("failed to select collections from db; error=%s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		defer rows.Close()

		collections := []HandleTopCollectionsResponseCollection{}
		for rows.Next() {
			var collection HandleTopCollectionsResponseCollection
			err = rows.Scan(
				&collection.CollectionKey,
				&collection.MintsCountPeriod,
				&collection.MintsCountTotal,
				&collection.Name,
				&collection.Symbol,
				&collection.Image,
				&collection.Size,
				&collection.Launchpad,
				&collection.MintPrice,
			)
			if err != nil {
				logger.Errorf("failed to scan row while selecting collections from db; error=%s", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
			}
			collections = append(collections, collection)
		}
		response := HandleTopCollectionsResponse{
			Collections: collections,
		}

		c.JSON(http.StatusOK, response)
	}
	return gin.HandlerFunc(fn)
}
