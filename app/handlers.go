package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/seanyang20/banking/service"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"` // for encoding and decoding structs
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

// // ResponseWriter is what we are sending back to client
// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World!")
// }

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"Ashish", "New Delhi", "110075"},
	// 	{"Rob", "New Delhi", "110075"},
	// }

	customers, _ := ch.service.GetAllCustomer()

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