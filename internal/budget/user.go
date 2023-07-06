package budget

// use the user model to store the user's budget data
import (
	"github.com/bajeti-lab/bajeti/internal/models"
)

// UserBudget is a struct that holds the user's budget data
type UserBudget struct {
	// the user's budget data
	Budget models.Budget
	// the user's income data
	Income []models.Income
	// the user's expense data
	Expenses []models.Expense
	// the user's expense categories data
	ExpenseCategories []models.ExpenseCategories
}

