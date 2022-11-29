package handlers

import (
	"net/http"

	"github.com/subramanyam-searce/banking/dbconnect"
)

func ViewAllTransactions(w http.ResponseWriter, r *http.Request) {
	transactions := dbconnect.GetAllTransactions()

	SendResponse(w, r, transactions)
}
