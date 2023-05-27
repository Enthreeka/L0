package repo

import (
	"context"

	"github.com/Enthreeka/L0/internal/entity"
)

type Delivery interface {
	GetByID(ctx context.Context, id string) (*entity.Delivery, error)
	Create(ctx context.Context, id string, delivery entity.Delivery)
	DeleteByID(ctx context.Context, id string)
}

type Item interface{}

type Order interface {
	GetByID(ctx context.Context, id string) (*entity.Order, error)
	Create(ctx context.Context, id string, order entity.Order) error
	DeleteByID(ctx context.Context, id string)
}

type Payment interface{}
