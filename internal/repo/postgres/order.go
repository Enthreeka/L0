package postgres

import (
	"context"

	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/internal/repo"
	"github.com/Enthreeka/L0/pkg/db"
)

type repoOrder struct {
	db *db.Postgres
}

func NewOrderRepository(db *db.Postgres) repo.Order {
	return &repoOrder{
		db: db,
	}
}

// Create implements repo.Order
func (r *repoOrder) Create(ctx context.Context, id string, delivery entity.Order) {
	panic("unimplemented")
}

// DeleteByID implements repo.Order
func (r *repoOrder) DeleteByID(ctx context.Context, id string) {
	panic("unimplemented")
}

// GetByID implements repo.Order
func (r *repoOrder) GetByID(ctx context.Context, id string) (*entity.Order, error) {
	panic("unimplemented")
}
