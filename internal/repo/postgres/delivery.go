package postgres

import (
	"context"

	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/internal/repo"
	"github.com/Enthreeka/L0/pkg/db"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type repoDelivery struct {
	db *db.Postgres
}

func NewDeliveryRepository(db *db.Postgres) repo.Delivery {
	return &repoDelivery{
		db: db,
	}
}

func (r *repoDelivery) Create(ctx context.Context, id string, delivery entity.Delivery) error {

	query := `	INSERT INTO delivery VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.Pool.Exec(ctx, query, id, delivery.Phone, delivery.Name, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email)
	return err
}

func (r *repoDelivery) DeleteByID(ctx context.Context, id string) error {
	query := `DELETE FROM delivery WHERE order_uid = $1`

	_, err := r.db.Pool.Exec(ctx, query, id)
	return err
}

func (r *repoDelivery) GetByID(ctx context.Context, id string) (*entity.Delivery, error) {

	query := `SELECT * FROM delivery WHERE order_uid = $1`

	var delivery []entity.Delivery
	err := pgxscan.Select(ctx, r.db.Pool, &delivery, query, id)
	if err != nil {
		return nil, err
	}

	if len(delivery) > 0 {
		return &delivery[0], nil
	} else {
		return nil, err
	}
}

func (r *repoDelivery) GetAll(ctx context.Context) (*[]entity.Delivery, error) {

	query := `SELECT * FROM delivery`

	var delivery []entity.Delivery
	err := pgxscan.Select(ctx, r.db.Pool, &delivery, query)
	if err != nil {
		return nil, err
	}

	if len(delivery) > 0 {
		return &delivery, nil
	} else {
		return nil, err
	}

}
