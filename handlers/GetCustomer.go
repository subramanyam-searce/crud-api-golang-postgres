package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/subramanyam-searce/banking/dbconnect"
)

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("strconvAtoiError:", err)
		SendResponse(w, r, err)
	} else {
		customer := dbconnect.GetCustomer(strconv.Itoa(id))
		SendResponse(w, r, customer)
	}
}
