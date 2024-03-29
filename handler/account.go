package handler

import (
	"app/errs"
	"app/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type accountHandler struct {
	accSrv service.AccountService
}

func NewaccountHandler(accSrv service.AccountService) accountHandler {
	return accountHandler{accSrv: accSrv}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	CustomerId, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	if r.Header.Get("content-type") != "application/json" {
		handlerError(w, errs.NewValidatorError("request body incorrect format"))
		return
	}

	request := service.NewAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handlerError(w, errs.NewValidatorError("request body incorrect format"))
		return
	}

	response, err := h.accSrv.NewAccount(CustomerId, request)
	if err != nil {
		handlerError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h accountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	CustomerId, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	responses, err := h.accSrv.GetAccount(CustomerId)
	if err != nil {
		handlerError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(responses)

}
