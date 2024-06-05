package postgres

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"
	vote "vote/genproto"
	"vote/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// ErrPublicVoteNotFound represents an error when a public vote is not found.
var ErrPublicVoteNotFound = errors.New("public vote not found")

// PublicVote provides database operation for public votes
type PublicVote struct {
	Db *pgx.Conn
}

// NewPublicVote creates a new instance of PublicVote
func NewPublicVote(db *pgx.Conn) *PublicVote {
	return &PublicVote{Db: db}
}

// Create creates a new public vote in the database
func (PublicVote *PublicVote) Create(ctx context.Context, publicVote *vote.PublicVoteCreate) (*vote.PublicVote, error) {
	// Generate a new UUID for the public vote
	publicVoteID := uuid.New().String()

	query := `
		INSERT INTO 
			public_vote (
				id,
				election_id, 
				public_id
			) 
		VALUES (
				$1, 
				$2, 
				$3
			)
		RETURNING 
			id, 
			election_id,
			public_id,
			created_at
	`
	var (
		dbPublicVote vote.PublicVote
		createdAt    time.Time
	)

	err := PublicVote.Db.QueryRow(ctx, query,
		publicVoteID,
		publicVote.ElectionId,
		publicVote.PublicId,
	).Scan(
		&dbPublicVote.Id,
		&dbPublicVote.ElectionId,
		&dbPublicVote.PublicId,
		&createdAt,
	)
	if err != nil {
		slog.Error("Error creating public vote: %v", err)
		return nil, err
	}

	// Convert time.Time to string
	dbPublicVote.CreatedAt = createdAt.String()

	return &dbPublicVote, nil
}

// GetByPublicId gets a public vote by its public ID
func (PublicVote *PublicVote) GetByPublicVoteId(ctx context.Context, req *vote.PublicVoteById) (*vote.PublicVote, error) {
	var (
		publicVote vote.PublicVote
		createdAt  time.Time
	)

	query := `
		SELECT
			id,
			election_id,
			public_id,
			created_at
		FROM 
			public_vote 
		WHERE 
			id = $1 
	`
	err := PublicVote.Db.QueryRow(ctx, query, req.Id).Scan(
		&publicVote.Id,
		&publicVote.ElectionId,
		&publicVote.PublicId,
		&createdAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Error("Error getting public vote by public ID: %v", ErrPublicVoteNotFound)
			return nil, ErrPublicVoteNotFound
		}
		slog.Error("Error getting public vote by public ID: %v", err)
		return nil, err
	}

	// Convert time.Time to string
	publicVote.CreatedAt = createdAt.String()

	return &publicVote, nil
}

// GetAll returns all public votes (non-deleted) with public votes and public votes count
func (PublicVote *PublicVote) GetAll(ctx context.Context, publicVoteReq *vote.GetAllPublicVoteReq) (*vote.GetAllPublicVoteRes, error) {
	var publicVoteRes vote.GetAllPublicVoteRes

	// Get public votes with pagination
	query := `
		SELECT 
			id,
			election_id, 
			public_id,
			created_at 
		FROM public_vote
	`

	// filter adds filter to the query
	// In here I use key and params to avoid sql injection
	params := make(map[string]interface{}, 0)

	var arr []interface{}
	filter := ``

	if len(publicVoteReq.ElectionId) > 0 {
		params["election_id"] = publicVoteReq.ElectionId
		filter += " WHERE election_id = :election_id"
	}

	if len(publicVoteReq.PublicId) > 0 {
		params["public_id"] = publicVoteReq.PublicId
		if filter == "" {
			filter += " WHERE public_id = :public_id"
		} else {
			filter += " AND public_id = :public_id"
		}
	}
	query += filter
	query, arr = helper.ReplaceQueryParams(query, params)
	rows, err := PublicVote.Db.Query(ctx, query, arr...)
	if err != nil {
		slog.Error("Error getting public votes: %v", err)
		return nil, fmt.Errorf("error getting public votes: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			createdAt  time.Time
			publicVote vote.PublicVote
		)

		err = rows.Scan(
			&publicVote.Id,
			&publicVote.ElectionId,
			&publicVote.PublicId,
			&createdAt,
		)
		if err != nil {
			slog.Error("Error scanning public vote row: %v", err)
			return nil, fmt.Errorf("error scanning public vote row: %w", err)
		}
		publicVote.CreatedAt = createdAt.String()
		publicVoteRes.PublicVotes = append(publicVoteRes.PublicVotes, &publicVote)
		publicVoteRes.Count++
	}
	return &publicVoteRes, nil
}
