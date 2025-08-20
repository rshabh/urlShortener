package main

import (
	"URLSHORTENER/internal/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.RegisterRoutes()
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
