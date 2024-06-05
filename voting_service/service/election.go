package service

import (
	"context"
	"log/slog"
	vote "vote/genproto"
	"vote/storage"
)

type ElectionService struct {
	stg *storage.Storage
	vote.UnimplementedElectionServiceServer
}

func NewElectionService(stg *storage.Storage) *ElectionService {
	return &ElectionService{stg: stg}
}

func (e *ElectionService) Create(ctx context.Context, electionReq *vote.ElectionCreate) (*vote.Election, error) {
	slog.Info("CreateElection Service", "election", electionReq)
	electionRes, err := e.stg.Election.Create(ctx, electionReq)
	if err != nil {
		slog.Error("error while CreateElection Service", "err", err)
		return nil, err
	}

	return electionRes, nil
}

func (e *ElectionService) Update(ctx context.Context, electionReq *vote.ElectionUpdate) (*vote.Void, error) {
	slog.Info("UpdateElection Service", "election", electionReq)
	err := e.stg.Election.Update(ctx, electionReq)
	if err != nil {
		slog.Error("error while UpdateElection Service", "err", err)
		return &vote.Void{}, err
	}

	return &vote.Void{}, nil
}

func (e *ElectionService) Delete(ctx context.Context, electionReq *vote.ElectionDelete) (*vote.Void, error) {
	slog.Info("DeleteElection Service", "election id", electionReq.Id)
	err := e.stg.Election.Delete(ctx, electionReq)
	if err != nil {
		slog.Error("error while DeleteElection Service", "err", err)
		return &vote.Void{}, err
	}

	return &vote.Void{}, nil
}

func (e *ElectionService) GetById(ctx context.Context, electionReq *vote.ElectionById) (*vote.Election, error) {
	slog.Info("GetElectionById Service", "election id", electionReq.Id)
	electionRes, err := e.stg.Election.GetById(ctx, electionReq)
	if err != nil {
		slog.Error("error while GetElectionById Service", "err", err)
		return nil, err
	}

	return electionRes, nil
}

func (e *ElectionService) GetAll(ctx context.Context, electionReq *vote.GetAllElectionReq) (*vote.GetAllElectionRes, error) {
	slog.Info("GetAllElection Service", "election req", electionReq)
	electionRes, err := e.stg.Election.GetAll(ctx, electionReq)
	if err != nil {
		slog.Error("error while GetAllElection Service", "err", err)
		return nil, err
	}

	return electionRes, nil
}

func (e *ElectionService) GetCandidateVoes(ctx context.Context, electionReq *vote.GetCandidateVotesReq) (*vote.GetCandidateVotesRes, error) {
	slog.Info("GetCandidateVotes Service", "election id", electionReq.Id)
	electionRes, err := e.stg.Election.GetCandidateVotes(ctx, electionReq)
	if err != nil {
		slog.Error("error while GetCandidateVotes Service", "err", err)
		return nil, err
	}

	return electionRes, nil
}
