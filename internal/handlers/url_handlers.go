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

	services.GetShortAndInsert(r.Context(), l)
	log.Println("The long url is saved in the db")
	json.NewEncoder(w).Encode(services.GetUrl(r.Context(), l))

}

func Redirect(w http.ResponseWriter, r *http.Request) {
	log.Println("redirect funcion called")
	s := chi.URLParam(r, "s")
	l := services.GetLong(r.Context(), s)
	log.Println("the short is" + s)
	log.Println("the long is " + l)
	http.Redirect(w, r, l, http.StatusMovedPermanently)
}
