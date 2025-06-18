package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
	databaseUrl := os.Getenv("DATABASE_URL")

	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Fatalf("Unable to parse DATABASE_URL: %v\n", err)
	}

	// ⚙️ Disable prepared statement caching
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	log.Println("✅ Database connected successfully")
}

func Close() {
	DB.Close()
}
