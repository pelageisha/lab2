package models

type Transaction struct {
	ID     int    `db:"id"`
	Amount int    `db:"amount"`
	Note   string `db:"note"`
	Type   Type   `db:"type"`
}

type Type struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}
