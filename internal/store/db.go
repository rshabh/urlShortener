package store

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

var DB *pgxpool.Pool

func InitDB() {

	connStr := viper.GetString("cn")
	log.Println(connStr)

	ctx := context.Background()
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
