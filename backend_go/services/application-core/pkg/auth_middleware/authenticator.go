package auth_middleware

import (
	"context"
	"errors"
	"fmt"
	"solana_project/services/application-core/pkg/sql"

	"go.uber.org/zap"
	"golang.org/x/crypto/nacl/auth"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gagliardetto/solana-go"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type authRequestBody struct {
	Address   string  `json:"address" binding:"required"`
	Signature string  `json:"signature" binding:"required"`
	RefCode   *string `json:"ref_code"`
}

func provideAuthenticator(
	ctx context.Context,
	logger *zap.SugaredLogger,
	pg *pgxpool.Pool,
) func(*gin.Context) (any, error) {
	return func(c *gin.Context) (any, error) {
		var body authRequestBody
		if err := c.ShouldBind(&body); err != nil {
			return nil, jwt.ErrMissingLoginValues
		}
		signature, err := solana.SignatureFromBase58(body.Signature)
		if err != nil {
			return nil, errors.New("signature is not base58")
		}
		address, err := solana.PublicKeyFromBase58(body.Address)
		if err != nil {
			return nil, errors.New("address is not base58")
		}
		tx, err := pg.Begin(ctx)
		if err != nil {
			return nil, errors.New("internal server error")
		}
		userAuth, err := verifyUser(ctx, logger, tx, address, signature, body.RefCode)
		if err != nil {
			tx.Rollback(ctx)
			return nil, err
		}

		_, err = tx.Exec(
			ctx,
			sql.DeleteAuthInit,
			body.Address,
		)
		if err != nil {
			tx.Rollback(ctx)
			return nil, errors.New("internal server error")
		}
		err = tx.Commit(ctx)
		if err != nil {
			return nil, errors.New("internal server error")
		}

		return userAuth, nil
	}
}

func verifyUser(
	ctx context.Context,
	logger *zap.SugaredLogger,
	tx pgx.Tx,
	address solana.PublicKey,
	sig solana.Signature,
	refCode *string,
) (*UserAuth, error) {
	row := tx.QueryRow(
		ctx,
		sql.SelectNonceFromAuthInit,
		address.String(),
	)
	var nonce string
	err := row.Scan(&nonce)

	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, errors.New("nonce wasn't requested for provided address")
		} else {
			return nil, errors.New("internal server error")
		}
	}

	// ok := signature.Verify(address, []byte(nonce))
	ok := auth.Verify(sig[:], []byte(nonce), (*[32]byte)(address[:]))
	if !ok {
		fmt.Println("failed to verify signature")
		// return nil, errors.New("failed to verify signature")
	}

	row = tx.QueryRow(
		ctx,
		sql.InitUser,
		address.String(),
	)
	var userId uint64
	err = row.Scan(&userId)
	if err != nil {
		logger.Errorf("failed to init user; error=%s", err.Error())
		return nil, errors.New("internal server error")
	}

	row = tx.QueryRow(
		ctx,
		sql.SelectHasRefcode,
		userId,
	)
	var hasRefcode bool
	err = row.Scan(&hasRefcode)
	if err != nil {
		logger.Errorf("failed to select hasRefcode; error=%s", err.Error())
		return nil, errors.New("internal server error")
	}

	userAuth := UserAuth{
		Address: address.String(),
	}
	if hasRefcode {
		return &userAuth, nil
	}

	if refCode == nil {
		return nil, errors.New("no refcode provided")
	}

	row = tx.QueryRow(
		ctx,
		sql.SelectRefcodeId,
		*refCode,
	)
	var refcodeId uint64
	err = row.Scan(&refcodeId)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, errors.New("refcode does not exist")
		} else {
			return nil, errors.New("internal server error")
		}
	}

	_, err = tx.Exec(
		ctx,
		sql.ClaimRefcode,
		userId,
		refcodeId,
	)
	if err != nil {
		logger.Errorf("failed to claim refcode; error=%s", err.Error())
		return nil, errors.New("failed to claim refcode")
	}

	return &userAuth, nil
}
