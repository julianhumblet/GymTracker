package configfile

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var Configuration configuration

type configuration struct {
	Server struct {
		ServerPort int `json:"serverPort"`
	} `json:"server"`
	Database struct {
		Host         string `json:"host"`
		Port         int    `json:"port"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		DatabaseName string `json:"databaseName"`
	} `json:"database"`
}

func InitConfigfile(pathConfigfile string) error {

	if !fileExists(pathConfigfile) {

		configfile, err := createConfigfile(pathConfigfile)
		if err != nil {
			return err
		}

		// Turn the struct into JSON format.
		encodedJsonStruct, err := json.MarshalIndent(Configuration, "", "\t")
		if err != nil {
			return err
		}

		// Write the JSON data to the created configfile.
		_, err = configfile.Write(encodedJsonStruct)
		if err != nil {
			return err
		}

		log.Println("Configure the settings in the configfile.")
		os.Exit(1)
	}

	// Open the configfile.
	configfile, err := openFile(pathConfigfile)
	if err != nil {
		return err
	}
	defer configfile.Close()

	// Define the configfile and assign to global variable.
	err = json.NewDecoder(configfile).Decode(&Configuration)
	if err != nil {
		return err
	}

	log.Println("success: configfile has been created/initialized succesfully.")

	return nil
}

func fileExists(path string) bool {

	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}

	log.Printf("error checking configfile existence: %s", err)
	os.Exit(1)

	return false
}

func createConfigfile(filePath string) (*os.File, error) {

	configfile, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("error: creating configfile: %s", err)
	}

	return configfile, nil
}

func openFile(filePath string) (*os.File, error) {

	openedFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error: opening configfile: %s", err)
	}

	return openedFile, nil
}
