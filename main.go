package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/onimarufl/mux/service/testservice"
)

func main() {

	service := testservice.NewHandler(
		testservice.NewService(),
	)

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/allgroceries", service.AllGroceries)

	log.Fatal(http.ListenAndServe(":8080", r))
}
