package handlers

import (
	"URLSHORTENER/internal/models"
	"URLSHORTENER/internal/services"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SaveInMap(w http.ResponseWriter, r *http.Request) {
	var l models.Long

	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	services.GetShort(l)
	log.Println("The long url is saved in the map")

	m := services.GetMap()
	s := "http://localhost:8080/redirect/" + m[l.Long]
	json.NewEncoder(w).Encode(s)

}

func Redirect(w http.ResponseWriter, r *http.Request) {
	s := chi.URLParam(r, "s")
	l := services.GetLongFromShort(s)
	log.Println("the short is" + s)
	log.Println("the long is " + l)
	http.Redirect(w, r, l, http.StatusMovedPermanently)
}

func GetMap(w http.ResponseWriter, r *http.Request) {
	m := services.GetMap()
	json.NewEncoder(w).Encode(m)
}
