package main

import (
	"encoding/json"
	"fmt"
	"modules/domain"
	"time"
)

func main() {
	address := domain.NewAddress(
		"Brazil",
		"SP",
		"São Paulo",
		"Vila Mariana",
		"Rua Domingos de Morais",
		"123",
		"Apto 123",
		"teste",
		"04007-010",
	)
	user := domain.NewUser(
		"John Doe",
		"123.456.789-00",
		time.Now(),
		[]string{"+55 11 99999-9999"},
		[]domain.Address{*address},
	)

	userObj, _ := json.Marshal(user)

	fmt.Println(
		"User:",
		string(userObj),
	)
}
