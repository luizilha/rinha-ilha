package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func PostgresConnection() *sql.DB {
	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", 
    os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
