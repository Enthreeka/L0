package server

import (
	"context"
	"fmt"
	"time"

	"github.com/Enthreeka/L0/internal/config"
	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/internal/repo/cache"
	"github.com/Enthreeka/L0/internal/repo/postgres"
	"github.com/Enthreeka/L0/pkg/db"
	"github.com/Enthreeka/L0/pkg/logger"
)

func Run(log *logger.Logger, config *config.Config) error {

	db, err := db.NewConnect(context.Background(), config.Db.URL)
	if err != nil {
		return err
	}

	orderRepo := postgres.NewOrderRepository(db)

	//cache := make(map[string]entity.Order)

	//orderCache := cache.NewOrderCache(cache)

	order := entity.Order{
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

	err = orderRepo.Create(context.Background(), "b63feb7b2b84b", order)
	if err != nil {
		return err
	}

	c := make(map[string]entity.Order)

	cache := cache.NewOrderCache(c)

	cache.Create(context.TODO(), "b63feb7b2b84b", order)

	cache.DeleteByID(context.TODO(), "b63feb7b2b84b")

	or, err := cache.GetByID(context.TODO(), "b63feb7b2b84b")
	if err != nil {
		return err
	}

	fmt.Println(or)

	return nil
}
