package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Gres struct {
	Pool *pgxpool.Pool
}

func New(url string) (Gres, error) {
	pool, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		return Gres{}, fmt.Errorf("can't connect to db: %w", err)
	}
	return Gres{
		Pool: pool,
	}, nil
}
