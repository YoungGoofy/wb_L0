package postgresql

import (
	"context"
	"fmt"
	"github.com/YoungGoofy/WB_L0/internal/config"
	"github.com/YoungGoofy/WB_L0/internal/utils"
	"github.com/jackc/pgx/v4/pgxpool"
	"strconv"
	"time"
)

func NewClient(ctx context.Context, cfg *config.Config) (pool *pgxpool.Pool, err error) {
	localData := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseName)
	maxAttempts, _ := strconv.Atoi(cfg.MaxAttempts)
	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, localData)
		if err != nil {
			return err
		}
		return nil
	}, maxAttempts, time.Second*5)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
