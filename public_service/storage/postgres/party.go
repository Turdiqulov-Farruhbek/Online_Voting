package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	public "github.com/myfirstgo/online_voting_service/public_service/genproto"
)

// PartyDb provides database operations for parties
type PartyDb struct {
	Db *pgx.Conn
}

// NewParty creates a new instance of PartyDb
func NewParty(db *pgx.Conn) *PartyDb {
	return &PartyDb{Db: db}
}

// Create a new party record
func (party *PartyDb) Create(ctx context.Context, req *public.PartyCreate) error {
	query := `
		INSERT INTO party(id, name, slogan, opened_date, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
	`
	_, err := party.Db.Exec(ctx, query, req.Id, req.Name, req.Slogan, req.OpenedDate, req.Description)
	if err != nil {
		return fmt.Errorf("failed to insert party record: %v", err)
	}
	return nil
}

// Update an existing party record
func (party *PartyDb) Update(ctx context.Context, req *public.PartyUpdate) error {
	query := `
		UPDATE party
		SET name = $1, slogan = $2, opened_date = $3, description = $4, updated_at = NOW()
		WHERE id = $5
	`
	_, err := party.Db.Exec(ctx, query, req.Name, req.Slogan, req.OpenedDate, req.Description, req.Id)
	if err != nil {
		return fmt.Errorf("failed to update party record: %v", err)
	}
	return nil
}

// Delete an existing party record
func (party *PartyDb) Delete(ctx context.Context, req *public.PartyDelete) error {
	query := `
		UPDATE party_table
		SET deleted_at = NOW()
		WHERE id = $1
	`
	_, err := party.Db.Exec(ctx, query, req.Id)
	if err != nil {
		return fmt.Errorf("failed to delete party record: %v", err)
	}
	return nil
}

// GetById retrieves a party record by ID
func (party *PartyDb) GetById(ctx context.Context, req *public.PartyById) (*public.Party, error) {
	query := `
		SELECT id, name, slogan, opened_date, description, created_at, updated_at, COALESCE(deleted_at, 0)
		FROM party
		WHERE id = $1 AND (deleted_at IS NULL OR deleted_at = 0)
	`
	row := party.Db.QueryRow(ctx, query, req.Id)

	var res public.Party
	err := row.Scan(&res.Id, &res.Name, &res.Slogan, &res.OpenedDate, &res.Description, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get party record by ID: %v", err)
	}
	return &res, nil
}

// GetAll retrieves all party records
func (party *PartyDb) GetAll(ctx context.Context, req *public.GetAllPartyRequest) (*public.GetAllPartyResponse, error) {
	query := `
		SELECT id, name, slogan, opened_date, description, created_at, updated_at, COALESCE(deleted_at, 0)
		FROM party
		WHERE (opened_date = $1 OR $1 = '') 
		AND (name = $2 OR $2 = '') 
		AND (slogan = $3 OR $3 = '') 
		AND (description = $4 OR $4 = '')
		AND (deleted_at IS NULL OR deleted_at = 0)
	`
	rows, err := party.Db.Query(ctx, query, req.OpenedDate, req.Name, req.Slogan, req.Description)
	if err != nil {
		return nil, fmt.Errorf("failed to get all party records: %v", err)
	}
	defer rows.Close()

	var res public.GetAllPartyResponse
	for rows.Next() {
		var item public.Party
		err := rows.Scan(&item.Id, &item.Name, &item.Slogan, &item.OpenedDate, &item.Description, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan party record: %v", err)
		}
		res.Parties = append(res.Parties, &item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return &res, nil
}
