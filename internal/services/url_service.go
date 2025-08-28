package services

import (
	"URLSHORTENER/internal/models"
	"URLSHORTENER/internal/store"
	"context"
	"log"
)

func GetShortAndInsert(ctx context.Context, l models.Long, fk_user_id string) {
	s := String(5)
	log.Println(s)
	log.Println(l.Long)
	store.InsertURL(ctx, s, l.Long, fk_user_id)
}

func GetUrl(ctx context.Context, l models.Long, u string) string {
	m, err := store.FindShortFromLong(ctx, l.Long, u)
	if err != nil {
		log.Println("some error occured in saveInDB function")
	}
	s := "http://localhost:8080/redirect/" + m

	return s
}

func GetLong(ctx context.Context, s string, u string) string {
	l, err := store.FindLongFromShort(ctx, s, u)
	if err != nil {
		log.Println("error occured in redirect function")
	}

	return l
}
