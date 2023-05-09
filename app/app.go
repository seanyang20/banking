package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seanyang20/banking/domain"
	"github.com/seanyang20/banking/service"
)

func Start() {
	// // define routes (contains default multiplexer)
	// http.HandleFunc("/greet", greet)
	// http.HandleFunc("/customers", getAllCustomers)

	// // starting server
	// log.Fatal(http.ListenAndServe("localhost:8000", nil))

	// // ------
	// // custom mux
	// // -------
	// mux := http.NewServeMux()

	// // define routes
	// mux.HandleFunc("/greet", greet)
	// mux.HandleFunc("/customers", getAllCustomers)

	// // starting server
	// log.Fatal(http.ListenAndServe("localhost:8000", mux))

	// // ------
	// // gorilla/ mux
	// // -------
	// router := mux.NewRouter()
	// // define routes
	// router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	// router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	// router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// // starting server
	// log.Fatal(http.ListenAndServe("localhost:8000", router))

	// ------
	// implementing hexagonal architecture
	// -------
	router := mux.NewRouter()
	//wiring (injected Database Adapter here (previously was stub))
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

// func createCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Post request received")
// }

// func getCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Fprint(w, vars["customer_id"])
// }
