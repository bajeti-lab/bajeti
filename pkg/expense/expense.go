package expense

import (
	"github.com/bajeti-lab/bajeti/internal/server"
	"github.com/bajeti-lab/bajeti/pkg/utils"
	"gorm.io/gorm"
)

func CreateExpense(db *gorm.DB, title string, amount float64, description string, date string, userID uint, accountID uint) {
	// Create the expense
	e := &server.Expense{
		Title:       title,
		Amount:      amount,
		Description: description,
		Date:        date,
		UserID:      userID,
	}

	// Enter the expense into the database
	utils.CreateOrUpdate(db, e)
}

func GetExpense(db *gorm.DB, id int) interface{} {
	// Get the expense
	e := &server.Expense{}
	expense := utils.GetRecordByID(db, e, id)
	return expense
}

func GetAllExpenses(db *gorm.DB) interface{} {
	// Get all expenses
	e := &server.Expense{}
	expenses := utils.GetAllRecords(db, e)
	return expenses
}

func DeleteExpense(db *gorm.DB, id int) error {
	// Delete the expense
	e := &server.Expense{}
	result := db.Delete(e, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
