package main

import (
	"github.com/seanyang20/banking/app"
	"github.com/seanyang20/banking/logger"
)

func main() {

	//log.Println("Starting our application...")
	logger.Info("Starting our application...")
	app.Start()

}
