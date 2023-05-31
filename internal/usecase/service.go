package usecase

import (
	"context"

	"github.com/Enthreeka/L0/internal/entity"
)

type Delivery interface {
	CreateDelivery(ctx context.Context, id string, delivery entity.Delivery) error
	GetByID(ctx context.Context, id string) (*entity.Delivery, error)
	SaveAllToCache(ctx context.Context) error
}

type Item interface {
	CreateItem(ctx context.Context, id string, item entity.Item) error
	GetByID(ctx context.Context, id string) (*[]entity.Item, error)
	SaveAllToCache(ctx context.Context) error
}

type Order interface {
	CreateOrder(ctx context.Context, id string, order entity.Order) error
	GetByID(ctx context.Context, id string) (*entity.Order, error)
	SaveAllToCache(ctx context.Context) error
}

type Payment interface {
	CreatePayment(ctx context.Context, id string, payment entity.Payment) error
	GetByID(ctx context.Context, id string) (*entity.Payment, error)
	SaveAllToCache(ctx context.Context) error
}
