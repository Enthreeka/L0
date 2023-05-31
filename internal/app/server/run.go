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
)

func Run(log *logger.Logger, config *config.Config) error {

	db, err := db.NewConnect(context.Background(), config.Db.URL)
	if err != nil {
		return err
	}

	defer db.Close()

	app := fiber.New()

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
	orderHandler := http.NewOrderHandler(orderService, deliveryService, itemService, paymentService)

	api := app.Group("/api")

	api.Get("/:id", orderHandler.SearchOrder)

	log.Info("Starting http server: %s:%s", config.Server.TypeServer, config.Server.Port)

	if err = app.Listen(fmt.Sprintf(":%s", config.Server.Port)); err != nil {
		log.Fatal("Server listening failed:%s", err)
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
	// 	ChrtID:      999123,
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

	// orderObj := usecase.Order{
	// 	Order:    order,
	// 	Payment:  payment,
	// 	Item:     []entity.Item{item},
	// 	Delivery: delivery,
	// }

	return nil
}
