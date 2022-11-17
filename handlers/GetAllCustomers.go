package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/subramanyam-searce/banking/dbconnect"
)

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := dbconnect.GetAllCustomers()

	json.NewEncoder(w).Encode(customers)
}
