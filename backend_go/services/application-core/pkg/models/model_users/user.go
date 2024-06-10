package model_users

import (
	"context"
	"solana_project/services/application-core/pkg/sql"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Address string  `json:"address"`
	Name    *string `json:"name"`
	Image   *string `json:"image"`
}

func InitUser(ctx context.Context, pg *pgxpool.Pool, address string) (uint64, error) {
	row := pg.QueryRow(
		ctx,
		sql.InitUser,
		address,
	)
	var userId uint64
	err := row.Scan(&userId)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func SelectUser(ctx context.Context, pg *pgxpool.Pool, address string) (*User, error) {
	row := pg.QueryRow(
		ctx,
		sql.SelectUser,
		address,
	)
	var user User
	err := row.Scan(&user.Address, &user.Name, &user.Image)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) Upsert(ctx context.Context, pg *pgxpool.Pool) error {
	_, err := pg.Exec(
		ctx,
		sql.UpsertUser,
		u.Address,
		u.Name,
		u.Image,
	)
	if err != nil {
		return err
	}
	return nil
}

func HasRefcode(ctx context.Context, pg *pgxpool.Pool, userId uint64) (bool, error) {
	row := pg.QueryRow(ctx, sql.SelectHasRefcode, userId)
	var hasRefcode bool
	err := row.Scan(&hasRefcode)
	if err != nil {
		return false, err
	}
	return hasRefcode, nil
}
