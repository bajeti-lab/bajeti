package models

import "time"

type Budget struct {
	Period time.Time
}

type User struct {
	UserId 		uint
	Username 	string
	Email 		string
	Password 	string
}

type ExpenseCategories struct {
	CategoryId 	uint
	Title       string
	Description string
}

type Expense struct{
	ExpenseId	uint
	Name		string
	Description	string
	CategoryId	int64
}

type Income struct{
	IncomeId	uint
	Description string
	Amount		float64
}