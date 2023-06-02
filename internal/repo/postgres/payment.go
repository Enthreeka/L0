package postgres

import (
	"context"

	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/internal/repo"
	"github.com/Enthreeka/L0/pkg/db"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type repoPayment struct {
	db *db.Postgres
}

func NewPaymentRepository(db *db.Postgres) repo.Payment {
	return &repoPayment{
		db: db,
	}
}

func (r *repoPayment) Create(ctx context.Context, id string, payment entity.Payment) error {

	query := `INSERT INTO payment VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`

	_, err := r.db.Pool.Exec(ctx, query, id, payment.Transaction, payment.RequestID, payment.Currency, payment.Provider, payment.Amount,
		payment.PaymentDt, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee)

	return err
}

func (r *repoPayment) DeleteByID(ctx context.Context, id string) error {

	query := `DELETE FROM payment WHERE order_uid = $1`

	_, err := r.db.Pool.Exec(ctx, query, id)
	return err
}

func (r *repoPayment) GetByID(ctx context.Context, id string) (*entity.Payment, error) {

	query := `SELECT * FROM payment WHERE order_uid = $1`

	var payment []entity.Payment
	err := pgxscan.Select(ctx, r.db.Pool, &payment, query, id)
	if err != nil {
		return nil, err
	}

	if len(payment) > 0 {
		return &payment[0], nil
	} else {
		return nil, err
	}
}

func (r *repoPayment) GetAll(ctx context.Context) (*[]entity.Payment, error) {
	query := `SELECT * FROM payment`

	var payment []entity.Payment
	err := pgxscan.Select(ctx, r.db.Pool, &payment, query)
	if err != nil {
		return nil, err
	}

	if len(payment) > 0 {
		return &payment, nil
	} else {
		return nil, err
	}
}
