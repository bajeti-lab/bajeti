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

/*
Account is a representation of an account in the database.
Accounts belong to a Budget and an Expense.
Multiple accounts can belong to a single Budget or Expense.
*/
type Account struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Title     string
	Amount    float64
	BudgetID  uint
	ExpenseID uint
}

type Budget struct {
	gorm.Model
	Amount    float64
	StartDate string `gorm:"datatypes:DATE"`
	EndDate   string `gorm:"datatypes:DATE"`
	UserID    uint
	Accounts  []*Account
}

type Expense struct {
	gorm.Model
	ID          uint `gorm:"primary_key"`
	Title       string
	Amount      float64
	Description string
	Date        string `gorm:"datatypes:DATE"`
	UserID      uint
	Accounts    []*Account
}

type Income struct {
	gorm.Model
	ID          uint `gorm:"primary_key"`
	Title       string
	Amount      float64
	Desctiption string
	Date        string `gorm:"datatypes:DATE"`
	UserID      uint
}

type Saving struct {
	gorm.Model
	ID          uint `gorm:"primary_key"`
	Title       string
	Amount      float64
	Description string
	Date        string `gorm:"datatypes:DATE"`
	UserID      uint
}

type Goal struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	Title       string
	Amount      float64
	Description string
	Date        string `gorm:"datatypes:DATE"`
	UserID      uint
}

type Transaction struct {
	gorm.Model
	ID          uint `gorm:"primary_key"`
	Title       string
	Amount      float64
	Description string
	Date        string `gorm:"datatypes:DATE"`
	UserID      uint
}
