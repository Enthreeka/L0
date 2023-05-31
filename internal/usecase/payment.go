package usecase

import (
	"context"

	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/internal/repo"
	"github.com/Enthreeka/L0/pkg/logger"
)

type paymentService struct {
	db    repo.Payment
	cache repo.Payment

	log *logger.Logger
}

func NewPaymentService(db repo.Payment, cache repo.Payment, log *logger.Logger) Payment {
	return &paymentService{
		db:    db,
		cache: cache,
		log:   log,
	}
}

func (p *paymentService) CreatePayment(ctx context.Context, id string, payment entity.Payment) error {

	err := p.db.Create(ctx, id, payment)
	if err != nil {
		p.log.Error("Error with create payment in db %v", err)
		return err
	}

	err = p.cache.Create(ctx, id, payment)
	if err != nil {
		p.log.Error("Error with create payment in cache %v", err)
		return err
	}

	return nil
}

func (p *paymentService) GetByID(ctx context.Context, id string) (*entity.Payment, error) {

	payment, err := p.cache.GetByID(ctx, id)
	if err != nil {
		p.log.Error("Error with get id %v", err)
	}

	return payment, nil
}

func (p *paymentService) SaveAllToCache(ctx context.Context) error {

	payment, err := p.db.GetAll(ctx)
	if err != nil {
		p.log.Error("Error to get all data from payment db %v", err)
		return err
	}

	for _, v := range *payment {
		p.cache.Create(ctx, v.OrderUID, v)
	}

	return nil

}
