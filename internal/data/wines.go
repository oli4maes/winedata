package data

import (
	"database/sql"
)

type Wine struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Version int    `json:"version"`
}

type WineModel struct {
	DB *sql.DB
}
