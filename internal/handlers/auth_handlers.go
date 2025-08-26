package handlers

import (
	"URLSHORTENER/internal/models"
	"URLSHORTENER/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	services.Register(r.Context(), u)
	json.NewEncoder(w).Encode("user saved")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var ul models.UserLogin
	err := json.NewDecoder(r.Body).Decode(&ul)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	ts := services.Login(r.Context(), ul)
	if ts == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, ts)

}
