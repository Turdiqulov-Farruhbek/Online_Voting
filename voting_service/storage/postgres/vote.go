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

// ErrVoteNotFound represents an error when a vote is not found.
var ErrVoteNotFound = errors.New("vote not found")

// Vote provides database operation for votes
type Vote struct {
	Db *pgx.Conn
}

// NewVote creates a new instance of Vote
func NewVote(db *pgx.Conn) *Vote {
	return &Vote{Db: db}
}

// Create creates a new vote in the database
func (Vote *Vote) Create(ctx context.Context, voteReq *vote.VoteCreate) (*vote.Vote, error) {
	// Generate a new UUID for the vote
	voteID := uuid.New().String()

	query := `
		INSERT INTO 
			vote (
				id,
				candidate_id
			) 
		VALUES (
				$1, 
				$2
			)
		RETURNING 
			id, 
			candidate_id,
			created_at
	`
	var (
		dbVote    vote.Vote
		createdAt time.Time
	)

	err := Vote.Db.QueryRow(ctx, query,
		voteID,
		voteReq.CandidateId,
	).Scan(
		&dbVote.Id,
		&dbVote.CandidateId,
		&createdAt,
	)
	if err != nil {
		slog.Error("Error creating vote: %v", err)
		return nil, err
	}

	// Convert time.Time to string
	dbVote.CreatedAt = createdAt.String()

	return &dbVote, nil
}

// GetAll returns all votes (non-deleted) with votes and votes count
func (Vote *Vote) GetAll(ctx context.Context, voteReq *vote.GetAllVoteReq) (*vote.GetAllVoteRes, error) {
	var voteRes vote.GetAllVoteRes

	// Get votes with pagination
	query := `
		SELECT 
			id,
			candidate_id,
			created_at 
		FROM vote
	`

	// filter adds filter to the query
	// In here I use key and params to avoid sql injection
	params := make(map[string]interface{}, 0)

	var arr []interface{}
	filter := ""

	if len(voteReq.CandidateId) > 0 {
		params["candidate_id"] = voteReq.CandidateId
		filter += " WHERE candidate_id = :candidate_id"
	}
	query += filter
	query, arr = helper.ReplaceQueryParams(query, params)
	rows, err := Vote.Db.Query(ctx, query, arr...)
	if err != nil {
		slog.Error("Error getting votes: %v", err)
		return nil, fmt.Errorf("error getting votes: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			createdAt time.Time
			vote      vote.Vote
		)

		err = rows.Scan(
			&vote.Id,
			&vote.CandidateId,
			&createdAt,
		)
		if err != nil {
			slog.Error("Error scanning vote row: %v", err)
			return nil, fmt.Errorf("error scanning vote row: %w", err)
		}
		vote.CreatedAt = createdAt.String()
		voteRes.Votes = append(voteRes.Votes, &vote)
		voteRes.Count++
	}
	return &voteRes, nil
}

// GetById gets a vote by its ID
func (Vote *Vote) GetById(ctx context.Context, req *vote.VoteById) error {
	query := `
		SELECT
			id,
			candidate_id,
			created_at
		FROM 
			vote 
		WHERE 
			id = $1 
	`
	var (
		voteID      string
		candidateID string
		createdAt   time.Time
	)
	err := Vote.Db.QueryRow(ctx, query, req.Id).Scan(
		&voteID,
		&candidateID,
		&createdAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Error("Error getting vote by ID: %v", ErrVoteNotFound)
			return ErrVoteNotFound
		}
		slog.Error("Error getting vote by ID: %v", err)
		return err
	}

	return nil
}
