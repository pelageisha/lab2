package main

import (
	"log"
	"net/http"

	"lab2/db"
	"lab2/handler"
	"lab2/models"
	"lab2/repository"
	"lab2/service"

	"github.com/gorilla/mux"
)

func GroupByType(transactions []models.Transaction) map[models.TransactionType][]models.Transaction {
	grouped := make(map[models.TransactionType][]models.Transaction)
	for _, t := range transactions {
		grouped[t.Type] = append(grouped[t.Type], t)
	}
	return grouped
}

func CalculateTotal(transactions []models.Transaction) float64 {
	var total float64
	for _, t := range transactions {
		total += t.Amount
	}
	return total
}

func MaxExpenseType(transactions []models.Transaction) models.TransactionType {
	expenseMap := make(map[models.TransactionType]float64)
	for _, t := range transactions {
		if t.Amount < 0 {
			expenseMap[t.Type] += -t.Amount
		}
	}

	var maxType models.TransactionType
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
	db.InitDB()

	repo := repository.NewTransactionRepository(db.DB)
	srv := service.NewTransactionService(repo)
	h := handler.NewTransactionHandler(srv)

	r := mux.NewRouter()
	r.HandleFunc("/transactions", h.CreateTransaction).Methods("POST")
	r.HandleFunc("/transactions/{id}", h.GetTransaction).Methods("GET")

	log.Println("Сервер запущенний на порту :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
