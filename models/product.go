package models

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float32 `json:"price"`
	Discount     float32 `json:"discount"`
	Availability int     `json:"availability"`
	Category     string  `json:"category"`
}
