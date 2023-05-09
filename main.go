package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"` // for encoding and decoding structs
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

func main() {

	// define routes (contains default multiplexer)
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

// ResponseWriter is what we are sending back to client
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Ashish", "New Delhi", "110075"},
		{"Rob", "New Delhi", "110075"},
	}

	// Response header
	w.Header().Add("Content-Type", "application/json")

	if r.Header.Get("Content-Type") == "application/xml" {
		// xml encoding
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		// json encoding
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
