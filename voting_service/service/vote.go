package service

import (
	"context"
	"log/slog"
	vote "vote/genproto"
	"vote/storage/postgres"
)

type VoteService struct {
	stg *postgres.Storage
	vote.UnimplementedVoteServiceServer
}

func NewVoteService(stg *postgres.Storage) *VoteService {
	return &VoteService{stg: stg}
}

func (v *VoteService) Create(ctx context.Context, voteReq *vote.VoteCreate) (*vote.Vote, error) {
	slog.Info("CreateVote Service", "vote", voteReq)
	voteRes, err := v.stg.VoteS.Create(ctx, voteReq)
	if err != nil {
		slog.Error("error while CreateVote Service", "err", err)
		return nil, err
	}

	return voteRes, nil
}

func (v *VoteService) GetAll(ctx context.Context, voteReq *vote.GetAllVoteReq) (*vote.GetAllVoteRes, error) {
	slog.Info("GetAllVote Service", "vote req", voteReq)
	voteRes, err := v.stg.VoteS.GetAll(ctx, voteReq)
	if err != nil {
		slog.Error("error while GetAllVote Service", "err", err)
		return nil, err
	}

	return voteRes, nil
}

func (v *VoteService) GetById(ctx context.Context, voteReq *vote.VoteById) (*vote.Void, error) {
	slog.Info("GetVoteById Service", "vote id", voteReq.Id)
	err := v.stg.VoteS.GetById(ctx, voteReq)
	if err != nil {
		slog.Error("error while GetVoteById Service", "err", err)
		return &vote.Void{}, err
	}

	return &vote.Void{}, nil
}
