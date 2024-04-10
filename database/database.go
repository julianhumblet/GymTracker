package database

import (
	"GymTracker/configfile"
	"database/sql"
	"fmt"
	"strconv"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegistration struct {
	Username         string `json:"username"`
	Password         string `json:"password"`
	RegistrationCode string `json:"registrationcode"`
}

// MODIFY TO DATABASE PROVIDER
func OpenDBCon() (*sql.DB, error) {

	dbCreds := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		configfile.Configuration.Database.Username, configfile.Configuration.Database.Password,
		configfile.Configuration.Database.Host, strconv.Itoa(configfile.Configuration.Database.Port),
		configfile.Configuration.Database.DatabaseName)

	dbCon, err := sql.Open("mysql", dbCreds)
	if err != nil {
		return nil, fmt.Errorf("error: opening database connection has failed: %s", err)
	}

	err = dbCon.Ping()
	if err != nil {
		return nil, fmt.Errorf("error: database connection test (ping) has failed: %s", err)
	}

	return dbCon, nil
}
