package store

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {

	ctx := context.Background()
	connStr := "postgres://postgres:27042002@localhost:5432/practisedb"

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("Could not ping DB: %v\n", err)
	}

	fmt.Println("✅ Connected to PostgreSQL")
	DB = pool

}
