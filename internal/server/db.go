package server

import (
	"github.com/bajeti-lab/bajeti/pkg/utils"
)

func OpenDatabase() {
	// Create the database
	db, err := utils.CreateDatabase()

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Budget{})
	db.AutoMigrate(&Expense{})
	db.AutoMigrate(&Income{})
	db.AutoMigrate(&Saving{})
	db.AutoMigrate(&Goal{})
	db.AutoMigrate(&Transaction{})

}
