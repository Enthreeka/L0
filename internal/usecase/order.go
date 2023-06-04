package usecase

import (
	"context"

	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/internal/repo"
	"github.com/Enthreeka/L0/pkg/logger"
)

type orderService struct {
	db    repo.Order
	cache repo.Order

	log *logger.Logger
}

func NewOrderService(db repo.Order, cache repo.Order, log *logger.Logger) Order {
	return &orderService{
		db:    db,
		cache: cache,
		log:   log,
	}
}

func (o *orderService) CreateOrder(ctx context.Context, id string, order entity.Order) error {

	err := o.db.Create(ctx, id, order)
	if err != nil {
		o.log.Error("Error with create order in db %v", err)
		return err
	}

	err = o.cache.Create(ctx, id, order)
	if err != nil {
		o.log.Error("Error with create order in cache %v", err)
		return err
	}

	return nil
}

func (o *orderService) GetByID(ctx context.Context, id string) (*entity.Order, error) {

	order, err := o.cache.GetByID(ctx, id)
	if err != nil {
		o.log.Error("Error with get id %v", err)
		return nil, err
	}

	return order, nil
}

func (o *orderService) SaveAllToCache(ctx context.Context) error {

	order, err := o.db.GetAll(ctx)
	if err != nil {
		o.log.Error("Error to get all data from delivery db %v", err)
		return err
	}

	if order != nil {
		for _, v := range *order {
			o.cache.Create(ctx, v.OrderUID, v)
		}
	}

	return nil
}

func (o *orderService) DeleteByID(ctx context.Context, id string) error {

	err := o.db.DeleteByID(ctx, id)
	if err != nil {
		o.log.Error("%s:", err)
		return err
	}

	err = o.cache.DeleteByID(ctx, id)
	if err != nil {
		o.log.Error("%s:", err)
		return err
	}

	return nil
}
