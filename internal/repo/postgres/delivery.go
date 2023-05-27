package postgres

import (
	"context"

	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/internal/repo"
	"github.com/Enthreeka/L0/pkg/db"
)

type repoDelivery struct {
	db *db.Postgres
}

func NewDeliveryRepository(db *db.Postgres) repo.Delivery {
	return &repoDelivery{
		db: db,
	}
}

// Create implements repo.Delivery
func (r *repoDelivery) Create(ctx context.Context, id string,delivery entity.Delivery) {


	panic("unimplemented")
}

// DeleteByID implements repo.Delivery
func (r *repoDelivery) DeleteByID(ctx context.Context, id string) {
	panic("unimplemented")
}

// GetByID implements repo.Delivery
func (r *repoDelivery) GetByID(ctx context.Context, id string) (*entity.Delivery, error) {
	panic("unimplemented")
}

// // GetByID implements repo.Delivery
// func (repoDelivery) GetByID(id string) (entity.Delivery, error) {
// 	panic("unimplemented")
// }
