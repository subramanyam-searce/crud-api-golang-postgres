package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter()

	for _, v := range Routes {
		router.HandleFunc(v.path, v.handler).Methods(v.method)
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}
