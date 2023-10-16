package savings

import (
	"github.com/bajeti-lab/bajeti/internal/server"
	"github.com/bajeti-lab/bajeti/pkg/utils"
	"gorm.io/gorm"
)

func CreateSavings(db *gorm.DB, title string, amount float64, description string, userID uint) {
	// Create the savings
	s := &server.Saving{
		Title:       title,
		Amount:      amount,
		Description: description,
		UserID:      userID,
	}

	// Enter the savings into the database
	utils.CreateOrUpdate(db, s)
}
