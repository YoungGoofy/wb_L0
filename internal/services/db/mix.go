package db

import (
	"context"
	"github.com/YoungGoofy/WB_L0/internal/models"
	"github.com/YoungGoofy/WB_L0/internal/services"
	"log"
)

type Repositories struct {
	pgRepo    services.PGRepository
	cacheRepo services.CacheRepository
}

func NewRepositories(pgRepository services.PGRepository, cacheRepository services.CacheRepository) *Repositories {
	return &Repositories{pgRepo: pgRepository, cacheRepo: cacheRepository}
}

func (r *Repositories) Create(ctx context.Context, orders *models.Orders) error {
	err := r.pgRepo.Create(ctx, orders)
	if err != nil {
		log.Fatal("Error for creating new order")
		return err
	}
	log.Println("Add order in database")
	return nil
}

func (r *Repositories) GetById(ctx context.Context, uid string) (*models.Orders, error) {
	if r.cacheRepo.GetSize() == 0 {
		UIDs, err := r.pgRepo.GetIds(ctx)
		if err != nil {
			return nil, nil
		}
		for _, id := range UIDs {
			order, err := r.pgRepo.GetFullById(ctx, id)
			if err != nil {
				return nil, nil
			}
			err = r.cacheRepo.Set(ctx, order)
			if err != nil {
				return nil, nil
			}
		}
	}
	cache, err := r.cacheRepo.GetById(ctx, uid)
	if err != nil {
		return nil, nil
	}
	if cache != nil {
		return cache, nil
	}
	order, err := r.pgRepo.GetFullById(ctx, uid)
	if err != nil {
		return nil, nil
	}

	err = r.cacheRepo.Set(ctx, order)
	if err != nil {
		return nil, nil
	}
	return order, nil
}
