package user

import (
	Address2 "modules/domain/address"
	"time"
)

type User struct {
	Name        string             `json:"name"`
	CPF         string             `json:"cpf"`
	BirthDate   time.Time          `json:"birthDate"`
	PhoneNumber []string           `json:"phoneNumber"`
	Address     []Address2.Address `json:"address"`
}

func NewUser(name string, CPF string, birthDate time.Time, phoneNumber []string, address []Address2.Address) *User {
	return &User{Name: name, CPF: CPF, BirthDate: birthDate, PhoneNumber: phoneNumber, Address: address}
}
