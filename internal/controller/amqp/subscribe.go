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

func (p *subscribeBroker) Subscribe(subject string) {

	p.stan.Subscribe(subject, func(msg *stan.Msg) {
		fmt.Printf("%s\n", msg)

		data := &entity.Data{}

		err := json.Unmarshal(msg.Data, data)
		if err != nil {
			fmt.Printf("failed to unmarshal data: %s", err)
		}

		p.orderService.CreateOrder(context.Background(), data.Order.OrderUID, data.Order)

		for _, el := range data.Item {
			p.itemService.CreateItem(context.Background(), el.OrderUID, el)
		}

		p.deliveryService.CreateDelivery(context.Background(), data.Delivery.OrderUID, data.Delivery)

		p.paymentService.CreatePayment(context.Background(), data.Payment.OrderUID, data.Payment)

		//fmt.Println(data)
	})

}
