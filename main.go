package main

import (
	"GymTracker/logfile"
	"fmt"
)

func init() {

	// Set the desired configuration settings here
	pathLogfile := "./logfile.log"

	fmt.Println("Initializing application...")

	// Initialize logfile
	_, err := logfile.InitLogfile(pathLogfile)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

}
