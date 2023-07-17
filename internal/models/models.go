package models

import "time"

type Budget struct {
	Period time.Time
}

type User struct {
	Username 	string
	Email 		string
	Password 	string
}

type Categories struct {
	CategoryId 	uint
	Title       string
	Description string
	Amount	  int64
}

type Expense struct{
	ExpenseId	uint
	Name		string
	Description	string
	Amount		int64
	CategoryId	int64
}

type Income struct{
	IncomeId	uint
	Description string
	Amount		float64
	CategoryId	uint
}
