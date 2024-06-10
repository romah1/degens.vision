package users

import (
	"context"
	"net/http"
	"solana_project/services/application-core/pkg/auth_middleware"
	"solana_project/services/application-core/pkg/models/model_users"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func HandleApiV1UpdateUser(
	ctx context.Context,
	logger *zap.SugaredLogger,
	pg *pgxpool.Pool,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model_users.User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		userAuth := auth_middleware.ExtractUserAuth(c)
		if user.Address != userAuth.Address {
			c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
			return
		}
		err := user.Upsert(ctx, pg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
