package utils

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable TimeZone=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_TIMEZONE"))
	DB, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	CreateDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", os.Getenv("DB_NAME"))
	DB.Exec(CreateDatabaseCommand)

	return DB
}

func CreateOrUpdate(db *gorm.DB, record interface{}) {
	db.Save(record)
}
