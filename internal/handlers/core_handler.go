package handlers

import (
	"database/sql"

	"github.com/tijanieneye10/restapi/internal/store"
)

type Handler struct {
	DB      *sql.DB
	Queries *store.Queries
}

func NewHandler(db *sql.DB, queries *store.Queries) *Handler {
	return &Handler{
		DB:      db,
		Queries: queries,
	}
}
