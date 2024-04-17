package domain

type Address struct {
	Country      string `json:"country"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	HouseNumber  string `json:"houseNumber"`
	Complement   string `json:"complement"`
	Reference    string `json:"reference"`
	ZipCode      string `json:"zipCode"`
}

func NewAddress(
	country string,
	state string,
	city string,
	neighborhood string,
	street string,
	houseNumber string,
	complement string,
	reference string,
	zipCode string,
) *Address {
	return &Address{
		Country:      country,
		State:        state,
		City:         city,
		Neighborhood: neighborhood,
		Street:       street,
		HouseNumber:  houseNumber,
		Complement:   complement,
		Reference:    reference,
		ZipCode:      zipCode,
	}
}
