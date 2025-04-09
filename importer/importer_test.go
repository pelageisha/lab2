package importer

import (
	"testing"

	"lab2/db"
	"lab2/repository"
)

func BenchmarkImportCSV(b *testing.B) {
	db.InitDB()
	repo := repository.NewTransactionRepository(db.DB)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ImportCSV("../import.csv", *repo)
	}
}
