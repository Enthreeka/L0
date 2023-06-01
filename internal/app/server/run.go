package server

import (
	"context"
	"fmt"
	"sync"

	"github.com/Enthreeka/L0/internal/config"
	"github.com/Enthreeka/L0/internal/controller/http"
	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/internal/repo/cache"
	"github.com/Enthreeka/L0/internal/repo/postgres"
	"github.com/Enthreeka/L0/internal/usecase"
	"github.com/Enthreeka/L0/pkg/db"
	"github.com/Enthreeka/L0/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func Run(log *logger.Logger, config *config.Config) error {

	db, err := db.NewConnect(context.Background(), config.Db.URL)
	if err != nil {
		return err
	}

	defer db.Close()

	engine := html.New("./public", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	order := make(map[string]entity.Order)
	payment := make(map[string]entity.Payment)
	delivery := make(map[string]entity.Delivery)
	item := make(map[string][]entity.Item)

	// Слой репозитория БД
	orderDB := postgres.NewOrderRepository(db)
	paymentDB := postgres.NewPaymentRepository(db)
	deliveryDB := postgres.NewDeliveryRepository(db)
	itemDB := postgres.NewItemRepo(db)

	// Слой репозитория кеша
	orderCache := cache.NewOrderCache(order)
	paymentCache := cache.NewPaymentCache(payment)
	deliveryCache := cache.NewDeliveryCache(delivery)
	itemCache := cache.NewItemCache(item)

	// Слой бизнес-логики
	orderService := usecase.NewOrderService(orderDB, orderCache, log)
	paymentService := usecase.NewPaymentService(paymentDB, paymentCache, log)
	deliveryService := usecase.NewDeliveryService(deliveryDB, deliveryCache, log)
	itemService := usecase.NewItemService(itemDB, itemCache, log)

	//Загрузка всех данных из бд в кеш при запуске сервера
	getFromDB := func() {
		err := orderService.SaveAllToCache(context.Background())
		if err != nil {
			log.Error("Error saving orders to cache:", err)
		}
		err = deliveryService.SaveAllToCache(context.Background())
		if err != nil {
			log.Error("Error saving delivery to cache:", err)
		}
		err = paymentService.SaveAllToCache(context.Background())
		if err != nil {
			log.Error("Error saving payments to cache:", err)
		}
		err = itemService.SaveAllToCache(context.Background())
		if err != nil {
			log.Error("Error saving items to cache:", err)
		}
	}

	var once sync.Once
	once.Do(getFromDB)

	// Слой обработчика
	orderHandler := http.NewOrderHandler(orderService, deliveryService, itemService, paymentService, log)

	api := app.Group("/api")

	api.Get("/", orderHandler.SearchOrder)
	api.Post("/search", orderHandler.SearchOrder)

	log.Info("Starting http server: %s:%s", config.Server.TypeServer, config.Server.Port)

	if err = app.Listen(fmt.Sprintf(":%s", config.Server.Port)); err != nil {
		log.Fatal("Server listening failed:%s", err)
	}

	return nil
}
