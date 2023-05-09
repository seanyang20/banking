package app

import (
	"log"
	"net/http"
)

func Start() {
	// // define routes (contains default multiplexer)
	// http.HandleFunc("/greet", greet)
	// http.HandleFunc("/customers", getAllCustomers)

	// // starting server
	// log.Fatal(http.ListenAndServe("localhost:8000", nil))

	// custom mux
	mux := http.NewServeMux()

	// define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))

}
