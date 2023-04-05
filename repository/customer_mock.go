package repository

import "errors"

type CustomerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() CustomerRepositoryMock {
	customers := []Customer{
		{CustomerId: 1001, Name: "bas bm", City: "BKK", ZipCode: "12356", Status: 1},
		{CustomerId: 1002, Name: "bas bmz", City: "II", ZipCode: "44", Status: 4},
		{CustomerId: 1003, Name: "bas bm2", City: "AB", ZipCode: "123", Status: 5},
	}

	return CustomerRepositoryMock{customers: customers}

}

func (r CustomerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r CustomerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerId == id {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}
