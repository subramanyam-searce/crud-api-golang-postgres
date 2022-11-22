package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subramanyam-searce/banking/dbconnect"
	"github.com/subramanyam-searce/banking/typedefs"
)

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)

	updatedCustomer := dbconnect.UpdateCustomer(vars["id"], vars["key"], vars["value"])

	json.NewEncoder(w).Encode(map[string]*typedefs.Customer{
		"updated_customer": updatedCustomer,
	})
}
