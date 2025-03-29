package models

type TransactionType struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Category string `json:"category" db:"category"`
}

type Transaction struct {
	ID     int             `json:"id" db:"id"`
	Amount float64         `json:"amount" db:"amount"`
	Type   TransactionType `json:"type"`
	Note   string          `json:"note"`
}
