package main

import (
	"GymTracker/configfile"
	"GymTracker/logfile"
	"GymTracker/webserver"
	"fmt"
	"log"
	"os"
)

func init() {

	// Set the desired configuration settings here
	pathLogfile := "./logfile.log"
	pathConfigfile := "./config.json"

	fmt.Println("Initializing application...")

	// Initialize logfile
	_, err := logfile.InitLogfile(pathLogfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Initialize configfile
	err = configfile.InitConfigfile(pathConfigfile)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func main() {

	fmt.Println("Application has initialized succesfully.")
	fmt.Println("Application is now starting.")

	portWebserver := configfile.Configuration.Webserver.Port

	// Start the webserver
	webserver.StartWebserver(portWebserver)
}
