package usecase

import (
	"context"

	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/internal/repo"
	"github.com/Enthreeka/L0/pkg/logger"
)

type itemService struct {
	db    repo.Item
	cache repo.Item

	log *logger.Logger
}

func NewItemService(db repo.Item, chache repo.Item, log *logger.Logger) Item {
	return &itemService{
		db:    db,
		cache: chache,
		log:   log,
	}
}

func (i *itemService) CreateItem(ctx context.Context, id string, item entity.Item) error {

	err := i.db.Create(ctx, id, item)
	if err != nil {
		i.log.Error("Error with create item in db %v", err)
		return err
	}

	err = i.cache.Create(ctx, id, item)
	if err != nil {
		i.log.Error("Error with create item in cache %v", err)
		return err
	}

	return nil

}

func (i *itemService) GetByID(ctx context.Context, id string) (*[]entity.Item, error) {

	item, err := i.cache.GetByID(ctx, id)
	if err != nil {
		i.log.Error("Error with get id %v", err)
		return nil, err
	}

	return item, nil
}

func (i *itemService) SaveAllToCache(ctx context.Context) error {

	item, err := i.db.GetAll(ctx)
	if err != nil {
		i.log.Error("Error to get all data from delivery db %v", err)
		return err
	}

	if item != nil {
		for _, v := range *item {
			i.cache.Create(ctx, v.OrderUID, v)
		}
	}

	return nil
}

func (i *itemService) DeleteByID(ctx context.Context, id string) error {

	err := i.db.DeleteByID(ctx, id)
	if err != nil {
		i.log.Error("%s:", err)
		return err
	}

	err = i.cache.DeleteByID(ctx, id)
	if err != nil {
		i.log.Error("%s:", err)
		return err
	}

	return nil
}
