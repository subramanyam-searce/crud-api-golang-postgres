package handlers

import (
	"encoding/json"
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
		json.NewEncoder(w).Encode(err)
	} else {
		customer := dbconnect.GetCustomer(id)
		json.NewEncoder(w).Encode(customer)
	}

}
