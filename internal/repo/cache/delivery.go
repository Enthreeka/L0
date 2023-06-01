package cache

import (
	"context"
	"fmt"
	"sync"

	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/internal/repo"
)

type cacheDelivery struct {
	cache map[string]entity.Delivery
	mu    sync.RWMutex
}

func NewDeliveryCache(cache map[string]entity.Delivery) repo.Delivery {
	return &cacheDelivery{
		cache: cache,
	}
}

func (c *cacheDelivery) Create(ctx context.Context, id string, delivery entity.Delivery) error {
	c.mu.RLock()
	c.cache[id] = delivery
	c.mu.RUnlock()

	return nil
}

func (c *cacheDelivery) DeleteByID(ctx context.Context, id string) error {

	delete(c.cache, id)

	return nil
}

func (c *cacheDelivery) GetByID(ctx context.Context, id string) (*entity.Delivery, error) {

	data, ok := c.cache[id]
	if !ok {
		return nil, fmt.Errorf("%s", "Order number for delivery invalible")
	}

	return &data, nil
}

func (*cacheDelivery) GetAll(ctx context.Context) (*[]entity.Delivery, error) {
	panic("unimplemented")
}
