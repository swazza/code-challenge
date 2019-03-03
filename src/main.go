package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"volteo-challenge/src/routes"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello, World!")
	}).Methods("GET")

	srv := &http.Server {
		Handler: r,
		Addr: "0.0.0.0:" + os.Getenv("PORT"),
	}

	routes.BuildingSubrouter(r)

	log.Fatal(srv.ListenAndServe())
}