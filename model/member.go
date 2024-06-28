package model

import (
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Id          string `json:"id"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
}

func GenerateNewMember() Member {
	payload := Member{}
	email := gofakeit.Email()
	payload.Id = email
	payload.Email = email
	payload.DisplayName = gofakeit.Name()

	return payload
}
