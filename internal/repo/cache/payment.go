package cache

import (
	"context"
	"fmt"

	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/internal/repo"
)

type cachePayment struct {
	cache map[string]entity.Payment
}

func NewPaymentCache(cache map[string]entity.Payment) repo.Payment {
	return &cachePayment{
		cache: cache,
	}
}

func (c *cachePayment) Create(ctx context.Context, id string, payment entity.Payment) error {
	if _, ok := c.cache[id]; ok {
		return fmt.Errorf("%s", "Номер заказа не верный")
	}

	c.cache[id] = payment

	return nil
}

func (c *cachePayment) DeleteByID(ctx context.Context, id string) error {
	delete(c.cache, id)

	return nil
}

func (c *cachePayment) GetByID(ctx context.Context, id string) (*entity.Payment, error) {
	data, ok := c.cache[id]
	if !ok {
		return nil, fmt.Errorf("Order number for payment invalible")
	}

	return &data, nil
}

func (*cachePayment) GetAll(ctx context.Context) (*[]entity.Payment, error) {
	panic("unimplemented")
}
