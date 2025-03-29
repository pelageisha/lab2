package repository

import (
	"lab2/models"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	InsertTransaction(tx models.Transaction) (models.Transaction, error)
	GetTransactionByID(id int) (models.Transaction, error)
}

type PostgresTransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &PostgresTransactionRepository{db: db}
}

func (r *PostgresTransactionRepository) InsertTransaction(tx models.Transaction) (models.Transaction, error) {
	var typeID int
	err := r.db.Get(&typeID, `
        SELECT id FROM transaction_types 
        WHERE name = $1 AND category = $2
    `, tx.Type.Name, tx.Type.Category)

	if err != nil {
		err = r.db.QueryRow(`
            INSERT INTO transaction_types (name, category)
            VALUES ($1, $2) RETURNING id
        `, tx.Type.Name, tx.Type.Category).Scan(&typeID)

		if err != nil {
			return tx, err
		}
	}

	err = r.db.QueryRow(`
        INSERT INTO transactions (amount, type_id, note)
        VALUES ($1, $2, $3) RETURNING id
    `, tx.Amount, typeID, tx.Note).Scan(&tx.ID)

	if err != nil {
		return tx, err
	}

	tx.Type.ID = typeID
	return tx, nil
}

func (r *PostgresTransactionRepository) GetTransactionByID(id int) (models.Transaction, error) {
	var tx models.Transaction
	query := `
        SELECT t.id, t.amount, t.note,
               tt.id as "type.id", tt.name as "type.name", tt.category as "type.category"
        FROM transactions t
        JOIN transaction_types tt ON t.type_id = tt.id
        WHERE t.id = $1
    `
	err := r.db.Get(&tx, query, id)
	return tx, err
}
