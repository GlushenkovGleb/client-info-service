package main

import (
	"client-info-service/internal/config"
	appHTTP "client-info-service/internal/controller/httpcontroller"
	appKafka "client-info-service/internal/controller/kafkaconsumer"
	"client-info-service/internal/infrastructure/producer"
	"client-info-service/internal/infrastructure/storage"
	"client-info-service/internal/infrastructure/webapi"
	"client-info-service/internal/service"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

/*
1. init config
2.  init logger
3. init router
4. start router
5. init producer
*/

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
	log.Info("Starting client-info-service...")
	log.Debug(fmt.Sprintf("config: %v", cfg.Kafka))

	db, err := storage.NewPostgresDB(cfg.DataBase)
	if err != nil {
		log.Error("Could not init database: %s", err)
		return
	}

	st := storage.NewStorage(db)
	enricher := webapi.NewClientInfoEnricher(cfg.ClientInfoEnricher)
	prod := producer.NewKafkaProducer(cfg.Kafka)
	s := service.New(st, enricher, prod)

	consumer := appKafka.NewConsumer(cfg.Kafka, log, s)
	go consumer.ProcessAddClientsQueue()

	handler := appHTTP.NewClientController(log, s)
	r := appHTTP.NewRouter(handler)
	http.ListenAndServe(cfg.HTTPServer.Address, r)

}
