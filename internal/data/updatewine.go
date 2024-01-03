package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

func (w WineModel) Update(wine *Wine) error {
	query := `
		UPDATE wines
		SET name = $1, version = version + 1
		WHERE id = $2 AND version = $3
		RETURNING version`

	args := []interface{}{
		wine.Name,
		wine.ID,
		wine.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := w.DB.QueryRowContext(ctx, query, args...).Scan(&wine.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}
