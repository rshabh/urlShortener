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

	ctx := context.Background()
	var connStr any = viper.Get("db_connectionStr")
	fmt.Println(connStr)

	pool, err := pgxpool.New(ctx, connStr.(string))
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
