package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {

	connString := "postgres://postgres:1234@localhost:5432/postgres"

	var err error

	DB, err = pgxpool.New(
		context.Background(),
		connString,
	)

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	err = DB.Ping(context.Background())

	if err != nil {
		log.Fatal("Database ping failed:", err)
	}

	log.Println("Database connected successfully")
}