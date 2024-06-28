package model

import (
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
	"time"
)

type Ecr struct {
	gorm.Model
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	CreatedBy   string    `json:"createdBy"`
	Description string    `json:"description"`
	RequestedAt time.Time `json:"requestedAt"`
}

func GenerateNewEcr() Ecr {
	payload := Ecr{}
	payload.Id = gofakeit.UUID()
	payload.Name = gofakeit.ProductName()
	payload.CreatedBy = gofakeit.Email()
	payload.RequestedAt = gofakeit.Date()

	return payload
}
