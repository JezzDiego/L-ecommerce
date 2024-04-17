package models

import (
	"time"
)

type User struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	CPF         string    `json:"cpf"`
	BirthDate   time.Time `json:"birthDate"`
	PhoneNumber []string  `json:"phoneNumber"`
	Address     []Address `json:"address"`
}
