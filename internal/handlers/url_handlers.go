package handlers

import (
	"URLSHORTENER/internal/models"
	"URLSHORTENER/internal/services"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SaveInDb(w http.ResponseWriter, r *http.Request) {
	var l models.Long

	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	services.GetShortAndInsert(l)
	log.Println("The long url is saved in the map")

	m, err := services.FindShortFromLong(l.Long)
	if err != nil {
		log.Println("some error occured in saveInDB function")
	}
	s := "http://localhost:8080/" + m
	json.NewEncoder(w).Encode(s)

}

func Redirect(w http.ResponseWriter, r *http.Request) {
	s := chi.URLParam(r, "s")
	l, err := services.FindLongFromShort(s)
	if err != nil {
		log.Println("error occured in redirect function")
	}
	log.Println("the short is" + s)
	log.Println("the long is " + l)
	http.Redirect(w, r, l, http.StatusMovedPermanently)
}
