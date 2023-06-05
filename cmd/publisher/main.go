package main

import (
	"github.com/Enthreeka/L0/internal/config"
	"github.com/Enthreeka/L0/internal/controller/amqp"
	"github.com/Enthreeka/L0/pkg/logger"
	"github.com/nats-io/stan.go"
)

func main() {

	configPath := "configs/config.json"

	log := logger.New()

	config, err := config.New(configPath)
	if err != nil {
		log.Fatal("Failed to load config: %s", err)
	}

	//Подключение к брокеру со стороны отправителя
	stanConn, err := stan.Connect(config.Nats.ClusterID, config.Nats.PublisherID)
	if err != nil {
		log.Error("failed to connect to stan %s:", err)
	}

	log.Info("Starting subscriber: clusterID:%s,PublisherID:%s", config.Nats.ClusterID, config.Nats.PublisherID)

	defer stanConn.Close()

	//Cлой обработчика со стороны брокера
	stream := amqp.NewPublish(stanConn, log)

	err = stream.Publish(config.Nats.Subject)
	if err != nil {
		log.Error("%s:", err)
	}
}
