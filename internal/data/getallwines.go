package data

import (
	"context"
	"fmt"
	"time"
)

func (w WineModel) GetAll(name string, filters Filters) ([]*Wine, Metadata, error) {
	query := fmt.Sprintf(`
        SELECT count(*) OVER(), id, name, version
        FROM wines
        WHERE  (to_tsvector('simple', name) @@ plainto_tsquery('simple', $1) OR $1 = '')
        ORDER BY %s %s, id ASC
        LIMIT $2 OFFSET $3`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{name, filters.limit(), filters.offset()}

	rows, err := w.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0

	wines := []*Wine{}

	for rows.Next() {
		var wine Wine

		err := rows.Scan(
			&totalRecords,
			&wine.ID,
			&wine.Name,
			&wine.Version,
		)

		if err != nil {
			return nil, Metadata{}, err
		}

		wines = append(wines, &wine)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return wines, metadata, nil
}
