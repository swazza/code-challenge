package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"volteo-challenge/src/resources"
)

func getBuilding(res http.ResponseWriter, req *http.Request) {
	building := resources.Building{Id:"1"}
	tmp := building.Get()
	buildingJson, _ := json.Marshal(tmp)
	fmt.Fprintf(res, string(buildingJson))
}

func createBuilding(res http.ResponseWriter, req *http.Request) {}

func BuildingSubrouter(router *mux.Router) {
	subrouter := router.PathPrefix("/buildings").Subrouter()
	subrouter.HandleFunc("/", getBuilding).Methods("GET")
	subrouter.HandleFunc("/", createBuilding).Methods("POST")
}