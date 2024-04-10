package logfile

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Create/open the logfile and set as default output for logs
func InitLogfile(pathLogfile string) (*os.File, error) {

	// Get the directory of the given logfile path
	dirLogfile := filepath.Dir(pathLogfile)

	// Check if the directory exists
	if !dirExists(dirLogfile) {
		return nil, fmt.Errorf("error initializing logfile: directory %s does not exist", dirLogfile)
	}

	// Create a logfile, if it does not exist in the directory
	logfile, err := os.OpenFile(pathLogfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("error creating/opening logfile: %s", err)
	}

	// Define the created logfile as the standard output for the logs
	log.SetOutput(logfile)

	log.Println("success: logfile has been initialized succesfully.")

	return logfile, nil
}

// Check if a directory exists
func dirExists(dir string) bool {

	_, err := os.Stat(dir)

	return !os.IsNotExist(err)
}
