package domain

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/seanyang20/banking/errs"
	"github.com/seanyang20/banking/logger"
)

type CustomerRepositoryDb struct {
	// client *sql.DB
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	// var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		// rows, err = d.client.Query(findAllSql)
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		// rows, err = d.client.Query(findAllSql, status)
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// rows, err := d.client.Query(findAllSql)
	if err != nil {
		// log.Println("Error while querying customers table " + err.Error())
		logger.Error("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// sqlx helps us remove the code commented out below
	// err = sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	// log.Println("Error while scanning customers " + err.Error())
	// 	logger.Error("Error while scanning customers " + err.Error())
	// 	return nil, errs.NewUnexpectedError("Unexpected database error")
	// }

	// // looping over all the rows in the query result
	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	// 	if err != nil {
	// 		// log.Println("Error while scanning customers " + err.Error())
	// 		logger.Error("Error while scanning customers " + err.Error())
	// 		return nil, errs.NewUnexpectedError("Unexpected database error")
	// 	}
	// 	customers = append(customers, c)
	// }
	return customers, nil

}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	// before sqlx
	// row := d.client.QueryRow(customerSql, id)
	// var c Customer
	// err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)

	// after sqlx
	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		// log.Println("Error while scanning customer " + err.Error())
		// return nil, err
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found") // when is db is up (docker)
		} else {
			// log.Println("Error while scanning customer " + err.Error())
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error") // when db is down (docker)
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// client, err := sql.Open("mysql", "root:Moopie1497$!@tcp(localhost:3306)/banking")
	// client, err := sqlx.Open("mysql", "root:Moopie1497$!@tcp(localhost:3306)/banking")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
