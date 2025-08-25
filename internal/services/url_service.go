package services

import (
	"URLSHORTENER/internal/models"
	"URLSHORTENER/internal/store"
	"context"
	"log"
)

func GetShortAndInsert(ctx context.Context, l models.Long) {
	s := String(5)
	log.Println(s)
	log.Println(l.Long)
	store.InsertURL(ctx, s, l.Long)
}

func GetUrl(ctx context.Context, l models.Long) string {
	m, err := store.FindShortFromLong(ctx, l.Long)
	if err != nil {
		log.Println("some error occured in saveInDB function")
	}
	s := "http://localhost:8080/" + m

	return s
}

func GetLong(ctx context.Context, s string) string {
	l, err := store.FindLongFromShort(ctx, s)
	if err != nil {
		log.Println("error occured in redirect function")
	}

	return l
}
