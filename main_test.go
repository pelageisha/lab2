package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateTotal(t *testing.T) {
	tests := []struct {
		name        string
		input       []Transaction
		expectedSum float64
	}{
		{
			name: "Only income",
			input: []Transaction{
				{Amount: 1000},
				{Amount: 2000},
			},
			expectedSum: 3000,
		},
		{
			name: "Income and expense",
			input: []Transaction{
				{Amount: 5000},
				{Amount: -2000},
				{Amount: -1000},
			},
			expectedSum: 2000,
		},
		{
			name:        "Empty list",
			input:       []Transaction{},
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
	office := TransactionType{ID: 1, Name: "Оренда"}
	salary := TransactionType{ID: 2, Name: "Зарплата"}
	supplies := TransactionType{ID: 3, Name: "Канцелярія"}

	tests := []struct {
		name        string
		input       []Transaction
		expectedMax TransactionType
	}{
		{
			name: "Оренда найбільша",
			input: []Transaction{
				{Amount: -1000, Type: supplies},
				{Amount: -3000, Type: office},
				{Amount: -1500, Type: salary},
			},
			expectedMax: office,
		},
		{
			name: "Канцелярія найбільша",
			input: []Transaction{
				{Amount: -500, Type: office},
				{Amount: -5000, Type: supplies},
			},
			expectedMax: supplies,
		},
		{
			name:        "Немає витрат",
			input:       []Transaction{{Amount: 1000}, {Amount: 2000}},
			expectedMax: TransactionType{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := MaxExpenseType(tc.input)
			require.Equal(t, tc.expectedMax, result)
		})
	}
}
