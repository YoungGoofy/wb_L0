package services

import (
	"context"
	"github.com/YoungGoofy/WB_L0/internal/models"
)

type PGRepository interface {
	Create(ctx context.Context, order *models.Orders) error
	CreateOrder(ctx context.Context, orders *models.Orders) (*models.Orders, error)
	GetOrderById(ctx context.Context, uid string) (models.Orders, error)
	GetFullById(ctx context.Context, uid string) (*models.Orders, error)
}

type CacheRepository interface {
	Set(ctx context.Context, order *models.Orders) error
	GetById(ctx context.Context, uid string) (*models.Orders, error)
	Delete(ctx context.Context, uid string) error
	GetSize() int
}
