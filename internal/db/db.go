package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/log"
	_ "github.com/lib/pq" // PostgreSQL driver)
)

func NewDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Buat connection string
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	// connStr := `host= port=5432 user=postgres password=vini dbname=shortlink sslmode=disable`
	// connStr := "host=172.18.125.255 port=5432 user=postgres password=vini dbname=shortlink sslmode=disable"
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
