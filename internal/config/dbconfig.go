package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB(databaseURL string) *sql.DB {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("database not working %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("database not working %v", err)
	}

	fmt.Println("Connected to database")

	return db
}
