package main

import (
	"testing"

	"lab2/models"

	"github.com/stretchr/testify/require"
)

func TestCalculateTotal(t *testing.T) {
	tests := []struct {
		name        string
		input       []models.Transaction
		expectedSum float64
	}{
		{
			name: "Only income",
			input: []models.Transaction{
				{Amount: 1000},
				{Amount: 2000},
			},
			expectedSum: 3000,
		},
		{
			name: "Income and expense",
			input: []models.Transaction{
				{Amount: 5000},
				{Amount: -2000},
				{Amount: -1000},
			},
			expectedSum: 2000,
		},
		{
			name:        "Empty list",
			input:       []models.Transaction{},
			expectedSum: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := CalculateTotal(tc.input)
			require.Equal(t, tc.expectedSum, result)
		})
	}
}

func TestMaxExpenseType(t *testing.T) {
	office := models.TransactionType{ID: 1, Name: "Оренда"}
	salary := models.TransactionType{ID: 2, Name: "Зарплата"}
	supplies := models.TransactionType{ID: 3, Name: "Канцелярія"}

	tests := []struct {
		name        string
		input       []models.Transaction
		expectedMax models.TransactionType
	}{
		{
			name: "Оренда найбільша",
			input: []models.Transaction{
				{Amount: -1000, Type: supplies},
				{Amount: -3000, Type: office},
				{Amount: -1500, Type: salary},
			},
			expectedMax: office,
		},
		{
			name: "Канцелярія найбільша",
			input: []models.Transaction{
				{Amount: -500, Type: office},
				{Amount: -5000, Type: supplies},
			},
			expectedMax: supplies,
		},
		{
			name:        "Немає витрат",
			input:       []models.Transaction{{Amount: 1000}, {Amount: 2000}},
			expectedMax: models.TransactionType{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := MaxExpenseType(tc.input)
			require.Equal(t, tc.expectedMax, result)
		})
	}
}
