package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() (*sql.DB, error) {

	url := os.Getenv("S1_DATABASE_URL")

	if url == "" {
		log.Fatalln("missing environment variable S1_DATABASE_URL")
	}

	db, error := sql.Open("pgx", url)

	if error != nil {
		return nil, error
	}

	error = db.Ping()

	if error != nil {
		return nil, error
	}

	return db, nil
}
