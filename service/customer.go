package service

type CustomerResponse struct {
	CustomerId int    `json:"customer_id"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
}

type CustomerServices interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(int) (*CustomerResponse, error)
}
