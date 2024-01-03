package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

func (w WineModel) Get(id int64) (*Wine, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, name, version
		FROM wines
		WHERE id = $1`

	var wine Wine

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := w.DB.QueryRowContext(ctx, query, id).Scan(
		&wine.ID,
		&wine.Name,
		&wine.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &wine, nil
}
