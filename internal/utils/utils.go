package utils

import (
	"time"
	"fmt"
)

// GetCurrentTime returns the current time as a formatted string
func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// CalculateTotal calculates the total amount from a list of values
func CalculateTotal(values []float64) float64 {
	total := 0.0
	for _, value := range values {
		total += value
	}
	return total
}
// Capture expenses details
// The items should be stored as a list 
// Iterate list to calculate the total 
func getExpenditures() {
	expenses := make(map[string]float64)
	total := 0.0

// Collect expense details
for {
	var expenseName string
	var expenseAmount float64

	fmt.Print("Enter expense name (or 'q' to quit):")
	fmt.Scanln(&expenseName)

	// Break the loop if the user enters 'q'
	if expenseName == "q" {
		break
	}

	fmt.Print("Enter expense amount: ")
	fmt.Scanln(&expenseAmount)

	expenses[expenseName] = expenseAmount
	total += expenseAmount
}

fmt.Println("\nMonthly Expenditure:")
fmt.Println("====================")

// Print each expense
for name, amount := range expenses {
	fmt.Printf("%s: $%.2f\n", name, amount)
	fmt.Printf("Total: $%.2f\n", total)
}
}
