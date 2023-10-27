package storage

import (
	"client-info-service/internal/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Client interface {
	SaveClient(info model.ClientInfo) error
	GetClients() ([]model.ClientInfo, error)
	DeleteClient(id uuid.UUID) error
	UpdateClient(id uuid.UUID) error
}

type Storage struct {
	Client
}

func NewStorage(db *sqlx.DB) *Storage {
	client := NewClientStorage(db)
	return &Storage{client}
}
