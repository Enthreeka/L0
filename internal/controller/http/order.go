package http

import (
	"context"

	"github.com/Enthreeka/L0/internal/usecase"
	"github.com/Enthreeka/L0/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type orderHandler struct {
	orderService    usecase.Order
	deliveryService usecase.Delivery
	itemService     usecase.Item
	paymentService  usecase.Payment

	log *logger.Logger
}

func NewOrderHandler(orderService usecase.Order, deliveryService usecase.Delivery, itemService usecase.Item, paymentService usecase.Payment, log *logger.Logger) *orderHandler {
	return &orderHandler{
		orderService:    orderService,
		deliveryService: deliveryService,
		itemService:     itemService,
		paymentService:  paymentService,
		log:             log,
	}
}

func (o *orderHandler) SearchOrder(c *fiber.Ctx) error {

	if c.Method() == fiber.MethodPost {
		o.log.Info("Start search order")

		id := c.FormValue("id")
		o.log.Info("Get order with id: %s", id)

		order, err := o.orderService.GetByID(context.Background(), id)
		if err != nil {
			return err
		}

		payment, err := o.paymentService.GetByID(context.Background(), id)
		if err != nil {
			return err
		}

		delivery, err := o.deliveryService.GetByID(context.Background(), id)
		if err != nil {
			return err
		}

		item, err := o.itemService.GetByID(context.Background(), id)
		if err != nil {
			return err
		}

		o.log.Info("Search order completed successfully")
		return c.Render("index", fiber.Map{
			"Order":    order,
			"Payment":  payment,
			"Delivery": delivery,
			"Item":     item,
		})
	} else {
		o.log.Info("Get start page")
	}
	return c.Render("index", fiber.Map{})
}
