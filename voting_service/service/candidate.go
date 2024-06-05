package service

import (
	"context"
	"errors"
	"log/slog"
	vote "vote/genproto"
	"vote/storage"
)

type CandidateService struct {
	stg *storage.Storage
	vote.UnimplementedCandidateServiceServer
}

func NewCandidateService(stg *storage.Storage) *CandidateService {
	return &CandidateService{stg: stg}
}

func (c *CandidateService) Create(ctx context.Context, candidateReq *vote.CandidateCreate) (*vote.Candidate, error) {
	slog.Info("CreateCandidate Service", "candidate", candidateReq)
	valElection, err := c.stg.Election.ValidElectionDate(ctx, &candidateReq.ElectionId)
	if err != nil || valElection {
		slog.Error("this election date is not valid")
		if !valElection {
			return nil, errors.New("this election date is not valid")
		}
		return nil, err
	}
	candidateRes, err := c.stg.Candidate.Create(ctx, candidateReq)
	if err != nil {
		slog.Error("error while CreateCandidate Service", "err", err)
		return nil, err
	}

	return candidateRes, nil
}

func (c *CandidateService) Update(ctx context.Context, candidateReq *vote.CandidateUpdate) (*vote.Void, error) {
	slog.Info("UpdateCandidate Service", "candidate", candidateReq)
	valElection, err := c.stg.Election.ValidElectionDate(ctx, &candidateReq.ElectionId)
	if err != nil || valElection {
		slog.Error("this election date is not valid")
		if !valElection {
			return nil, errors.New("this election date is not valid")
		}
		return nil, err
	}
	err = c.stg.Candidate.Update(ctx, candidateReq)
	if err != nil {
		slog.Error("error while UpdateCandidate Service", "err", err)
		return &vote.Void{}, err
	}

	return &vote.Void{}, nil
}

func (c *CandidateService) Delete(ctx context.Context, candidateReq *vote.CandidateDelete) (*vote.Void, error) {
	slog.Info("DeleteCandidate Service", "candidate id", candidateReq.Id)
	err := c.stg.Candidate.Delete(ctx, candidateReq)
	if err != nil {
		slog.Error("error while DeleteCandidate Service", "err", err)
		return &vote.Void{}, err
	}

	return &vote.Void{}, nil
}

func (c *CandidateService) GetById(ctx context.Context, candidateReq *vote.CandidateById) (*vote.Candidate, error) {
	slog.Info("GetCandidateById Service", "candidate id", candidateReq.Id)
	candidateRes, err := c.stg.Candidate.GetById(ctx, candidateReq)
	if err != nil {
		slog.Error("error while GetCandidateById Service", "err", err)
		return nil, err
	}

	return candidateRes, nil
}

func (c *CandidateService) GetAll(ctx context.Context, candidateReq *vote.GetAllCandidateReq) (*vote.GetAllCandidateRes, error) {
	slog.Info("GetAllCandidate Service", "candidate req", candidateReq)
	candidateRes, err := c.stg.Candidate.GetAll(ctx, candidateReq)
	if err != nil {
		slog.Error("error while GetAllCandidate Service", "err", err)
		return nil, err
	}

	return candidateRes, nil
}
