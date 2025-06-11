package db

import (
	"database/sql"

	"github.com/gofiber/fiber/v2/log"
	_ "github.com/lib/pq" // PostgreSQL driver)
)

func NewDB() (*sql.DB, error) {
	connStr := "host=172.18.125.255 port=5432 user=postgres password=vini dbname=shortlink sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Info("Connected to the database successfully")
	return db, nil
}
