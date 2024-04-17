package models

import (
	"time"
)

type User struct {
	Name        string    `json:"name"`
	CPF         string    `json:"cpf"`
	BirthDate   time.Time `json:"birthDate"`
	PhoneNumber []string  `json:"phoneNumber"`
	Address     []Address `json:"address"`
}
