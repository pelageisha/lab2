package main

import (
	"lab2/db"
	"lab2/importer"
	"lab2/repository"
)

func main() {
	db.InitDB()
	repo := repository.NewTransactionRepository(db.DB)

	importer.ImportCSV("import.csv", repo)
}
