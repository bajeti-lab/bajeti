package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDatabase() (*gorm.DB, error) {
	// load the environment variable
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create and open the database
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_TIMEZONE"),
	)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		CreateDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", os.Getenv("DB_NAME"))
		DB.Exec(CreateDatabaseCommand)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Error opening database")
		}
	}

	return DB, err
}

// create a record in the database if there is none or update it if it exists
func CreateOrUpdate(db *gorm.DB, record interface{}) error {
	result := db.Save(record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// get a record from the database by its id
func GetRecordByID(db *gorm.DB, record interface{}, id int) interface{} {
	db.First(&record, id)
	return record
}

// get all records from the database
func GetAllRecords(db *gorm.DB, record interface{}) interface{} {
	db.Find(&record)
	return record
}

// delete a record from the database by its id
func DeleteRecordByID(db *gorm.DB, record interface{}, id int) error {
	db.Delete(&record, id)
	return nil
}
