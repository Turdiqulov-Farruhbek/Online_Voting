package postgres

import "github.com/jackc/pgx/v5"

// Vote provides database operation for users
type Vote struct {
	Db *pgx.Conn
}

// NewUser creates a new instance of Vote
func NewVote(db *pgx.Conn) *Vote {
	return &Vote{db}
}
