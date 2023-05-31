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
	}

	return order, nil
}

func (o *orderService) SaveAllToCache(ctx context.Context) error {

	order, err := o.db.GetAll(ctx)
	if err != nil {
		o.log.Error("Error to get all data from delivery db %v", err)
		return err
	}

	for _, v := range *order {
		o.cache.Create(ctx, v.OrderUID, v)
	}

	return nil
}

// type orderService struct {
// 	orderRepo    repo.Order
// 	itemRepo     repo.Item
// 	deliveryRepo repo.Delivery
// 	paymentRepo  repo.Payment
// 	orderCache   repo.Order

// 	log *logger.Logger
// }

// type Order struct {
// 	entity.Order    `json:"order"`
// 	entity.Delivery `json:"delivery"`
// 	entity.Payment  `json:"payment"`
// 	Item            []entity.Item `json:"item"`
// }

// func NewOrderService(orderRepo repo.Order, itemRepo repo.Item, deliveryRepo repo.Delivery, paymentRepo repo.Payment, orderCache repo.Order, log *logger.Logger) *orderService {
// 	return &orderService{
// 		orderRepo:    orderRepo,
// 		itemRepo:     itemRepo,
// 		deliveryRepo: deliveryRepo,
// 		paymentRepo:  paymentRepo,
// 		log:          log,
// 	}
// }

// func (o *orderService) GetOrderByID(ctx context.Context, id string) (*Order, error) {

// 	//	dataOrder := new(Order)

// 	//var dataOrder Order

// 	or, err := o.orderCache.GetByID(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	dt := Order{
// 		Order: *or,
// 	}

// 	// order, err := o.orderRepo.GetByID(ctx, id)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// if order != nil {
// 	// 	dataOrder.Order = *order
// 	// }

// 	// payment, err := o.paymentRepo.GetByID(ctx, id)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// if payment != nil {
// 	// 	dataOrder.Payment = *payment
// 	// }

// 	// delivery, err := o.deliveryRepo.GetByID(ctx, id)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// if delivery != nil {
// 	// 	dataOrder.Delivery = *delivery
// 	// }

// 	// item, err := o.itemRepo.GetByID(ctx, id)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// if item != nil {
// 	// 	dataOrder.Item = *item
// 	// }

// 	return &dt, nil
// }

// func (o *orderService) CreateOrder(ctx context.Context, id string, order Order) error {

// 	// err := o.orderCache.Create(ctx, id, or)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	err := o.orderRepo.Create(ctx, id, order.Order)
// 	if err != nil {
// 		return err
// 	}

// 	err = o.deliveryRepo.Create(ctx, id, order.Delivery)
// 	if err != nil {
// 		return err
// 	}

// 	err = o.itemRepo.Create(ctx, id, order.Item[0])
// 	if err != nil {
// 		return err
// 	}

// 	err = o.paymentRepo.Create(ctx, id, order.Payment)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
