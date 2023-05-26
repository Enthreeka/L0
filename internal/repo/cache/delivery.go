package cache

import (
	"context"
	"sync"

	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/internal/repo"
)

type cacheDelivery struct {
	cache map[string]entity.Delivery
	mu    sync.RWMutex
}

func NewDeliveryCache(cache map[string]entity.Delivery, mu sync.RWMutex) repo.Delivery {
	return &cacheDelivery{
		cache: cache,
		mu:    mu,
	}
}

// Create implements repo.Delivery
func (c *cacheDelivery) Create(ctx context.Context, id string, delivery entity.Delivery) {
	c.mu.RLock()
	c.cache[id] = delivery
	c.mu.RUnlock()

	panic("unimplemented")
}

// DeleteByID implements repo.Delivery
func (c *cacheDelivery) DeleteByID(ctx context.Context, id string) {

	delete(c.cache, id)

	panic("unimplemented")
}

// GetByID implements repo.Delivery
func (c *cacheDelivery) GetByID(ctx context.Context, id string) (*entity.Delivery, error) {

	
	//data, _ := d.cache[id]
	var data entity.Delivery
	c.cache[id] = data
	

	return &data, nil
	//panic("unimplemented")
}

// func (d *OrderCache) GetByID(id string) data []entity.Order {
// 	d.mu.RLock()
// 	//data, _ := d.cache[id]
// 	//var data []entity.Order
// 	d.cache[id] = data
// 	d.mu.RUnlock()

// 	return data
// }

// Реализовать repo постгреса, реализовать cache
