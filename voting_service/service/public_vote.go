package service

import (
	"context"
	"log/slog"
	vote "vote/genproto"
	"vote/storage/postgres"
)

type PublicVoteService struct {
	stg *postgres.Storage
	vote.UnimplementedPublicVoteServiceServer
}

func NewPublicVoteService(stg *postgres.Storage) *PublicVoteService {
	return &PublicVoteService{stg: stg}
}

func (p *PublicVoteService) Create(ctx context.Context, voteReq *vote.PublicVoteCreate) (*vote.PublicVote, error) {
	slog.Info("CreatePublicVote Service", "public vote", voteReq)
	voteRes, err := p.stg.PublicVoteS.Create(ctx, voteReq)
	if err != nil {
		slog.Error("error while CreatePublicVote Service", "err", err)
		return nil, err
	}

	return voteRes, nil
}

func (p *PublicVoteService) GetAll(ctx context.Context, voteReq *vote.GetAllPublicVoteReq) (*vote.GetAllPublicVoteRes, error) {
	slog.Info("GetAllPublicVote Service", "public vote req", voteReq)
	voteRes, err := p.stg.PublicVoteS.GetAll(ctx, voteReq)
	if err != nil {
		slog.Error("error while GetAllPublicVote Service", "err", err)
		return nil, err
	}

	return voteRes, nil
}

func (p *PublicVoteService) GetById(ctx context.Context, voteReq *vote.PublicVoteById) (*vote.PublicVote, error) {
	slog.Info("GetPublicVoteById Service", "public vote id", voteReq.Id)
	voteRes, err := p.stg.PublicVoteS.GetByPublicVoteId(ctx, voteReq)
	if err != nil {
		slog.Error("error while GetPublicVoteById Service", "err", err)
		return nil, err
	}

	return voteRes, nil
}
