package service

import (
	"app/errs"
	"app/logs"
	"app/repository"
	"database/sql"
)

type CustomerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) CustomerService {
	return CustomerService{custRepo: custRepo}
}

func (s CustomerService) GetCustomes() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	CustomerResponses := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{
			CustomerId: customer.CustomerId,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		CustomerResponses = append(CustomerResponses, custResponse)
	}
	return CustomerResponses, nil
}

func (s CustomerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		return nil, errs.NewUnexpectedError()
	}

	custResponse := CustomerResponse{
		CustomerId: customer.CustomerId,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &custResponse, nil
}
