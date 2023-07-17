package budget

import (
	// import models
	"github.com/bajeti-lab/bajeti/internal/models"
)

// a function to create a user, given, name, email and password
func CreateUser(name, email, password string) (models.User, error){
	// create a new user
	user := models.User{
		Username:   name,
		Email:    	email,
		Password: 	password,
	}

	return user, nil
}

