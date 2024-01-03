package data

import (
	"context"
	"github.com/oli4maes/winedata/internal/validator"
	"time"
)

func (m WineModel) Insert(wine *Wine) error {
	query := `
		INSERT INTO wines (name)
		VALUES ($1)
		RETURNING id, version`

	args := []interface{}{wine.Name}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&wine.ID, &wine.Version)
}

func ValidateWine(v *validator.Validator, wine *Wine) {
	v.Check(wine.Name != "", "name", "must be provided")
}
