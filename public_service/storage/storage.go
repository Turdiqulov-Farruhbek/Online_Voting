package storage

import (
	"context"

	public "github.com/myfirstgo/online_voting_service/public_service/genproto"
)

type StorageI interface {
	Party() PartyI
	Public() PublicI
}
type PartyI interface {
	Create(ctx context.Context, partyReq *public.PartyCreate) error
	Update(ctx context.Context, partyReq *public.PartyUpdate) error
	Delete(ctx context.Context, partyReq *public.PartyDelete) error
	GetById(ctx context.Context, partyReq *public.PartyById) (*public.Party, error)
	GetAll(ctx context.Context, partyReq *public.GetAllPartyRequest) (*public.GetAllPartyResponse, error)
}

type PublicI interface {
	Create(ctx context.Context, publicReq *public.PublicCreate) error
	Update(ctx context.Context, publicReq *public.PublicUpdate) error
	Delete(ctx context.Context, publicReq *public.PublicDelete) error
	GetById(ctx context.Context, publicReq *public.PublicById) (*public.Public, error)
	GetAll(ctx context.Context, publicReq *public.GetAllPublicReq) (*public.GetAllPublicRes, error)
}
