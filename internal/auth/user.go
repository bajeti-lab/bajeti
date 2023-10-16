package auth

import (
	"github.com/bajeti-lab/bajeti/internal/server"
	"github.com/bajeti-lab/bajeti/pkg/utils"
	"gorm.io/gorm"
)

// CreateUser creates a new user
func CreateUser(db *gorm.DB, username string, password string, email string) {
	// Create the user
	u := &server.User{
		Name:     username,
		Password: password,
		Email:    email,
	}

	// Enter the user into the database
	utils.CreateOrUpdate(db, u)
}

// GetUser gets a user by ID
func GetUser(db *gorm.DB, id int) interface{} {
	// Get the user
	u := &server.User{}
	user := utils.GetRecordByID(db, u, id)
	return user
}

// GetAllUsers gets all users
func GetAllUsers(db *gorm.DB) interface{} {
	// Get all users
	u := &server.User{}
	users := utils.GetAllRecords(db, u)
	return users
}

// DeleteUser deletes a user by ID
func DeleteUser(db *gorm.DB, id int) error {
	// Delete the user
	u := &server.User{}
	result := db.Delete(u, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
