package models

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
