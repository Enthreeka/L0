package postgres

import (
	"context"

	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/internal/repo"
	"github.com/Enthreeka/L0/pkg/db"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type repoOrder struct {
	db *db.Postgres
}

func NewOrderRepository(db *db.Postgres) repo.Order {
	return &repoOrder{
		db: db,
	}
}

func (r *repoOrder) Create(ctx context.Context, id string, order entity.Order) error {

	query := `INSERT INTO "order" VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,CURRENT_TIMESTAMP,$10)`

	_, err := r.db.Pool.Exec(ctx, query, id, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature,
		order.CustomerID, order.DeliveryService, order.ShardKey, order.SmID, order.OofShard)
	return err
}

func (r *repoOrder) DeleteByID(ctx context.Context, id string) error {
	query := `DELETE FROM "order" WHERE order_uid = $1`

	_, err := r.db.Pool.Exec(ctx, query, id)
	return err
}

func (r *repoOrder) GetByID(ctx context.Context, id string) (*entity.Order, error) {

	query := `SELECT * FROM "order" WHERE order_uid = $1`

	var order []entity.Order
	err := pgxscan.Select(ctx, r.db.Pool, &order, query, id)
	if err != nil {
		return nil, err
	}

	if len(order) > 0 {
		return &order[0], nil
	} else {
		return nil, err
	}
}

// GetAll implements repo.Order
func (r *repoOrder) GetAll(ctx context.Context) (*[]entity.Order, error) {

	query := `SELECT * FROM "order"`

	var order []entity.Order
	err := pgxscan.Select(ctx, r.db.Pool, &order, query)
	if err != nil {
		return nil, err
	}

	if len(order) > 0 {
		return &order, nil
	} else {
		return nil, err
	}
}
