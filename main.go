package main

import (
	"fmt"
)

type TransactionType struct {
	ID       int
	Name     string
	Category string
}

type Transaction struct {
	ID     int
	Amount float64
	Type   TransactionType
	Note   string
}

func GroupByType(transactions []Transaction) map[TransactionType][]Transaction {
	grouped := make(map[TransactionType][]Transaction)
	for _, t := range transactions {
		grouped[t.Type] = append(grouped[t.Type], t)
	}
	return grouped
}

func CalculateTotal(transactions []Transaction) float64 {
	var total float64
	for _, t := range transactions {
		total += t.Amount
	}
	return total
}

func MaxExpenseType(transactions []Transaction) TransactionType {
	expenseMap := make(map[TransactionType]float64)
	for _, t := range transactions {
		if t.Amount < 0 {
			expenseMap[t.Type] += -t.Amount
		}
	}

	var maxType TransactionType
	var maxAmount float64
	for typ, amount := range expenseMap {
		if amount > maxAmount {
			maxAmount = amount
			maxType = typ
		}
	}
	return maxType
}

func main() {
	salary := TransactionType{1, "Зарплата", "Доход"}
	office := TransactionType{2, "Оренда офісу", "Витрати"}
	supplies := TransactionType{3, "Канцелярія", "Витрати"}

	transactions := []Transaction{
		{1, 5000, salary, "Зарплата за березень"},
		{2, -1500, office, "Оренда офісу"},
		{3, -300, supplies, "Папір, ручки"},
		{4, -2000, office, "Оренда офісу за квітень"},
	}

	grouped := GroupByType(transactions)
	total := CalculateTotal(transactions)
	maxExpense := MaxExpenseType(transactions)

	fmt.Println("Загальний баланс:", total)
	fmt.Println("Тип з найбільшою сумою витрат:", maxExpense.Name)
	fmt.Println("Групування транзакцій:")
	for typ, txs := range grouped {
		fmt.Println("-", typ.Name, "→", len(txs), "шт.")
	}
}
