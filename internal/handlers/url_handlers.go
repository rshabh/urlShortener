package handlers

import (
	"URLSHORTENER/internal/models"
	"URLSHORTENER/internal/services"
	"URLSHORTENER/internal/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SaveInDb(w http.ResponseWriter, r *http.Request) {
	//// for jwt

	w.Header().Set("Content-Type", "application/json")
	// tokenString := r.Header.Get("Authorization")
	// if tokenString == "" {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	fmt.Fprint(w, "Missing authorization header")
	// 	return
	// }
	// tokenString = tokenString[len("Bearer "):]

	cookie, err := r.Cookie("jwt_token")

	if err != nil {
		fmt.Fprint(w, "No cookies found")
		return
	}

	fmt.Println("cookie found")

	tokenString := cookie.Value

	claims, lerr := utils.VerifyToken(tokenString)

	if lerr != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}

	u := claims.Uid

	fmt.Println(u)

	fmt.Println("token verified from cookie")

	fmt.Println("Welcome to the the protected area of saving in db")

	var l models.Long

	lerrr := json.NewDecoder(r.Body).Decode(&l)
	if lerrr != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	services.GetShortAndInsert(r.Context(), l, u)
	log.Println("The long url is saved in the db")
	json.NewEncoder(w).Encode(services.GetUrl(r.Context(), l, u))

}

func Redirect(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt_token")

	if err != nil {
		fmt.Fprint(w, "No cookies found")
		return
	}

	fmt.Println("cookie found")

	tokenString := cookie.Value

	claims, lerr := utils.VerifyToken(tokenString)

	if lerr != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}

	u := claims.Uid
	log.Println("redirect funcion called")
	s := chi.URLParam(r, "s")
	l := services.GetLong(r.Context(), s, u)
	log.Println("the short is" + s)
	log.Println("the long is :=  " + l)
	http.Redirect(w, r, l, http.StatusMovedPermanently)
}
