package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	public "github.com/myfirstgo/online_voting_service/public_service/genproto"
)

// PublicDb provides database operation for candidates
type PublicDb struct {
	Db *pgx.Conn
}

// NewCandidate creates a new instance of PublicDb
func NewPublic(db *pgx.Conn) *PublicDb {
	return &PublicDb{Db: db}
}
func (pub *PublicDb) Create(ctx context.Context, publicReq *public.PublicCreate) error {
	return nil
}
func (pub *PublicDb) Update(ctx context.Context, publicReq *public.PublicUpdate) error {
	return nil
}
func (pub *PublicDb) Delete(ctx context.Context, publicReq *public.PublicDelete) error {
	return nil
}
func (pub *PublicDb) GetById(ctx context.Context, publicReq *public.PublicById) (*public.Public, error) {
	return nil, nil
}
func (pub *PublicDb) GetAll(ctx context.Context, publicReq *public.GetAllPublicReq) (*public.GetAllPublicRes, error) {
	return nil, nil
}
