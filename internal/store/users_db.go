package store

import (
	"URLSHORTENER/internal/models"
	"context"
	"log"
	"time"

	"github.com/google/uuid"
)

func CreateUser(ctx context.Context, u models.User) {
	u.Id = uuid.New()
	u.CreatedAt = time.Now()
	query := `INSERT INTO public.users (id,uname,password, "createdAt") VALUES ($1, $2, $3, $4)`

	_, err := DB.Exec(ctx, query, u.Id, u.Uname, u.Password, u.CreatedAt)

	if err != nil {
		log.Println("Error in create user query", err)
		return
	}

	log.Println("user entered succesfully")

}

func FindUserByUname(ctx context.Context, uname string) models.User {
	var u models.User
	query := `SELECT id, uname, password, "createdAt" FROM public.users WHERE uname = $1;`
	err := DB.QueryRow(ctx, query, uname).Scan(&u.Id, &u.Uname, &u.Password, &u.CreatedAt)
	if err != nil {
		log.Println("some error ocuured in findUserByUname function", err)
	}

	return u
}
