package auth_middleware

import (
	"context"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var identityKey = "id"

type UserAuth struct {
	Address string
}

func ExtractUserAuth(c *gin.Context) *UserAuth {
	claims := jwt.ExtractClaims(c)
	return &UserAuth{
		Address: claims[identityKey].(string),
	}
}

func ProvideAuthMiddleware(
	ctx context.Context,
	logger *zap.SugaredLogger,
	pg *pgxpool.Pool,
) (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("y9h2RPRj1J1q2wTpYA"),
		Timeout:     time.Hour * 24 * 365, // year
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*UserAuth); ok {
				return jwt.MapClaims{
					identityKey: v.Address,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &UserAuth{
				Address: claims[identityKey].(string),
			}
		},
		Authenticator: provideAuthenticator(ctx, logger, pg),
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		return nil, err
	}

	return authMiddleware, nil
}
