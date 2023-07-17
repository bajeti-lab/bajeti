package utils

import (
	"time"
	"github.com/bajeti-lab/bajeti/internal/models"
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

// implement a factory handler to create necessary budget categories
func CreateCategory(categorystruct struct {}) {
	category := models.Categories{
		CategoryId: 1,
		Title: categorystruct.Title,
		Description: categorystruct.Description,
		Amount: 0,
	}
}
