package auth

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"solana_project/services/application-core/pkg/sql"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type HandleApiV1AuthInitRequestBody struct {
	Address string `json:"address" binding:"required"`
}

type HandleApiV1AuthInitResponseBody struct {
	Nonce string `json:"nonce"`
}

func HandleApiV1AuthInit(
	ctx context.Context,
	logger *zap.SugaredLogger,
	pgxpool *pgxpool.Pool,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body HandleApiV1AuthInitRequestBody
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		nonce := genNonce(body.Address, 10)
		_, err := pgxpool.Exec(
			ctx,
			sql.UpsertAuthInit,
			body.Address,
			nonce,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		c.JSON(http.StatusOK, HandleApiV1AuthInitResponseBody{
			Nonce: nonce,
		})
	}
}

var letterRunes = []rune("123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func genNonce(address string, n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return fmt.Sprintf(`Degens.vision needs to verify that the wallet %s belongs to you.

This is a simple signature process, with no associated fees, blockchain operations, fund transfers, or rights transfers of any kind.	

The signature is only for authorization purposes and nothing more.

Nonce: %s`, address, string(b))
}
