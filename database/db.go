package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB // Package-level variable to hold the database connection

func InitDB() (*sql.DB, error) {
	connStr := "postgres://postgres:0000@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	log.Println("connected to db")

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}
