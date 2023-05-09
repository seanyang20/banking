package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func main() {

	// define routes (contains default multiplexer)
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Ashish", "New Delhi", "110075"},
		{"Rob", "New Delhi", "110075"},
	}

	w.Header().Add("Content-Type", "application/json")

	// encodes all of our customers in JSON format
	json.NewEncoder(w).Encode(customers)
}
