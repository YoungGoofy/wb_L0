package db

import (
	"context"
	"encoding/json"
	"github.com/YoungGoofy/WB_L0/internal/models"
	"github.com/YoungGoofy/WB_L0/internal/services/cache"
)

type OrderCache struct {
	cache *cache.Cache
}

func NewCache(cache *cache.Cache) *OrderCache {
	cache.Data = make(map[string]string)
	return &OrderCache{cache: cache}
}

func (oc *OrderCache) Set(ctx context.Context, order *models.Orders) error {
	orderBytes, err := json.Marshal(order)
	if err != nil {
		return err
	}

	oc.cache.Put(order.OrderUID, string(orderBytes))
	return nil
}

func (oc *OrderCache) GetById(ctx context.Context, uid string) (*models.Orders, error) {
	item, err := oc.cache.Get(uid)
	if err != nil {
		return nil, err
	}

	var order models.Orders

	err = json.Unmarshal([]byte(item), &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (oc *OrderCache) Delete(ctx context.Context, uid string) {
	oc.cache.Del(uid)
}

func (oc *OrderCache) GetSize() int {
	return len(oc.cache.Data)
}
