package budget

import (
	"github.com/bajeti-lab/bajeti/internal/server"
	"github.com/bajeti-lab/bajeti/pkg/utils"
	"gorm.io/gorm"
)

func CreateBudget(db *gorm.DB, amount float64, startDate string, endDate string, userID uint) {
	// Create the budget
	b := &server.Budget{
		Amount:    amount,
		StartDate: startDate,
		EndDate:   endDate,
		UserID:    userID,
	}

	// Enter the budget into the database
	utils.CreateOrUpdate(db, b)
}

func GetBudget(db *gorm.DB, id int) interface{} {
	// Get the budget
	b := &server.Budget{}
	budget := utils.GetRecordByID(db, b, id)
	return budget
}

func GetAllBudgets(db *gorm.DB) interface{} {
	// Get all budgets
	b := &server.Budget{}
	budgets := utils.GetAllRecords(db, b)
	return budgets
}

func DeleteBudget(db *gorm.DB, id int) error {
	// Delete the budget
	b := &server.Budget{}
	result := db.Delete(b, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
