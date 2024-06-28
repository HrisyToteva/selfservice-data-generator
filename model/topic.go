package model

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type Topic struct {
	gorm.Model
	Id             string    `json:"id"`
	CapabilityId   string    `json:"capabilityId"`
	KafkaClusterId string    `json:"kafkaClusterId"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Partitions     int       `json:"partitions"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"createdAt"`
	CreatedBy      string    `json:"createdBy"`
	ModifiedAt     time.Time `json:"modifiedAt"`
	ModifiedBy     string    `json:"ModifiedBy"`
	Retention      string    `json:"retention"`
}

func fakeTopicStatus() string {
	statusValues := []string{"Provisioned"} //, "Requested", "In Progress"
	statusIndex := rand.Intn(len(statusValues))
	return statusValues[statusIndex]
}

func GenerateNewTopic(capabilityId string) Topic {
	payload := Topic{}
	payload.Name = fmt.Sprintf("pub.%s", gofakeit.ProductName())
	payload.Id = gofakeit.UUID()
	payload.Description = gofakeit.ProductDescription()
	payload.CreatedAt = gofakeit.Date()
	payload.CreatedBy = gofakeit.Email()
	payload.Status = fakeTopicStatus()
	payload.ModifiedAt = gofakeit.Date()
	payload.ModifiedBy = gofakeit.Email()
	payload.Partitions = 3
	payload.Retention = "1d"
	payload.CapabilityId = capabilityId
	payload.KafkaClusterId = "prod1"

	return payload
}
