package model

import (
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
	"time"
)

type Membership struct {
	gorm.Model
	Id           string    `json:"id"`
	UserId       string    `json:"userId"`
	CapabilityId string    `json:"capabilityId"`
	CreatedAt    time.Time `json:"capabilityId"`
}

func GenerateNewMembership(userId string, capabilityId string) Membership {
	payload := Membership{}
	payload.Id = gofakeit.UUID()
	payload.CapabilityId = capabilityId
	payload.UserId = userId
	payload.CreatedAt = time.Now()

	return payload
}
