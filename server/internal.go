package server

import (
	"fmt"
	"log"
	"net/http"
	"vartul14/locus/api"

	"github.com/gorilla/mux"
)

func Start() {
	go startExternalHTTPServer()
}

func startExternalHTTPServer() {
	r := mux.NewRouter()

	addDirectionsAPIs(r)

	fmt.Println("Starting External HTTP Server at 8070")
	log.Fatal(http.ListenAndServe(":8070", r))

}

func addDirectionsAPIs(r *mux.Router) {
	r.HandleFunc("/directions", api.GetDirections).
		Methods("GET")
}
