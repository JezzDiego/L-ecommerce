package domain

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float32 `json:"price"`
	Discount     float32 `json:"discount"`
	Availability int     `json:"availability"`
	Category     string  `json:"category"`
}

func NewProduct(
	id int,
	name string,
	description string,
	prince float32,
	discount float32,
	availability int,
	category string,
) *Product {
	return &Product{
		Id:           id,
		Name:         name,
		Description:  description,
		Price:        prince,
		Discount:     discount,
		Availability: availability,
		Category:     category,
	}
}
