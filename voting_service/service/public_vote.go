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

func (s *PublicVoteService) Create(ctx context.Context, req *vote.PublicVoteCreate) (*vote.PublicVoteRes, error) {
	slog.Info("PublicVoteService Create", "req", req)
	res, err := s.stg.PublicVoteS.Create(ctx, req)
	if err != nil {
		slog.Error("error while PublicVoteService Create", "err", err)
		return nil, err
	}
	return res, nil
}

func (s *PublicVoteService) GetByIdPublic(ctx context.Context, req *vote.PublicVoteById) (*vote.PublicVoteRes, error) {
	slog.Info("PublicVoteService GetByIdPublic", "req", req)
	res, err := s.stg.PublicVoteS.GetByIdPublic(ctx, req)
	if err != nil {
		slog.Error("error while PublicVoteService GetByIdPublic", "err", err)
		return nil, err
	}
	return res, nil
}

func (s *PublicVoteService) GetByIdVote(ctx context.Context, req *vote.VoteById) (*vote.VoteRes, error) {
	slog.Info("PublicVoteService GetByIdVote", "req", req)
	res, err := s.stg.PublicVoteS.GetByIdVote(ctx, req)
	if err != nil {
		slog.Error("error while PublicVoteService GetByIdVote", "err", err)
		return nil, err
	}
	return res, nil
}

func (s *PublicVoteService) GetAllPublic(ctx context.Context, req *vote.GetAllPublicVoteReq) (*vote.GetAllPublicVoteRes, error) {
	slog.Info("PublicVoteService GetAllPublic", "req", req)
	res, err := s.stg.PublicVoteS.GetAllPublic(ctx, req)
	if err != nil {
		slog.Error("error while PublicVoteService GetAllPublic", "err", err)
		return nil, err
	}
	return res, nil
}

func (s *PublicVoteService) GetAllVote(ctx context.Context, req *vote.GetAllVoteReq) (*vote.GetAllVoteRes, error) {
	slog.Info("PublicVoteService GetAllVote", "req", req)
	res, err := s.stg.PublicVoteS.GetAllVote(ctx, req)
	if err != nil {
		slog.Error("error while PublicVoteService GetAllVote", "err", err)
		return nil, err
	}
	return res, nil
}
