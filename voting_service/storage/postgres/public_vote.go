package postgres

import "github.com/jackc/pgx/v5"

// PublicVote provides database operation for users
type PublicVote struct {
	Db *pgx.Conn
}

// NewUser creates a new instance of PublicVote
func NewPublicVote(db *pgx.Conn) *PublicVote {
	return &PublicVote{db}
}
