package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler handles the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Bajeti")
}

// BudgetsHandler handles the budgets page
func BudgetsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Budgets")
}

// TransactionsHandler handles the transactions page
func TransactionsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Transactions")
}
