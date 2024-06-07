package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	public "github.com/myfirstgo/online_voting_service/public_service/genproto"
)

// PublicDb provides database operations for candidates
type PublicDb struct {
	Db *pgx.Conn
}

// NewPublic creates a new instance of PublicDb
func NewPublic(db *pgx.Conn) *PublicDb {
	return &PublicDb{Db: db}
}

// Create a new public record
func (pub *PublicDb) Create(ctx context.Context, publicReq *public.PublicCreate) error {
	query := `
		INSERT INTO public (id, name, last_name, phone, email, birthday, gender, party_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	birthday, err := time.Parse("2006-01-02", publicReq.Birthday)
	if err!= nil {
        return fmt.Errorf("failed to parse birthday: %v", err)
    }

	_, err = pub.Db.Exec(ctx, query, publicReq.Id, publicReq.Name, publicReq.LastName, publicReq.Phone, publicReq.Email, birthday, publicReq.Gender, publicReq.PartyId)
	if err != nil {
		return fmt.Errorf("failed to insert public record: %v", err)
	}
	return nil
}

// Update an existing public record
func (pub *PublicDb) Update(ctx context.Context, publicReq *public.PublicUpdate) error {
	query := `
		UPDATE public
		SET name = $1, last_name = $2, phone = $3, email = $4, birthday = $5, gender = $6, party_id = $7, updated_at = NOW()
		WHERE id = $8
	`
	_, err := pub.Db.Exec(ctx, query, publicReq.Name, publicReq.LastName, publicReq.Phone, publicReq.Email, publicReq.Birthday, publicReq.Gender, publicReq.PartyId, publicReq.Id)
	if err != nil {
		return fmt.Errorf("failed to update public record: %v", err)
	}
	return nil
}

// Delete an existing public record
func (pub *PublicDb) Delete(ctx context.Context, publicReq *public.PublicDelete) error {
	query := `
		UPDATE public
		SET deleted_at = NOW()
		WHERE id = $1
	`
	_, err := pub.Db.Exec(ctx, query, publicReq.Id)
	if err != nil {
		return fmt.Errorf("failed to delete public record: %v", err)
	}
	return nil
}

// GetById retrieves a public record by ID
func (pub *PublicDb) GetById(ctx context.Context, publicReq *public.PublicById) (*public.Public, error) {
	query := `
		SELECT id, name, last_name, phone, email, birthday, gender, party_id, created_at, updated_at, COALESCE(deleted_at, 0)
		FROM public
		WHERE id = $1 AND (deleted_at IS NULL OR deleted_at = 0)
	`
	row := pub.Db.QueryRow(ctx, query, publicReq.Id)

	var pubResp public.Public
	err := row.Scan(&pubResp.Id, &pubResp.Name, &pubResp.LastName, &pubResp.Phone, &pubResp.Email, &pubResp.Birthday, &pubResp.Gender, &pubResp.PartyId, &pubResp.CreatedAt, &pubResp.UpdatedAt, &pubResp.DeletedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get public record by ID: %v", err)
	}
	return &pubResp, nil
}

// GetAll retrieves all public records
func (pub *PublicDb) GetAll(ctx context.Context, publicReq *public.GetAllPublicReq) (*public.GetAllPublicRes, error) {
	query := `
		SELECT id, name, last_name, phone, email, birthday, gender, party_id, created_at, updated_at, COALESCE(deleted_at, 0)
		FROM public
		WHERE (party_id = $1 OR $1 = '') 
		AND (name = $2 OR $2 = '') 
		AND (last_name = $3 OR $3 = '') 
		AND (phone = $4 OR $4 = '') 
		AND (email = $5 OR $5 = '') 
		AND (birthday = $6 OR $6 = '') 
		AND (gender = $7 OR $7 = '')
		AND (deleted_at IS NULL OR deleted_at = 0)
	`
	rows, err := pub.Db.Query(ctx, query, publicReq.PartyId, publicReq.Name, publicReq.LastName, publicReq.Phone, publicReq.Email, publicReq.Birthday, publicReq.Gender)
	if err != nil {
		return nil, fmt.Errorf("failed to get all public records: %v", err)
	}
	defer rows.Close()

	var pubRes public.GetAllPublicRes
	for rows.Next() {
		var pubItem public.Public
		err := rows.Scan(&pubItem.Id, &pubItem.Name, &pubItem.LastName, &pubItem.Phone, &pubItem.Email, &pubItem.Birthday, &pubItem.Gender, &pubItem.PartyId, &pubItem.CreatedAt, &pubItem.UpdatedAt, &pubItem.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan public record: %v", err)
		}
		pubRes.Publics = append(pubRes.Publics, &pubItem)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return &pubRes, nil
}
