package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
	dsn := "user=postgres password=polina2003 dbname=labdb sslmode=disable"

	var err error
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln("Помилка підключення до бази даних:", err)
	}

	log.Println("Успішне підключення до бази даних")
}
