package amqp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/internal/usecase"
	"github.com/Enthreeka/L0/pkg/logger"
	"github.com/nats-io/stan.go"
)

type subscribeBroker struct {
	stan stan.Conn
	log  *logger.Logger

	orderService    usecase.Order
	deliveryService usecase.Delivery
	itemService     usecase.Item
	paymentService  usecase.Payment
}

func NewSubcribe(stan stan.Conn, log *logger.Logger, orderService usecase.Order, deliveryService usecase.Delivery, itemService usecase.Item, paymentService usecase.Payment) Subscribe {
	return &subscribeBroker{
		stan:            stan,
		log:             log,
		orderService:    orderService,
		deliveryService: deliveryService,
		itemService:     itemService,
		paymentService:  paymentService,
	}
}

func (p *subscribeBroker) Subscribe(subject string) error {

	_, err := p.stan.Subscribe(subject, func(msg *stan.Msg) {
		p.log.Info("get data in subject:%s from publisher", subject)

		data := &entity.Data{}

		err := json.Unmarshal(msg.Data, data)
		if err != nil {
			fmt.Printf("failed to unmarshal data: %s", err)
		}

		//Можно добавить регулярные выражения для валидации входящих данных

		err = p.orderService.CreateOrder(context.Background(), data.Order.OrderUID, data.Order)
		if err != nil {
			p.log.Error("failed to get order in %s from publisher %s:", subject, err)
		}

		for _, el := range data.Item {
			p.itemService.CreateItem(context.Background(), el.OrderUID, el)
		}

		p.deliveryService.CreateDelivery(context.Background(), data.Delivery.OrderUID, data.Delivery)

		p.paymentService.CreatePayment(context.Background(), data.Payment.OrderUID, data.Payment)

		err = p.checkByID(context.Background(), data.Order.OrderUID)
		if err != nil {
			p.log.Error("error in the transaction %s:", err)
		}
	})
	if err != nil {
		p.log.Error("failed with subscribe method %v:", err)
		return err
	}
	return nil
}

// Костыльная транзакция.
// В случае, если кортеж хотя бы в одной из таблиц не создался, то кортежи в других таблицах с этим id удаляются.
// В случае заказа с 2 item, если один удовлетворяет требованиям, а другой нет, то транзакция не будет применена.
func (p *subscribeBroker) checkByID(ctx context.Context, id string) error {

	countErr := 0

	if _, err := p.orderService.GetByID(ctx, id); err != nil {
		countErr++
	}

	if _, err := p.paymentService.GetByID(ctx, id); err != nil {
		countErr++
	}

	if _, err := p.itemService.GetByID(ctx, id); err != nil {
		countErr++
	}

	if _, err := p.deliveryService.GetByID(ctx, id); err != nil {
		countErr++
	}

	if countErr > 0 {

		if err := p.deliveryService.DeleteByID(ctx, id); err != nil {
			return err
		}

		if err := p.itemService.DeleteByID(ctx, id); err != nil {
			return err
		}

		if err := p.paymentService.DeleteByID(ctx, id); err != nil {
			return err
		}

		if err := p.orderService.DeleteByID(ctx, id); err != nil {
			return err
		}

		p.log.Info("Delete all filed in interconnected tables with problems")
	}

	return nil
}
