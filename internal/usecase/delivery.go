package usecase

import (
	"context"

	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/internal/repo"
	"github.com/Enthreeka/L0/pkg/logger"
)

type deliveryService struct {
	db    repo.Delivery
	cache repo.Delivery

	log *logger.Logger
}

func NewDeliveryService(db repo.Delivery, cache repo.Delivery, log *logger.Logger) Delivery {
	return &deliveryService{
		db:    db,
		cache: cache,
		log:   log,
	}
}

func (d *deliveryService) CreateDelivery(ctx context.Context, id string, delivery entity.Delivery) error {

	err := d.db.Create(ctx, id, delivery)
	if err != nil {
		d.log.Error("Error with create delivery in db %v", err)
		return err
	}

	err = d.cache.Create(ctx, id, delivery)
	if err != nil {
		d.log.Error("Error with create delivery in cache %v", err)
		return err
	}

	return nil
}

func (d *deliveryService) GetByID(ctx context.Context, id string) (*entity.Delivery, error) {

	delivery, err := d.cache.GetByID(ctx, id)
	if err != nil {
		d.log.Error("Error with get id %v", err)
	}

	return delivery, nil
}

func (d *deliveryService) SaveAllToCache(ctx context.Context) error {

	delivery, err := d.db.GetAll(ctx)
	if err != nil {
		d.log.Error("Error to get all data from delivery db %v", err)
		return err
	}

	for _, v := range *delivery {
		d.cache.Create(ctx, v.OrderUID, v)
	}

	return nil
}
