package service

import (
	"lab2/models"
	"lab2/repository"
)

type TransactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(r repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		repo: r,
	}
}

func (s *TransactionService) CreateTransaction(tx models.Transaction) (models.Transaction, error) {
	return s.repo.InsertTransaction(tx)
}

func (s *TransactionService) GetTransaction(id int) (models.Transaction, error) {
	return s.repo.GetTransactionByID(id)
}

func (s *TransactionService) CalculateTotal(transactions []models.Transaction) float64 {
	var total float64
	for _, t := range transactions {
		total += t.Amount
	}
	return total
}

func (s *TransactionService) GroupByType(transactions []models.Transaction) map[models.TransactionType][]models.Transaction {
	grouped := make(map[models.TransactionType][]models.Transaction)
	for _, t := range transactions {
		grouped[t.Type] = append(grouped[t.Type], t)
	}
	return grouped
}

func (s *TransactionService) MaxExpenseType(transactions []models.Transaction) models.TransactionType {
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
