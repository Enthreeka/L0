package repo

import (
	"context"

	"github.com/Enthreeka/L0/internal/entity"
)

type Delivery interface {
	GetByID(ctx context.Context, id string) (*entity.Delivery, error)
	Create(ctx context.Context, id string, delivery entity.Delivery) error
	DeleteByID(ctx context.Context, id string) error
	GetAll(ctx context.Context) (*[]entity.Delivery, error)
}

type Item interface {
	GetByID(ctx context.Context, id string) (*[]entity.Item, error)
	Create(ctx context.Context, id string, item entity.Item) error
	DeleteByID(ctx context.Context, id string) error
	GetAll(ctx context.Context) (*[]entity.Item, error)
}

type Order interface {
	GetByID(ctx context.Context, id string) (*entity.Order, error)
	Create(ctx context.Context, id string, order entity.Order) error
	DeleteByID(ctx context.Context, id string) error
	GetAll(ctx context.Context) (*[]entity.Order, error)
}

type Payment interface {
	GetByID(ctx context.Context, id string) (*entity.Payment, error)
	Create(ctx context.Context, id string, payment entity.Payment) error
	DeleteByID(ctx context.Context, id string) error
	GetAll(ctx context.Context) (*[]entity.Payment, error)
}
