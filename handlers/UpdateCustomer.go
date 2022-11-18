package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/subramanyam-searce/banking/dbconnect"
	"github.com/subramanyam-searce/banking/typedefs"
)

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	id_str := mux.Vars(r)["id"]
	id, err := strconv.Atoi(id_str)
	if err != nil {
		fmt.Println(err)
	}
	rBody := map[string]any{}
	var updatedCustomer *typedefs.Customer

	json.NewDecoder(r.Body).Decode(&rBody)

	customer := dbconnect.GetCustomer(id)

	if customer != nil {
		for k, v := range rBody {
			updatedCustomer = dbconnect.UpdateCustomer(id, k, v)
		}
		SendResponse(w, r, map[string]*typedefs.Customer{"updated_customer": updatedCustomer})
	} else {
		w.WriteHeader(http.StatusNotFound)
		SendResponse(w, r, map[string]string{"message": "Error: User with id " + strconv.Itoa(id) + " not found"})
	}
}
