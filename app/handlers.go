package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seanyang20/banking/service"
)

// type Customer struct {
// 	Name    string `json:"full_name" xml:"name"` // for encoding and decoding structs
// 	City    string `json:"city" xml:"city"`
// 	Zipcode string `json:"zip_code" xml:"zipcode"`
// }

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

	customers, err := ch.service.GetAllCustomer()

	// Response header
	w.Header().Add("Content-Type", "application/json")

	// if r.Header.Get("Content-Type") == "application/xml" {
	// 	// xml encoding
	// 	w.Header().Add("Content-Type", "application/xml")
	// 	xml.NewEncoder(w).Encode(customers)
	// } else {
	// 	// json encoding
	// 	w.Header().Add("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(customers)
	// }

	if err != nil {
		// w.Header().Add("Content-Type", "application/json") // setting error in json format
		// w.WriteHeader(err.Code)
		// json.NewEncoder(w).Encode(err.AsMessage())
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		// w.Header().Add("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode(customer)
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		// w.Header().Add("Content-Type", "application/json") // setting error in json format
		// w.WriteHeader(err.Code)
		// json.NewEncoder(w).Encode(err.AsMessage())
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		// w.Header().Add("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode(customer)
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
