package amqp

import (
	"encoding/json"
	"time"

	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/pkg/logger"
	"github.com/nats-io/stan.go"
)

type publishBroker struct {
	stan stan.Conn

	log *logger.Logger
}

func NewPublish(stan stan.Conn, log *logger.Logger) Publish {
	return &publishBroker{
		stan: stan,
		log:  log,
	}
}

func (p *publishBroker) Publish(subject string) {

	p.log.Info("Publish the data")

	order := createData()
	o, _ := json.Marshal(order)

	p.stan.Publish(subject, o)

	p.log.Info("Data publish successfully")
}

func createData() *entity.Data {

	order := entity.Order{
		OrderUID:          "SAS",
		TrackNumber:       "WBILMTESTTRACK",
		Entry:             "WBIL",
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        "test",
		DeliveryService:   "meest",
		ShardKey:          "9",
		SmID:              99,
		DateCreated:       time.Now(),
		OofShard:          "1",
	}

	payment := entity.Payment{
		OrderUID:     "SAS",
		Transaction:  "b563feb7b2b84b6test",
		RequestID:    "",
		Currency:     "USD",
		Provider:     "wbpay",
		Amount:       1817,
		PaymentDt:    1637907727,
		Bank:         "alpg",
		DeliveryCost: 1500,
		GoodsTotal:   317,
		CustomFee:    0,
	}

	item := []entity.Item{
		{
			OrderUID:    "SAS",
			ChrtID:      909,
			TrackNumber: "WBILMTESTTRACK",
			Price:       343,
			RID:         "ab4219fsd087a764ae0btest",
			Name:        "ilya",
			Sale:        30,
			Size:        "0",
			TotalPrice:  316,
			NmID:        141241,
			Brand:       "versache",
			Status:      202,
		}, {
			OrderUID:    "SA",
			ChrtID:      910,
			TrackNumber: "WBILMTESTTRACK",
			Price:       343,
			RID:         "ab4219fsd087a764ae0btest",
			Name:        "ilya",
			Sale:        30,
			Size:        "0",
			TotalPrice:  316,
			NmID:        141241,
			Brand:       "versache",
			Status:      202,
		},
	}

	delivery := entity.Delivery{
		OrderUID: "SAS",
		Name:     "test",
		Phone:    "+9720000001",
		Zip:      "2639809",
		City:     "nino",
		Address:  "Ploshad Mira 15",
		Region:   "Kraiot",
		Email:    "test@gmail.com",
	}

	orderObj := &entity.Data{
		Order:    order,
		Payment:  payment,
		Delivery: delivery,
		Item:     item,
	}

	return orderObj
}
