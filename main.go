package main

import (
	"encoding/json"
	"fmt"
	"modules/models"
)

func main() {

	product, _ := json.Marshal(
		&models.Product{
			Id:           123456789,
			Name:         "Test Product",
			Description:  "Nice Product",
			Category:     "New",
			Price:        90.5,
			Discount:     30.5,
			Availability: 12,
		},
	)

	fmt.Println(
		"Product:",
		string(product),
	)
}
