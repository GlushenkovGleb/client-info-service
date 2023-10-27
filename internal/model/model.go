package model

import "github.com/google/uuid"

type ClientInfo struct {
	Client
	Age       int    `db:"age" json:"age"`
	Gender    string `db:"gender" json:"gender"`
	CountryId string `db:"country_id" json:"countryId"`
}

type Client struct {
	Id         uuid.UUID `json:"id" validate:"omitempty,uuid" db:"id"`
	Name       string    `json:"name" validate:"required" db:"name"`
	Surname    string    `json:"surname" validate:"required" db:"surname"`
	Patronymic string    `json:"patronymic,omitempty" db:"patronymic"`
}
