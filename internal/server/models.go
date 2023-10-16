package server

import (
	"github.com/jinzhu/gorm"
)

// User is a representation of a user in the database
type User struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"unique;not null"`
	Email        string `gorm:"type:varchar(100);unique;not null"`
	Password     string `gorm:"not null"`
	Budgets      []Budget
	Expenses     []Expense
	Incomes      []Income
	Savings      []Saving
	Goals        []Goal
	Transactions []Transaction
}

type Account struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	Title     string
	BudjetID  int
	ExpenseID int
}

type Budget struct {
	gorm.Model
	Amount    float64
	StartDate string `gorm:"datatypes:DATE"`
	EndDate   string `gorm:"datatypes:DATE"`
	UserID    int
	AccountID int `gorm:"foreignkey:AccountID"`
}

type Expense struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	Title       string
	Amount      float64
	Desctiption string
	Date        string `gorm:"datatypes:DATE"`
	UserID      int
	AccountID   int `gorm:"foreignkey:AccountID"`
}

type Income struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	Title       string
	Amount      float64
	Desctiption string
	Date        string `gorm:"datatypes:DATE"`
	UserID      int
}

type Saving struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	Title       string
	Amount      float64
	Description string
	Date        string `gorm:"datatypes:DATE"`
	UserID      int
}

type Goal struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	Title       string
	Amount      float64
	Description string
	Date        string `gorm:"datatypes:DATE"`
	UserID      int
}

type Transaction struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	Title       string
	Amount      float64
	Description string
	Date        string `gorm:"datatypes:DATE"`
	UserID      int
}
