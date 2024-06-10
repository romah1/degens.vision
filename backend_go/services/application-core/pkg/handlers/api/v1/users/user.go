package users

import (
	"context"
	"net/http"
	"solana_project/services/application-core/pkg/models/model_users"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func HandleApiV1UsersUser(
	ctx context.Context,
	logger *zap.SugaredLogger,
	pg *pgxpool.Pool,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		address := c.Param("address")
		user, err := model_users.SelectUser(ctx, pg, address)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
