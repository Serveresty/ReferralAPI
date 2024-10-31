package repository

import "github.com/jackc/pgx/v5"

type RefRepository struct {
	db *pgx.Conn
}

func NewRefRepository(db *pgx.Conn) *RefRepository {
	return &RefRepository{db: db}
}
