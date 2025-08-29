package main

import (
	"URLSHORTENER/internal/routes"
	"URLSHORTENER/internal/store"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	store.InitDB()
	r := routes.RegisterRoutes()
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
