package amqp

import (
	"encoding/json"
	"fmt"

	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/pkg/logger"
	"github.com/brianvoe/gofakeit/v6"
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

func (p *publishBroker) Publish(subject string) error {

	p.log.Info("Publish the data")

	order := createData()
	o, _ := json.Marshal(order)

	err := p.stan.Publish(subject, o)

	if err != nil {
		p.log.Error("failed with publish method %v:", err)
		return err
	}

	p.log.Info("Data publish successfully")

	return nil
}

func createData() *entity.Data {

	UUID := fmt.Sprintf("b563eb7b2b84b6test%v", gofakeit.Number(1, 100))

	fmt.Println(UUID, "- UUID")

	order := entity.Order{
		OrderUID:          UUID,
		TrackNumber:       gofakeit.Word(),
		Entry:             gofakeit.Word(),
		Locale:            "ru",
		InternalSignature: "",
		CustomerID:        gofakeit.UUID(),
		DeliveryService:   gofakeit.Word(),
		ShardKey:          gofakeit.Word(),
		SmID:              gofakeit.Number(1, 100),
		DateCreated:       gofakeit.Date(),
		OofShard:          gofakeit.Word(),
	}

	payment := entity.Payment{
		OrderUID:     UUID,
		Transaction:  UUID,
		RequestID:    "",
		Currency:     gofakeit.Currency().Short,
		Provider:     gofakeit.Word(),
		Amount:       gofakeit.Number(1, 100),
		PaymentDt:    gofakeit.Number(100000, 10000000),
		Bank:         gofakeit.Word(),
		DeliveryCost: gofakeit.Number(100, 9999),
		GoodsTotal:   gofakeit.Number(100, 1000),
		CustomFee:    gofakeit.Number(0, 10),
	}

	item := []entity.Item{
		{
			OrderUID:    UUID,
			ChrtID:      gofakeit.Number(900, 1000),
			TrackNumber: gofakeit.Word(),
			Price:       gofakeit.Number(0, 99999),
			RID:         gofakeit.Word(),
			Name:        gofakeit.Name(),
			Sale:        gofakeit.Number(0, 100),
			Size:        string(gofakeit.Number(1, 5)),
			TotalPrice:  gofakeit.Number(0, 9999),
			NmID:        gofakeit.Number(100000, 200000),
			Brand:       gofakeit.Word(),
			Status:      gofakeit.Number(200, 500),
		},
		{
			OrderUID:    UUID,
			ChrtID:      gofakeit.Number(900, 1000),
			TrackNumber: gofakeit.Word(),
			Price:       gofakeit.Number(0, 99999),
			RID:         gofakeit.Word(),
			Name:        gofakeit.Name(),
			Sale:        gofakeit.Number(0, 100),
			Size:        string(gofakeit.Number(1, 5)),
			TotalPrice:  gofakeit.Number(0, 9999),
			NmID:        gofakeit.Number(100000, 200000),
			Brand:       gofakeit.Word(),
			Status:      gofakeit.Number(200, 500),
		},
	}

	delivery := entity.Delivery{
		OrderUID: UUID,
		Name:     gofakeit.Name(),
		Phone:    gofakeit.Phone(),
		Zip:      gofakeit.Zip(),
		City:     gofakeit.City(),
		Address:  gofakeit.Address().City,
		Region:   gofakeit.State(),
		Email:    gofakeit.Email(),
	}

	orderObj := &entity.Data{
		Order:    order,
		Payment:  payment,
		Delivery: delivery,
		Item:     item,
	}

	return orderObj
}
