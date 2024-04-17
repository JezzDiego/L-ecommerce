package models

type Order struct {
	IdOrder       int     `json:"id"`
	IdUser        string  `json:"id_user"`
	OrderStatus   string  `json:"order_status"`
	PaymentMethod string  `json:"payment_method"`
	TotalValue    float32 `json:"total_value"`
	ArriveDate    string  `json:"arrive_date"`
	ArriveAddress Address `json:"arrive_address"`
	Message       string  `json:"message"`
}
