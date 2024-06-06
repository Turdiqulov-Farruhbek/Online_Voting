package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"
	vote "vote/genproto"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// PublicVote provides database operation for public votes
type PublicVote struct {
	Db *pgx.Conn
}

// NewPublicVote creates a new pubvstance of PublicVote
func NewPublicVote(db *pgx.Conn) *PublicVote {
	return &PublicVote{Db: db}
}

// Create creates a new public vote and a corresponding vote in the database
func (pv *PublicVote) Create(ctx context.Context, in *vote.PublicVoteCreate) (*vote.PublicVoteRes, error) {
	tx, err := pv.Db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Generate a new UUID for the public vote if id is not exists
	if len(in.Id) == 0 {

		in.Id = uuid.New().String()
	}

	// Insert into public_vote table
	publicVoteQuery := `
        INSERT INTO public_vote (id, election_id, public_id)
        VALUES ($1, $2, $3)
        RETURNING id, election_id, public_id, created_at
    `
	var publicVoteRes vote.PublicVoteRes
	var createdAt time.Time
	err = tx.QueryRow(ctx, publicVoteQuery, in.Id, in.ElectionId, in.PublicId).Scan(
		&publicVoteRes.Id,
		&publicVoteRes.ElectionId,
		&publicVoteRes.PublicId,
		&createdAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert into public_vote: %w", err)
	}

	publicVoteRes.CreatedAt = createdAt.Format(time.RFC3339)

	// Insert into vote table
	voteQuery := `
        INSERT INTO vote (id, candidate_id)
        VALUES ($1, $2)
    `
	voteID := uuid.New().String()
	_, err = tx.Exec(ctx, voteQuery, voteID, in.CandidateId)
	if err != nil {
		return nil, fmt.Errorf("failed to insert into vote: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &publicVoteRes, nil
}

// GetByIdPublic gets a public_vote record by its ID and returns corresponding PublicVoteRes
func (pv *PublicVote) GetByIdPublic(ctx context.Context, in *vote.PublicVoteById) (*vote.PublicVoteRes, error) {
	var publicVoteRes vote.PublicVoteRes
	var createdAt time.Time

	query := `
        SELECT id, election_id, public_id, created_at
        FROM public_vote 
        WHERE id = $1
    `
	err := pv.Db.QueryRow(ctx, query, in.Id).Scan(
		&publicVoteRes.Id,
		&publicVoteRes.ElectionId,
		&publicVoteRes.PublicId,
		&createdAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("public_vote not found for ID: %s", in.Id)
		}
		return nil, fmt.Errorf("failed to fetch public_vote by ID: %w", err)
	}
	publicVoteRes.CreatedAt = createdAt.Format(time.RFC3339)

	return &publicVoteRes, nil
}

// GetByIdVote gets a vote record by its ID and returns corresponding VoteRes
func (pv *PublicVote) GetByIdVote(ctx context.Context, in *vote.VoteById) (*vote.VoteRes, error) {
	var voteRes vote.VoteRes
	var createdAt time.Time

	query := `
        SELECT id, created_at
        FROM vote 
        WHERE id = $1
    `
	err := pv.Db.QueryRow(ctx, query, in.Id).Scan(&voteRes.Id, &createdAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("vote not found for ID: %s", in.Id)
		}
		return nil, fmt.Errorf("failed to fetch vote by ID: %w", err)
	}
	voteRes.CreatedAt = createdAt.Format(time.RFC3339)
	return &voteRes, nil
}

// GetAllPublic gets all public votes based on optional filters
func (pv *PublicVote) GetAllPublic(ctx context.Context, in *vote.GetAllPublicVoteReq) (*vote.GetAllPublicVoteRes, error) {
	var publicVoteRes vote.GetAllPublicVoteRes
	query := `
        SELECT pv.id, pv.election_id, pv.public_id, pv.created_at
        FROM public_vote pv
    `
	var args []interface{}
	whereClause := ""
	if in.ElectionId != "" {
		whereClause += " WHERE pv.election_id = $1"
		args = append(args, in.ElectionId)
	}

	if in.PublicId != "" {
		if whereClause != "" {
			whereClause += " AND"
		} else {
			whereClause += " WHERE"
		}
		whereClause += " pv.public_id = $2"
		args = append(args, in.PublicId)
	}

	query += whereClause

	rows, err := pv.Db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch public votes: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var pv vote.PublicVoteRes
		var createdAt time.Time

		err := rows.Scan(
			&pv.Id,
			&pv.ElectionId,
			&pv.PublicId,
			&createdAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan public vote row: %w", err)
		}

		pv.CreatedAt = createdAt.Format(time.RFC3339)

		publicVoteRes.PublicVotes = append(publicVoteRes.PublicVotes, &pv)
		publicVoteRes.Count++
	}

	return &publicVoteRes, nil
}

// GetAllVote gets all votes with an optional candidate ID filter
func (pv *PublicVote) GetAllVote(ctx context.Context, in *vote.GetAllVoteReq) (*vote.GetAllVoteRes, error) {
	var voteRes vote.GetAllVoteRes
	query := `
        SELECT v.id, v.created_at
        FROM vote v
    `

	var args []interface{}
	if in.CandidateId != "" {
		query += " WHERE v.candidate_id = $1"
		args = append(args, in.CandidateId)
	}

	rows, err := pv.Db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch votes: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var v vote.VoteRes
		var createdAt time.Time
		err := rows.Scan(&v.Id, &createdAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan vote row: %w", err)
		}
		v.CreatedAt = createdAt.Format(time.RFC3339)
		voteRes.Votes = append(voteRes.Votes, &v)
		voteRes.Count++
	}

	return &voteRes, nil
}
