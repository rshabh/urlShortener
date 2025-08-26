package handlers

import (
	"URLSHORTENER/internal/models"
	"URLSHORTENER/internal/services"
	"URLSHORTENER/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SaveInDb(w http.ResponseWriter, r *http.Request) {
	//// for jwt

	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := utils.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}

	fmt.Println("Welcome to the the protected area of saving in db")

	var l models.Long

	lerr := json.NewDecoder(r.Body).Decode(&l)
	if lerr != nil {
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
	log.Println("the long is :=  " + l)
	http.Redirect(w, r, l, http.StatusMovedPermanently)
}
