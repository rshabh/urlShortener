package handlers

import (
	"URLSHORTENER/internal/models"
	"URLSHORTENER/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

	cookie := http.Cookie{
		Name:     "jwt_token",
		Value:    ts,
		Expires:  time.Now().Add(time.Hour * 24), // Match JWT expiry
		HttpOnly: true,                           // Important for security
		Secure:   true,                           // Use `Secure: true` in production with HTTPS
		Path:     "/",                            // Make cookie available across the entire site
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, cookie.Value)

}
