package refcodes

import (
	"context"
	"net/http"
	"solana_project/services/application-core/pkg/sql"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type HandleApiV1AdminRefcodesCreateRequest struct {
	Code      string `json:"code" binding:"required"`
	MaxUsages int    `json:"max_usages" binding:"required"`
}

func HandleAdminRefcodesCreate(
	ctx context.Context,
	logger *zap.SugaredLogger,
	pg *pgxpool.Pool,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request HandleApiV1AdminRefcodesCreateRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}

		_, err := pg.Exec(
			ctx,
			sql.InsertRefcode,
			request.Code,
			request.MaxUsages,
		)
		if err != nil {
			logger.Errorf("failed to create refcode; err=%s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
