package store

import (
	"context"
	"fmt"
)

func InsertURL(ctx context.Context, short string, long string) error {
	query := `INSERT INTO public.url (short, long) VALUES ($1, $2)`

	_, err := DB.Exec(ctx, query, short, long)
	if err != nil {
		return fmt.Errorf("failed to insert url: %w", err)
	}

	return nil
}

func FindLongFromShort(ctx context.Context, short string) (string, error) {
	var long string
	query := `SELECT long FROM public.url WHERE short = $1 `
	err := DB.QueryRow(ctx, query, short).Scan(&long)
	if err != nil {
		return "", err
	}

	return long, nil

}

func FindShortFromLong(ctx context.Context, long string) (string, error) {
	var short string
	query := `SELECT short FROM public.url WHERE long = $1 `
	err := DB.QueryRow(ctx, query, long).Scan(&short)
	if err != nil {
		return "", err
	}

	return short, nil

}
