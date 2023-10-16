package account

import (
	"github.com/bajeti-lab/bajeti/internal/server"
	"github.com/bajeti-lab/bajeti/pkg/utils"
	"gorm.io/gorm"
)

func CreateAccount(db *gorm.DB, name string, amount float64, userID int) {
	// Create the account
	a := &server.Account{
		Title:  name,
		Amount: amount,
	}

	// Enter the account into the database
	utils.CreateOrUpdate(db, a)
}

func GetAccount(db *gorm.DB, id int) interface{} {
	// Get the account
	a := &server.Account{}
	account := utils.GetRecordByID(db, a, id)
	return account
}

func GetAllAccounts(db *gorm.DB) interface{} {
	// Get all accounts
	a := &server.Account{}
	accounts := utils.GetAllRecords(db, a)
	return accounts
}

func DeleteAccount(db *gorm.DB, id int) error {
	// Delete the account
	a := &server.Account{}
	result := db.Delete(a, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
