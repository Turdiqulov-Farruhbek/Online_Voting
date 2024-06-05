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

// ErrCandidateNotFound represents an error when a candidate is not found.
var ErrCandidateNotFound = errors.New("candidate not found")

// CandidateDb provides database operation for candidates
type CandidateDb struct {
	Db *pgx.Conn
}

// NewCandidate creates a new instance of CandidateDb
func NewCandidate(db *pgx.Conn) *CandidateDb {
	return &CandidateDb{Db: db}
}

// Create creates a new candidate in the database
func (CandidateDb *CandidateDb) Create(ctx context.Context, candidate *vote.CandidateCreate) (*vote.Candidate, error) {
	// Generate a new UUID for the candidate
	candidateID := uuid.New().String()

	query := `
		INSERT INTO 
			candidate (
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
			created_at,
			updated_at,
			deleted_at
	`
	var (
		dbCandidate vote.Candidate
		createdAt   time.Time
		updatedAt   time.Time
	)

	err := CandidateDb.Db.QueryRow(ctx, query,
		candidateID,
		candidate.ElectionId,
		candidate.PublicId,
	).Scan(
		&dbCandidate.Id,
		&dbCandidate.ElectionId,
		&dbCandidate.PublicId,
		&createdAt,
		&updatedAt,
		&dbCandidate.DeletedAt,
	)
	if err != nil {
		slog.Error("Error creating candidate: %v", err)
		return nil, err
	}

	// Convert time.Time to string
	dbCandidate.CreatedAt = createdAt.String()
	dbCandidate.UpdatedAt = updatedAt.String()

	return &dbCandidate, nil
}

// Delete deletes a candidate by using deleted_at table(Soft delete).
func (CandidateDb *CandidateDb) Delete(ctx context.Context, id *vote.CandidateDelete) error {
	query := `
		UPDATE 
			candidate 
		SET 
			deleted_at = $1 
		WHERE 
			id = $2
	`
	_, err := CandidateDb.Db.Exec(ctx, query, time.Now().Unix(), id.Id)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Error("Error deleting candidate: %v", ErrCandidateNotFound)
			return ErrCandidateNotFound
		}
		slog.Error("Error deleting candidate: %v", err)
		return err
	}
	return nil
}

// GetById gets a candidate by its ID
func (CandidateDb *CandidateDb) GetById(ctx context.Context, id *vote.CandidateById) (*vote.Candidate, error) {
	var (
		candidate vote.Candidate
		createdAt time.Time
		updatedAt time.Time
	)

	query := `
		SELECT
			id,
			election_id, 
			public_id,
			created_at,
			updated_at,
			deleted_at
		FROM 
			candidate 
		WHERE 
			id = $1 AND deleted_at = 0
	`
	err := CandidateDb.Db.QueryRow(ctx, query, id.Id).Scan(
		&candidate.Id,
		&candidate.ElectionId,
		&candidate.PublicId,
		&createdAt,
		&updatedAt,
		&candidate.DeletedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Error("Error getting candidate by id: %v", ErrCandidateNotFound)
			return nil, ErrCandidateNotFound
		}
		slog.Error("Error getting candidate by ID: %v", err)
		return nil, err
	}

	// Convert time.Time to string
	candidate.CreatedAt = createdAt.String()
	candidate.UpdatedAt = updatedAt.String()

	return &candidate, nil
}

// GetAll returns all candidates (non-deleted) with candidates and candidates count
func (CandidateDb *CandidateDb) GetAll(ctx context.Context, candidateReq *vote.GetAllCandidateReq) (*vote.GetAllCandidateRes, error) {
	var candidateRes vote.GetAllCandidateRes

	// Get candidates with pagination
	query := `
		SELECT 
			id,
			election_id, 
			public_id,
			created_at,
			updated_at,
			deleted_at 
		FROM candidate 
	`

	// filter adds filter to the query
	// In here I use key and params to avoid sql injection
	params := make(map[string]interface{}, 0)

	var arr []interface{}
	filter := `WHERE deleted_at = 0`

	if len(candidateReq.ElectionId) > 0 {
		params["election_id"] = candidateReq.ElectionId
		filter += " AND election_id = :election_id"
	}

	if len(candidateReq.PublicId) > 0 {
		params["public_id"] = candidateReq.PublicId
		filter += " AND public_id = :public_id"
	}
	query += filter
	query, arr = helper.ReplaceQueryParams(query, params)
	rows, err := CandidateDb.Db.Query(ctx, query, arr...)
	if err != nil {
		slog.Error("Error getting candidates: %v", err)
		return nil, fmt.Errorf("error getting candidates: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			createdAt time.Time
			updatedAt time.Time
			candidate vote.Candidate
		)

		err = rows.Scan(
			&candidate.Id,
			&candidate.ElectionId,
			&candidate.PublicId,
			&createdAt,
			&updatedAt,
			&candidate.DeletedAt,
		)
		if err != nil {
			slog.Error("Error scanning candidate row: %v", err)
			return nil, fmt.Errorf("error scanning candidate row: %w", err)
		}
		candidate.CreatedAt = createdAt.String()
		candidate.UpdatedAt = updatedAt.String()
		candidateRes.Candidates = append(candidateRes.Candidates, &candidate)
		candidateRes.Count++
	}
	return &candidateRes, nil
}

// Update updates an existing candidate in the database
func (CandidateDb *CandidateDb) Update(ctx context.Context, candidate *vote.CandidateUpdate) error {
	var args []interface{}
	count := 1
	query := `
		UPDATE 
			candidate 
		SET `
	filter := ``

	if len(candidate.ElectionId) > 0 {
		filter += fmt.Sprintf(" election_id = $%d, ", count)
		args = append(args, candidate.ElectionId)
		count++
	}

	if len(candidate.PublicId) > 0 {
		filter += fmt.Sprintf(" public_id = $%d, ", count)
		args = append(args, candidate.PublicId)
		count++
	}

	if filter == "" {
		slog.Error("No fields provided for update.")
		return errors.New("no fields provided for update")
	}

	filter += fmt.Sprintf(`
			updated_at = NOW()
		WHERE
			id = $%d
		AND
			deleted_at = 0
			`, count)

	args = append(args, candidate.Id)
	query += filter

	_, err := CandidateDb.Db.Exec(ctx, query, args...)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Error("Error updating candidate: %v", ErrCandidateNotFound)
			return ErrCandidateNotFound
		}
		slog.Error("Error updating candidate: %v", err)
		return err
	}
	return nil
}
