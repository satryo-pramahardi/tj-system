package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init(dsn string) {
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to open DB: %v", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatalf("❌ Failed to ping DB: %v", err)
	}

	log.Println("✅ Database connected")
}
