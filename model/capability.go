package model

import (
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
	"time"
)

type Capability struct {
	gorm.Model
	Id                        string    `json:"id"`
	Name                      string    `json:"name"`
	Description               string    `json:"description"`
	CreatedAt                 time.Time `json:"createdAt"`
	CreatedBy                 string    `json:"createdBy"`
	Status                    string    `json:"status"`
	ModifiedAt                time.Time `json:"modifiedAt"`
	ModifiedBy                string    `json:"modifiedBy"`
	JsonMetadata              string    `json:"jsonMetadata"`
	JsonMetadataSchemaVersion int       `json:"JsonMetadataSchemaVersion"`
}

func GenerateNewCapability() Capability {
	payload := Capability{}
	payload.Name = gofakeit.ProductName()
	payload.Id = gofakeit.UUID()
	payload.Description = gofakeit.ProductDescription()
	payload.CreatedAt = gofakeit.Date()
	payload.CreatedBy = gofakeit.Email()
	payload.Status = "Active"
	payload.ModifiedAt = gofakeit.Date()
	payload.ModifiedBy = gofakeit.Email()
	payload.JsonMetadata = "{}"
	payload.JsonMetadataSchemaVersion = 1
	return payload
}
