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
	
	o.mu.RLock()
	o.cache[id] = order
	o.mu.RUnlock()
	return nil
}

func (o *cacheOrder) DeleteByID(ctx context.Context, id string) error {
	delete(o.cache, id)

	return nil
}

func (o *cacheOrder) GetByID(ctx context.Context, id string) (*entity.Order, error) {

	data, ok := o.cache[id]
	if !ok {
		return nil, fmt.Errorf("invalible order number: %s", id)
	}

	return &data, nil
}

func (*cacheOrder) GetAll(ctx context.Context) (*[]entity.Order, error) {
	panic("unimplemented")
}
