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

// ErrElectionNotFound represents an error when a election is not found.
var ErrElectionNotFound = errors.New("election not found")

// ElectionDb provides database operation for elections
type ElectionDb struct {
	Db *pgx.Conn
}

// NewElection creates a new instance of ElectionDb
func NewElection(db *pgx.Conn) *ElectionDb {
	return &ElectionDb{Db: db}
}

// Create creates a new election in the database
func (Electiondb *ElectionDb) Create(ctx context.Context, election *vote.ElectionCreate) (*vote.Election, error) {
	// Generate a new UUID for the election
	electionID := uuid.New().String()

	// Parse the open date and end date from strings to time.Time
	openDate, err := time.Parse("2006-01-02 15:04:05", election.OpenDate)
	if err != nil {
		slog.Error("Error parsing open date string")
		return nil, err
	}

	endDate, err := time.Parse("2006-01-02 15:04:05", election.EndDate)
	if err != nil {
		slog.Error("Error parsing end date string")
		return nil, err
	}

	query := `
		INSERT INTO 
			election (
				id,
				name, 
				open_date, 
				end_date
			) 
		VALUES (
				$1, 
				$2, 
				$3,
				$4
			)
		RETURNING 
			id, 
			name,
			open_date,
			end_date,
			created_at,
			updated_at,
			deleted_at
	`
	var (
		dbElection vote.Election
		createdAt  time.Time
		updatedAt  time.Time
	)

	err = Electiondb.Db.QueryRow(ctx, query,
		electionID,
		election.Name,
		openDate,
		endDate,
	).Scan(
		&dbElection.Id,
		&dbElection.Name,
		&openDate,
		&endDate,
		&createdAt,
		&updatedAt,
		&dbElection.DeletedAt,
	)
	if err != nil {
		slog.Error("Error creating election: %v", err)
		return nil, err
	}

	// Convert time.Time to string
	dbElection.CreatedAt = createdAt.String()
	dbElection.UpdatedAt = updatedAt.String()
	dbElection.OpenDate = openDate.String()
	dbElection.EndDate = endDate.String()

	return &dbElection, nil
}

// Delete deletes an election by using deleted_at table(Soft delete).
func (Electiondb *ElectionDb) Delete(ctx context.Context, id *vote.ElectionDelete) error {
	query := `
		UPDATE 
			election 
		SET 
			deleted_at = $1 
		WHERE 
			id = $2
	`
	_, err := Electiondb.Db.Exec(ctx, query, time.Now().Unix(), id.Id)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Error("Error deleting election: %v", ErrElectionNotFound)
			return ErrElectionNotFound
		}
		slog.Error("Error deleting election: %v", err)
		return err
	}
	return nil
}

// GetById gets an election by its ID
func (Electiondb *ElectionDb) GetById(ctx context.Context, id *vote.ElectionById) (*vote.Election, error) {
	var (
		election  vote.Election
		openDate  time.Time
		endDate   time.Time
		createdAt time.Time
		updatedAt time.Time
	)

	query := `
		SELECT
			id,
			name, 
			open_date, 
			end_date,
			created_at,
			updated_at,
			deleted_at
		FROM 
			election 
		WHERE 
			id = $1 AND deleted_at = 0
	`
	err := Electiondb.Db.QueryRow(ctx, query, id.Id).Scan(
		&election.Id,
		&election.Name,
		&openDate,
		&endDate,
		&createdAt,
		&updatedAt,
		&election.DeletedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Error("Error getting election by id: %v", ErrElectionNotFound)
			return nil, ErrElectionNotFound
		}
		slog.Error("Error getting election by ID: %v", err)
		return nil, err
	}

	// Convert time.Time to string
	election.CreatedAt = createdAt.String()
	election.UpdatedAt = updatedAt.String()
	election.OpenDate = openDate.String()
	election.EndDate = endDate.String()

	return &election, nil
}

// GetAll returns all elections (non-deleted) with elections and elections count
func (Electiondb *ElectionDb) GetAll(ctx context.Context, electionReq *vote.GetAllElectionReq) (*vote.GetAllElectionRes, error) {
	var electionRes vote.GetAllElectionRes

	// Get elections with pagination
	query := `
		SELECT 
			id,
			name, 
			open_date, 
			end_date,
			created_at,
			updated_at,
			deleted_at 
		FROM election 
	`

	// filter adds filter to the query
	// In here I use key and params to avoid sql injection
	params := make(map[string]interface{}, 0)

	var arr []interface{}
	filter := `WHERE deleted_at = 0`

	if len(electionReq.Name) > 0 {
		params["name"] = "%" + electionReq.Name + "%"
		filter += " AND name ILIKE :name"
	}

	if len(electionReq.OpenDate) > 0 {
		openDate, err := time.Parse("2006-01-02 15:04:05", electionReq.OpenDate)
		if err != nil {
			slog.Error("Error parsing open date string")
			return nil, err
		}
		params["open_date"] = openDate
		filter += " AND open_date = :open_date"
	}

	if len(electionReq.EndDate) > 0 {
		endDate, err := time.Parse("2006-01-02 15:04:05", electionReq.EndDate)
		if err != nil {
			slog.Error("Error parsing end date string")
			return nil, err
		}
		params["end_date"] = endDate
		filter += " AND end_date = :end_date"
	}
	query += filter
	query, arr = helper.ReplaceQueryParams(query, params)
	rows, err := Electiondb.Db.Query(ctx, query, arr...)
	if err != nil {
		slog.Error("Error getting elections: %v", err)
		return nil, fmt.Errorf("error getting elections: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			createdAt time.Time
			updatedAt time.Time
			openDate  time.Time
			endDate   time.Time
			election  vote.Election
		)

		err = rows.Scan(
			&election.Id,
			&election.Name,
			&openDate,
			&endDate,
			&createdAt,
			&updatedAt,
			&election.DeletedAt,
		)
		if err != nil {
			slog.Error("Error scanning election row: %v", err)
			return nil, fmt.Errorf("error scanning election row: %w", err)
		}
		election.CreatedAt = createdAt.String()
		election.UpdatedAt = updatedAt.String()
		election.OpenDate = openDate.String()
		election.EndDate = endDate.String()
		electionRes.Elections = append(electionRes.Elections, &election)
		electionRes.Count++
	}
	return &electionRes, nil
}

// Update updates an existing election in the database
func (Electiondb *ElectionDb) Update(ctx context.Context, election *vote.ElectionUpdate) error {
	var args []interface{}
	count := 1
	query := `
		UPDATE 
			election 
		SET `
	filter := ``

	if len(election.Name) > 0 {
		filter += fmt.Sprintf(" name = $%d, ", count)
		args = append(args, election.Name)
		count++
	}

	if len(election.OpenDate) > 0 {
		openDate, err := time.Parse("2006-01-02 15:04:05", election.OpenDate)
		if err != nil {
			slog.Error("Error parsing open date string")
			return err
		}
		filter += fmt.Sprintf(" open_date = $%d, ", count)
		args = append(args, openDate)
		count++
	}

	if len(election.EndDate) > 0 {
		endDate, err := time.Parse("2006-01-02 15:04:05", election.EndDate)
		if err != nil {
			slog.Error("Error parsing end date string")
			return err
		}
		filter += fmt.Sprintf(" end_date = $%d, ", count)
		args = append(args, endDate)
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

	args = append(args, election.Id)
	query += filter

	_, err := Electiondb.Db.Exec(ctx, query, args...)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Error("Error updating election: %v", ErrElectionNotFound)
			return ErrElectionNotFound
		}
		slog.Error("Error updating election: %v", err)
		return err
	}
	return nil
}

// GetCandidateVotes gets the votes for each candidate in an election
func (Electiondb *ElectionDb) GetCandidateVotes(ctx context.Context, req *vote.GetCandidateVotesReq) (*vote.GetCandidateVotesRes, error) {
	var (
		candidateRes   vote.GetCandidateVotesRes
		candidateVotes []*vote.CandidateElectionRes
	)
	query := `
		SELECT 
			c.id, 
			COUNT(v.id) 
		FROM 
			candidate c
		LEFT JOIN 
			vote v ON c.id = v.candidate_id
		WHERE 
			v.election_id = $1 
		GROUP BY 
			c.id
	`
	rows, err := Electiondb.Db.Query(ctx, query, req.Id)
	if err != nil {
		slog.Error("Error getting candidate votes: %v", err)
		return nil, fmt.Errorf("error getting candidate votes: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			candidateID string
			count       int64
		)
		err = rows.Scan(&candidateID, &count)
		if err != nil {
			slog.Error("Error scanning candidate votes row: %v", err)
			return nil, fmt.Errorf("error scanning candidate votes row: %w", err)
		}
		candidateVotes = append(candidateVotes, &vote.CandidateElectionRes{
			CandidateId: candidateID,
			Count:       count,
		})
	}
	candidateRes.CandidateRes = candidateVotes
	return &candidateRes, nil
}

func (Electiondb *ElectionDb) ValidElectionDate(ctx context.Context, id *string) (bool, error) {
	query := `
		SELECT
			COUNT(1)
		FROM 
			election 
		WHERE 
			open_date > now()
		AND
			id = $1 
		AND 
			deleted_at = 0
		
	`
	var count int64
	err := Electiondb.Db.QueryRow(ctx, query, id).Scan(&count)
	if err != nil || count < 1 {
		return false, err
	}
	return true, nil
}
