package postgres

import (
	"context"

	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/internal/repo"
	"github.com/Enthreeka/L0/pkg/db"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type repoItem struct {
	db *db.Postgres
}

func NewItemRepo(db *db.Postgres) repo.Item {
	return &repoItem{
		db: db,
	}
}

func (r *repoItem) Create(ctx context.Context, id string, item entity.Item) error {

	query := `INSERT INTO item VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err := r.db.Pool.Exec(ctx, query, item.ChrtID, item.TrackNumber, id, item.Price, item.RID, item.Name, item.Sale,
		item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status)

	return err
}

func (r *repoItem) DeleteByID(ctx context.Context, id string) error {
	query := `DELETE FROM item WHERE order_uid = $1`

	_, err := r.db.Pool.Exec(ctx, query, id)
	return err
}

func (r *repoItem) GetByID(ctx context.Context, id string) (*[]entity.Item, error) {

	query := `SELECT * FROM item WHERE order_uid = $1`

	var item []entity.Item
	err := pgxscan.Select(ctx, r.db.Pool, &item, query, id)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *repoItem) GetAll(ctx context.Context) (*[]entity.Item, error) {
	query := `SELECT * FROM item`

	var item []entity.Item
	err := pgxscan.Select(ctx, r.db.Pool, &item, query)
	if err != nil {
		return nil, err
	}

	if len(item) > 0 {
		return &item, nil
	} else {
		return nil, err
	}
}
