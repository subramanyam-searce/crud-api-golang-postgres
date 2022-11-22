package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subramanyam-searce/banking/dbconnect"
	"github.com/subramanyam-searce/banking/typedefs"
)

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	foundCustomer := dbconnect.GetCustomer(id)

	if foundCustomer != nil {
		err := dbconnect.DeleteCustomer(id)
		if err != nil {
			SendResponse(w, r, err)
		} else {
			SendResponse(w, r, map[string]typedefs.Customer{"data_deleted": *foundCustomer})
		}
	} else {
		SendResponse(w, r, map[string]string{"message": "User with id " + vars["id"] + " does not exist!"})
	}
}
