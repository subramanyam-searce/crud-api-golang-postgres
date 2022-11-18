package app

import (
	"net/http"

	"github.com/subramanyam-searce/banking/handlers"
)

type Route struct {
	path    string
	handler func(w http.ResponseWriter, r *http.Request)
	method  string
}

var Routes []Route = []Route{
	{"/api/customer", handlers.CreateCustomer, http.MethodPost},
	{"/api/customer", handlers.GetAllCustomers, http.MethodGet},
	{"/api/customer/{id}", handlers.GetCustomer, http.MethodGet},
	{"/api/customer/{id}", handlers.DeleteCustomer, http.MethodDelete},
	{"/api/customer/{id}", handlers.UpdateCustomer, http.MethodPut},
}
