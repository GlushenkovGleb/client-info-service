package producer

import (
	"client-info-service/internal/config"
	"client-info-service/internal/model"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	addClientWriter *kafka.Writer
}

func NewKafkaProducer(cfg config.Kafka) *KafkaProducer {
	addClientWriter := &kafka.Writer{
		Addr:  kafka.TCP(cfg.Address),
		Topic: cfg.AddClientTopic,
	}
	return &KafkaProducer{addClientWriter: addClientWriter}
}

func (p *KafkaProducer) AddClient(cl model.Client) error {
	clBytes, _ := json.Marshal(cl)
	err := p.addClientWriter.WriteMessages(context.Background(),
		kafka.Message{
			Value: clBytes,
		},
	)
	return err
}
