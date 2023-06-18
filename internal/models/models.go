package models

// Budget represents a budget in the application
type Budget struct {
	ID          int64
	Name        string
	Description string
	Amount      float64
}

// Transaction represents a transaction in the application
type Transaction struct {
	ID          int64
	Name        string
	Description string
	Amount      float64
	BudgetID    int64
}

// User represents a user in the application
type User struct {
	ID       int64
	Username string
	Password string
}
