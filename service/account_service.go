package service

import (
	"app/errs"
	"app/logs"
	"app/repository"
	"strings"
	"time"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService {
	return accountService{accRepo: accRepo}
}

func (s accountService) NewAccount(customerID int, request NewAccountRequest) (*AccountResponse, error) {
	//validate input
	if request.Amount < 5000 {
		return nil, errs.NewValidatorError("amount at lest 5,000")
	}
	if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
		return nil, errs.NewValidatorError("account type should bet saving or checking")
	}

	account := repository.Account{
		CustomerId:  customerID,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      1,
	}
	newAcc, err := s.accRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	responses := AccountResponse{
		AccountId:   newAcc.AccountId,
		OpeningDate: newAcc.OpeningDate,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}
	return &responses, nil
}

func (s accountService) GetAccount(customerID int) ([]AccountResponse, error) {
	accounts, err := s.accRepo.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	responses := []AccountResponse{}
	for _, account := range accounts {
		responses = append(responses, AccountResponse{
			AccountId:   account.AccountId,
			OpeningDate: account.OpeningDate,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      account.Status,
		})
	}
	return responses, nil
}
