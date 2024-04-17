package domain

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

func NewUser(
	name string,
	CPF string,
	birthDate time.Time,
	phoneNumber []string,
	address []Address,
) *User {
	return &User{
		Name:        name,
		CPF:         CPF,
		BirthDate:   birthDate,
		PhoneNumber: phoneNumber,
		Address:     address,
	}
}
