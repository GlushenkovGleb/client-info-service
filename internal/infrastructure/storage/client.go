package storage

import (
	"client-info-service/internal/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ClientStorage struct {
	db *sqlx.DB
}

func NewClientStorage(db *sqlx.DB) *ClientStorage {
	return &ClientStorage{db: db}
}

func (c *ClientStorage) SaveClient(info model.ClientInfo) error {
	_, err := c.db.NamedExec(
		`INSERT INTO clients(id, name, surname, patronymic, gender, age, country_id)
				VAlUES(:id, :name, :surname, :patronymic, :gender, :age, :country_id)`, info)
	return err
}

func (c *ClientStorage) GetClients() ([]model.ClientInfo, error) {
	var infos []model.ClientInfo
	err := c.db.Select(&infos, "SELECT * FROM clients")
	if err != nil {
		return []model.ClientInfo{}, err
	}
	return infos, nil
}

func (c *ClientStorage) DeleteClient(id uuid.UUID) error {
	_, err := c.db.Exec("DELETE FROM clients WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (c *ClientStorage) UpdateClient(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
