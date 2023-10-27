package service

import (
	"client-info-service/internal/infrastructure/producer"
	"client-info-service/internal/infrastructure/storage"
	"client-info-service/internal/infrastructure/webapi"
	"client-info-service/internal/model"
	"github.com/google/uuid"
)

type ClientInfo interface {
	PutClientInQueue(client model.Client) (uuid.UUID, error)
	AddClient(client model.Client) error
	GetClients() ([]model.ClientInfo, error)
	DeleteClientById(id uuid.UUID) error
	UpdateClient(id uuid.UUID) error
}

type Service struct {
	ClientInfo
}

func New(st *storage.Storage, enr *webapi.ClientInfoEnricher, prod *producer.KafkaProducer) *Service {
	clientInfo := newClientInfoService(st, enr, prod)
	return &Service{ClientInfo: clientInfo}
}
