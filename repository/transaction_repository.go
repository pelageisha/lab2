package repository

import (
	"lab2/models"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Save(transaction models.Transaction) error {
	_, err := r.db.NamedExec(`
		INSERT INTO transactions (amount, note, type)
		VALUES (:amount, :note, :type)
	`, transaction)
	return err
}
