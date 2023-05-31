package http

import (
	"context"
	"fmt"

	"github.com/Enthreeka/L0/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type orderHandler struct {
	orderService    usecase.Order
	deliveryService usecase.Delivery
	itemService     usecase.Item
	paymentService  usecase.Payment
}

func NewOrderHandler(orderService usecase.Order, deliveryService usecase.Delivery, itemService usecase.Item, paymentService usecase.Payment) *orderHandler {
	return &orderHandler{
		orderService:    orderService,
		deliveryService: deliveryService,
		itemService:     itemService,
		paymentService:  paymentService,
	}
}

func (o *orderHandler) SearchOrder(c *fiber.Ctx) error {

	id := c.FormValue("id")
	fmt.Println(id)

	order, err := o.orderService.GetByID(context.Background(), "21")
	if err != nil {
		return err
	}

	payment, err := o.paymentService.GetByID(context.Background(), "21")
	if err != nil {
		return err
	}

	delivery, err := o.deliveryService.GetByID(context.Background(), "21")
	if err != nil {
		return err
	}

	item, err := o.itemService.GetByID(context.Background(), "21")
	if err != nil {
		return err
	}

	return c.Render("index", fiber.Map{
		"Order":    order,
		"Payment":  payment,
		"Delivery": delivery,
		"Item":     item,
	})
}

// payment := entity.Payment{
// 	Transaction:  "b563feb7b2b84b6test",
// 	RequestID:    "",
// 	Currency:     "USD",
// 	Provider:     "wbpay",
// 	Amount:       1817,
// 	PaymentDt:    1637907727,
// 	Bank:         "alpg",
// 	DeliveryCost: 1500,
// 	GoodsTotal:   317,
// 	CustomFee:    0,
// }

// item := entity.Item{
// 	ChrtID:      99912,
// 	TrackNumber: "WBILMTESTTRACK",
// 	Price:       343,
// 	RID:         "ab4219fsd087a764ae0btest",
// 	Name:        "ilya",
// 	Sale:        30,
// 	Size:        "0",
// 	TotalPrice:  316,
// 	NmID:        141241,
// 	Brand:       "versache",
// 	Status:      202,
// }

// delivery := entity.Delivery{
// 	Name:    "test",
// 	Phone:   "+9720000000",
// 	Zip:     "2639809",
// 	City:    "nino",
// 	Address: "Ploshad Mira 15",
// 	Region:  "Kraiot",
// 	Email:   "test@gmail.com",
// }

// order := entity.Order{
// 	TrackNumber:       "WBILMTESTTRACK",
// 	Entry:             "WBIL",
// 	Locale:            "en",
// 	InternalSignature: "",
// 	CustomerID:        "test",
// 	DeliveryService:   "meest",
// 	ShardKey:          "9",
// 	SmID:              99,
// 	DateCreated:       time.Now(),
// 	OofShard:          "1",
// }
