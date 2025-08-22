package services

import (
	"URLSHORTENER/internal/models"
	"URLSHORTENER/internal/store"
	"context"
	"fmt"
	"log"
)

func InsertURL(short string, long string) error {
	query := `INSERT INTO public.url (short, long) VALUES ($1, $2)`

	_, err := store.DB.Exec(context.Background(), query, short, long)
	if err != nil {
		return fmt.Errorf("failed to insert url: %w", err)
	}

	return nil
}

func FindLongFromShort(short string) (string, error) {
	var long string
	query := `SELECT long FROM public.url WHERE short = $1 `
	err := store.DB.QueryRow(context.Background(), query, short).Scan(&long)
	if err != nil {
		return "", err
	}

	return long, nil

}

func FindShortFromLong(long string) (string, error) {
	var short string
	query := `SELECT short FROM public.url WHERE long = $1 `
	err := store.DB.QueryRow(context.Background(), query, long).Scan(&short)
	if err != nil {
		return "", err
	}

	return short, nil

}

func GetShortAndInsert(l models.Long) {
	s := String(5)
	log.Println(s)
	log.Println(l.Long)
	InsertURL(s, l.Long)
}
