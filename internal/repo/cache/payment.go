package cache

import (
	"context"
	"fmt"
	"sync"

	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/internal/repo"
)

type cachePayment struct {
	cache map[string]entity.Payment
	mu    sync.RWMutex
}

func NewPaymentCache(cache map[string]entity.Payment) repo.Payment {
	return &cachePayment{
		cache: cache,
	}
}

func (c *cachePayment) Create(ctx context.Context, id string, payment entity.Payment) error {
	c.mu.RLock()
	c.cache[id] = payment
	c.mu.RUnlock()
	return nil
}

func (c *cachePayment) DeleteByID(ctx context.Context, id string) error {
	delete(c.cache, id)

	return nil
}

func (c *cachePayment) GetByID(ctx context.Context, id string) (*entity.Payment, error) {
	data, ok := c.cache[id]
	if !ok {
		return nil, fmt.Errorf("%s", "order number for payment invalible")
	}

	return &data, nil
}

func (*cachePayment) GetAll(ctx context.Context) (*[]entity.Payment, error) {
	panic("unimplemented")
}
