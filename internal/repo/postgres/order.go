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
func (r *repoOrder) Create(ctx context.Context, id string, order entity.Order) error {

	
	query := `INSERT INTO "order" VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`
	_, err := r.db.Pool.Exec(ctx, query, id, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature,
		 order.CustomerID, order.DeliveryService, order.ShardKey, order.SmID, order.DateCreated, order.OofShard)
	return err
}

// DeleteByID implements repo.Order
func (r *repoOrder) DeleteByID(ctx context.Context, id string) {
	panic("unimplemented")
}

// GetByID implements repo.Order
func (r *repoOrder) GetByID(ctx context.Context, id string) (*entity.Order, error) {
	panic("unimplemented")
}
