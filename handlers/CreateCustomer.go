package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/subramanyam-searce/banking/dbconnect"
	"github.com/subramanyam-searce/banking/typedefs"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	data := typedefs.Customer{}
	json.NewDecoder(r.Body).Decode(&data)

	id := data.Id
	foundCustomer := dbconnect.GetCustomer(id)

	if foundCustomer == nil {
		err := dbconnect.CreateCustomer(&data)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode(map[string]typedefs.Customer{"data_added": data})
		}
	} else {
		json.NewEncoder(w).Encode(map[string]string{"message": "User with id " + strconv.Itoa(id) + " already exist!"})
	}

}
