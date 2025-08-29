package store

import (
	"context"
	"fmt"
)

func InsertURL(ctx context.Context, short string, long string, fk_user_id string) error {
	// 1️⃣ Check if the URL already exists for this user
	q1 := `SELECT COUNT(*) FROM public.url WHERE long = $1 AND fk_user_id = $2`
	var count int
	err := DB.QueryRow(ctx, q1, long, fk_user_id).Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to get row count: %w", err)
	}

	// 2️⃣ Insert only if no duplicate
	if count == 0 {
		query := `INSERT INTO public.url (short, long, fk_user_id) VALUES ($1, $2, $3)`
		_, err := DB.Exec(ctx, query, short, long, fk_user_id)
		if err != nil {
			return fmt.Errorf("count is 0 but could not insert: %w", err)
		}
	}

	return nil
}

func FindLongFromShort(ctx context.Context, short string, fk_user_id string) (string, error) {
	var long string
	query := `SELECT long FROM public.url WHERE short = $1 && fk_user_id = $2 `
	err := DB.QueryRow(ctx, query, short).Scan(&long)
	if err != nil {
		return "", err
	}

	return long, nil

}

func FindShortFromLong(ctx context.Context, long string, fk_user_id string) (string, error) {
	var short string
	query := `SELECT short FROM public.url WHERE long = $1 AND fk_user_id = $2`
	err := DB.QueryRow(ctx, query, long, fk_user_id).Scan(&short)
	if err != nil {
		return "", err
	}

	return short, nil

}
