package database

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegistration struct {
	Username         string `json:"username"`
	Password         string `json:"password"`
	RegistrationCode string `json:"registrationcode"`
}

// ADD DATABASE CONNECTION
