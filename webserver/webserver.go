package webserver

import (
	"GymTracker/database"
	"encoding/json"
	"net/http"
)

// Define constant values
const (
	invalidRequestBody      = "Invalid request body"
	internalError           = "Internal server error"
	unauthorized            = "Unauthorized"
	usernameExists          = "Username already exists"
	invalidCredentials      = "Invalid credentials"
	methodNotAllowed        = "Method not allowed"
	registrationCodeInvalid = "Registration code is not valid"
)

// MODIFY API KEY FUNCTIONALITY

func StartWebserver(port int) {

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registrationHandler)
}

// Valid API keys
var validAPIKeys = map[string]bool{
	"test": true,
}

// Function to check if the given API key is valid
func isAPIKeyValid(apiKey string) bool {

	return validAPIKeys[apiKey]
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	// Check if the used method is post
	if r.Method != http.MethodPost {
		http.Error(w, methodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	// Check if the given API key is valid
	apiKey := r.Header.Get("Authorization")
	if !isAPIKeyValid(apiKey) {
		http.Error(w, unauthorized, http.StatusUnauthorized)
		return
	}

	// Insert post data into UserLogin struct
	var user database.UserLogin
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, invalidRequestBody, http.StatusBadRequest)
		return
	}

	// ADD CREDENTIAL CHECK IN DATABASE
	// IF CORRECT REDIRECT TO DASHBOARD
}

func registrationHandler(w http.ResponseWriter, r *http.Request) {

	// Check if the used method is post
	if r.Method != http.MethodPost {
		http.Error(w, methodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	// Check if the given API key is valid
	apiKey := r.Header.Get("Authorization")
	if !isAPIKeyValid(apiKey) {
		http.Error(w, unauthorized, http.StatusUnauthorized)
		return
	}

	// Insert post data into UserRegister struct
	var user database.UserRegistration
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, invalidRequestBody, http.StatusBadRequest)
		return
	}

	// HANDLE DATABASE STUFF
	// REDIRECT TO LOGIN PAGE IF SUCCESFUL
}
