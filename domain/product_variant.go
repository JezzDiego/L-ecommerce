package domain

type ProductVariant struct {
	IdProduct int `json:"id_product"`
	IdVariant int `json:"id_variant"`
}

func NewProductVariant(
	idProduct int,
	idVariant int,
) *ProductVariant {
	return &ProductVariant{
		IdProduct: idProduct,
		IdVariant: idVariant,
	}
}
