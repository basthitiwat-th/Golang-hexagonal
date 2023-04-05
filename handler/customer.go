package handler

import (
	"app/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	custSrv service.CustomerService
}

func NewCustomerHandler(custSrv service.CustomerService) CustomerHandler {
	return CustomerHandler{custSrv: custSrv}
}

func (h CustomerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customer, err := h.custSrv.GetCustomes()
	if err != nil {
		handlerError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customer)

}

func (h CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	customer, err := h.custSrv.GetCustomer(customerID)
	if err != nil {
		handlerError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customer)

}
