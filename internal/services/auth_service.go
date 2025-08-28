package services

import (
	"URLSHORTENER/internal/models"
	"URLSHORTENER/internal/store"
	"URLSHORTENER/internal/utils"
	"context"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var cost int = bcrypt.DefaultCost

type contextKey string

const UserKey contextKey = "uuid"

func Register(ctx context.Context, u models.User) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), cost)

	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}

	u.Password = string(hashedPassword)
	store.CreateUser(ctx, u)

}

func Login(ctx context.Context, ul models.UserLogin) string {

	u := store.FindUserByUname(ctx, ul.Uname)
	if u.Uname == "" {
		log.Println("no such user is present")
		log.Panic("no user found")
	}

	hp := u.Password

	lerr := bcrypt.CompareHashAndPassword([]byte(hp), []byte(ul.Password))
	if lerr != nil {
		fmt.Println("password does not match")
		return ""
	}

	//create token logic

	ts, err := utils.CreateToken(u.Id.String())

	if err != nil {
		fmt.Println("No username found")
		return ""
	}

	return ts

}
