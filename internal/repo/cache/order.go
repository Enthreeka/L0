package cache

import (
	"context"
	"fmt"
	"sync"

	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/internal/repo"
)

type cacheOrder struct {
	cache map[string]entity.Order
	mu    sync.RWMutex
}

func NewOrderCache(cache map[string]entity.Order) repo.Order {
	return &cacheOrder{
		cache: cache,
	}
}

func (o *cacheOrder) Create(ctx context.Context, id string, order entity.Order) error {
	o.mu.Lock()
	o.cache[id] = order
	o.mu.Unlock()

	return nil
}

func (o *cacheOrder) DeleteByID(ctx context.Context, id string) {

	delete(o.cache, id)
}

func (o *cacheOrder) GetByID(ctx context.Context, id string) (*entity.Order, error) {

	data, ok := o.cache[id]
	if !ok {
		return nil, fmt.Errorf("Номер заказа не верный")
	}

	return &data, nil
}
