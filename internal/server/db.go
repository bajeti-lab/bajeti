package server

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDatabase() {
	// Load environment variables
	// dialect := os.Getenv("DB_DIALECT")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	// Connect to database
	dsn := "host=" + host + " port=" + port + " user=" + user + " dbname=" + dbname + " password=" + password + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
