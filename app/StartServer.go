package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/").Subrouter()

	for _, v := range Routes {
		subRouter.HandleFunc(v.path, v.handler).Methods(v.method)
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}
