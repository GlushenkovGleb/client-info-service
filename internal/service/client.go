package service

import (
	"client-info-service/internal/infrastructure/producer"
	"client-info-service/internal/infrastructure/storage"
	"client-info-service/internal/infrastructure/webapi"
	"client-info-service/internal/model"
	"fmt"
	"github.com/google/uuid"
)

type ClientInfoService struct {
	st       *storage.Storage
	enricher *webapi.ClientInfoEnricher
	prod     *producer.KafkaProducer
}

func newClientInfoService(st *storage.Storage, enr *webapi.ClientInfoEnricher, prod *producer.KafkaProducer) *ClientInfoService {
	return &ClientInfoService{st: st, enricher: enr, prod: prod}
}

func (c *ClientInfoService) PutClientInQueue(client model.Client) (uuid.UUID, error) {
	client.Id = uuid.New()

	err := c.prod.AddClient(client)
	if err != nil {
		return [16]byte{}, err
	}

	return client.Id, nil
}

func (c *ClientInfoService) AddClient(client model.Client) error {
	clientInfo, err := c.enricher.Enrich(client)
	if err != nil {
		return err
	}

	err = c.st.SaveClient(clientInfo)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (c *ClientInfoService) GetClients() ([]model.ClientInfo, error) {
	infos, err := c.st.GetClients()
	if err != nil {
		return []model.ClientInfo{}, err
	}

	return infos, nil
}

func (c *ClientInfoService) DeleteClientById(id uuid.UUID) error {
	err := c.st.DeleteClient(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *ClientInfoService) UpdateClient(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
