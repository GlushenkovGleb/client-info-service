package kafkaconsumer

import (
	"client-info-service/internal/config"
	"client-info-service/internal/model"
	"client-info-service/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log/slog"
)

type Consumer struct {
	log    *slog.Logger
	reader *kafka.Reader
	s      *service.Service
}

func NewConsumer(cfg config.Kafka, log *slog.Logger, s *service.Service) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{cfg.Address},
		Topic:     cfg.AddClientTopic,
		GroupID:   cfg.GroupId,
		Partition: cfg.Partition,
		MaxBytes:  cfg.MaxBytes,
	})
	return &Consumer{log, reader, s}
}

func (c *Consumer) ProcessAddClientsQueue() {
	// TODO: implement graceful shutdown
	for {
		msg, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		var client model.Client
		err = json.Unmarshal(msg.Value, &client)
		if err != nil {
			c.log.Error(fmt.Sprintf("Could not deserialize: %v", err))
			return
		}
		c.log.Debug(fmt.Sprintf("I got this client: %v", client))
		// process message
		err = c.s.AddClient(client)

		if err != nil {
			c.log.Error(fmt.Sprintf("Could not add client: %v", err))
		}
	}
	if err := c.reader.Close(); err != nil {
		c.log.Error("failed to close reader: ", err)
	}
}
