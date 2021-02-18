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

	addTransactionsAPIs(r)

	fmt.Println("Starting External HTTP Server at 8070")
	log.Fatal(http.ListenAndServe(":8070", r))

}

func addTransactionsAPIs(r *mux.Router) {
	r.HandleFunc("/transactionservice/transaction/", api.UpdateTransactions).
		Methods("PUT")

	r.HandleFunc("/transactionservice/transaction/", api.GetTransactions).
		Methods("GET")
	
	r.HandleFunc("/transactionservice/types/", api.GetTransactionsFromTypes).
		Methods("GET")
	
	r.HandleFunc("/transactionservice/sum/", api.GetTransactionsSum).
		Methods("GET")
}
