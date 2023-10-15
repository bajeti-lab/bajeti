package server

import (
	"github.com/jinzhu/gorm"
)

// User is a representation of a user in the database
type User struct {
	gorm.Model
	ID           int    `gorm:"primary_key"`
	name         string `gorm:"unique_index;not null"`
	email        string `gorm:"type:varchar(100);unique_index;not null"`
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
	ID        int `gorm:"primary_key"`
	Title     string
	BudjetID  int
	ExpenseID int
}

type Budget struct {
	gorm.Model
	Amount    float64
	startDate string `gorm:"datatypes:DATE"`
	endDate   string `gorm:"datatypes:DATE"`
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
	ID          int `gorm:"primary_key"`
	Title       string
	Amount      float64
	Description string
	Date        string `gorm:"datatypes:DATE"`
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
