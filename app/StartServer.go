package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subramanyam-searce/banking/handlers"
)

func StartServer() {
	router := mux.NewRouter()

	//POST - Create Users
	router.HandleFunc("/api/customer", handlers.CreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/api/customer", handlers.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/api/customer/{id}", handlers.GetCustomer).Methods(http.MethodGet)
	router.HandleFunc("/api/customer/{id}", handlers.DeleteCustomer).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", router))
}
