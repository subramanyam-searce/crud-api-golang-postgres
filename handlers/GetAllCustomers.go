package handlers

import (
	"net/http"

	"github.com/subramanyam-searce/banking/dbconnect"
)

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := dbconnect.GetAllCustomers()

	SendResponse(w, r, customers)
}
