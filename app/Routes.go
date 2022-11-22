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
	{"/customer", handlers.CreateCustomer, http.MethodPost},
	{"/customer", handlers.GetAllCustomers, http.MethodGet},
	{"/customer/{id:[0-9]+}", handlers.GetCustomer, http.MethodGet},
	{"/customer/{id:[0-9]+}", handlers.DeleteCustomer, http.MethodDelete},
	{"/customer/{id:[0-9]+}/{key:(?:name|age|email)}/{value}", handlers.UpdateCustomer, http.MethodPut},
}
