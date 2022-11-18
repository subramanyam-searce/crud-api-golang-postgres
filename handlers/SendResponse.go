package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/subramanyam-searce/banking/typedefs"
)

func SendResponse(w http.ResponseWriter, r *http.Request, v any) {
	w.Header().Add("Content-Type", "application/json")
	var nilCustomer *typedefs.Customer = nil

	if v != nilCustomer {
		err := json.NewEncoder(w).Encode(v)
		if err != nil {
			fmt.Println("jsonEncodeError:", err)
		}
	} else {
		fmt.Println("404")
		w.WriteHeader(http.StatusNotFound)
	}
}
