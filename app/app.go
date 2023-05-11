package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/seanyang20/banking/domain"
	"github.com/seanyang20/banking/logger"
	"github.com/seanyang20/banking/service"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}

func Start() {

	sanityCheck()

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
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	// define routes
	// router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	// router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	// router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.
		HandleFunc("/customers", ch.getAllCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).
		Methods(http.MethodGet).
		Name("GetCustomer")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).
		Methods(http.MethodPost).
		Name("NewAccount")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).
		Methods(http.MethodPost).
		Name("NewTransaction")

	am := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())

	// starting server
	// log.Fatal(http.ListenAndServe("localhost:8000", router))
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logger.Info(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

// func createCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Post request received")
// }

// func getCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Fprint(w, vars["customer_id"])
// }

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
