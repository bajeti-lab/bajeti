package budget

import (
	"github.com/bajeti-lab/bajeti/internal/models"
)

func CreateExpense( name, description string, amount int64) (models.Expense, error) {
	// Create expenses
	expense := models.Expense{
		Name: name,
		Description: description,
		Amount: amount,
	}

	return expense, nil
}


