package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	public "github.com/myfirstgo/online_voting_service/public_service/genproto"
)

// PartyDb provides database operation for candidates
type PartyDb struct {
	Db *pgx.Conn
}

// NewCandidate creates a new instance of PartyDb
func NewParty(db *pgx.Conn) *PartyDb {
	return &PartyDb{Db: db}
}

func (part *PartyDb) Create(ctx context.Context, partyReq *public.PartyCreate) error {
	return nil
}
func (part *PartyDb) Update(ctx context.Context, partyReq *public.PartyUpdate) error {
	return nil
}
func (part *PartyDb) Delete(ctx context.Context, partyReq *public.PartyDelete) error {
	return nil
}
func (part *PartyDb) GetById(ctx context.Context, partyReq *public.PartyById) (*public.Party, error) {
	return nil, nil
}
func (part *PartyDb) GetAll(ctx context.Context, partyReq *public.GetAllPartyRequest) (*public.GetAllPartyResponse, error) {
	return nil, nil
}
