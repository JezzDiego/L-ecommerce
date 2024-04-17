package domain

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

func NewOrder(
	idOrder int,
	idUser string,
	orderStatus string,
	paymentMethod string,
	totalValue float32,
	arriveDate string,
	arriveAddress Address,
	message string,
) *Order {
	return &Order{
		IdOrder:       idOrder,
		IdUser:        idUser,
		OrderStatus:   orderStatus,
		PaymentMethod: paymentMethod,
		TotalValue:    totalValue,
		ArriveDate:    arriveDate,
		ArriveAddress: arriveAddress,
		Message:       message,
	}
}
