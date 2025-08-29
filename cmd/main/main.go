package main

import (
	"URLSHORTENER/internal/routes"
	"URLSHORTENER/internal/store"
	"log"
	"net/http"
)

func main() {

	store.InitDB()
	r := routes.RegisterRoutes()
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
