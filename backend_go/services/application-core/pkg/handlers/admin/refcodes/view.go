package refcodes

import (
	"context"
	"net/http"
	"solana_project/services/application-core/pkg/sql"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type HandleAdminRefcodesViewResponse struct {
	Refcodes []HandleAdminRefcodesViewResponseRefcode `json:"refcodes"`
}
type HandleAdminRefcodesViewResponseRefcode struct {
	Code          string `json:"code"`
	MaxUsages     int    `json:"max_usages"`
	CurrentUsages int    `json:"current_usages"`
}

func HandleAdminRefcodesView(
	ctx context.Context,
	logger *zap.SugaredLogger,
	pg *pgxpool.Pool,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		refcodes, err := selectRefcodes(ctx, pg)
		if err != nil {
			logger.Errorf("failed to select refcodes; err=%s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		response := HandleAdminRefcodesViewResponse{
			Refcodes: refcodes,
		}
		c.JSON(http.StatusOK, response)
	}
}

func selectRefcodes(
	ctx context.Context,
	pg *pgxpool.Pool,
) ([]HandleAdminRefcodesViewResponseRefcode, error) {
	rows, err := pg.Query(
		ctx,
		sql.SelectRefcodes,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []HandleAdminRefcodesViewResponseRefcode
	for rows.Next() {
		var refode HandleAdminRefcodesViewResponseRefcode
		err = rows.Scan(
			&refode.Code,
			&refode.MaxUsages,
			&refode.CurrentUsages,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, refode)
	}
	return res, nil
}
