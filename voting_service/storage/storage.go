package storage

import (
	"context"

	vote "vote/genproto"
)

type StorageI interface {
	Election() ElectionI
	Candidate() CandidateI
	Vote() VoteI
	PublicVote() PublicVoteI
}

type ElectionI interface {
	Create(ctx context.Context, electionReq *vote.ElectionCreate) (*vote.Election, error)
	Update(ctx context.Context, electionReq *vote.ElectionUpdate) error
	Delete(ctx context.Context, electionReq *vote.ElectionDelete) error
	GetById(ctx context.Context, electionReq *vote.ElectionById) (*vote.Election, error)
	GetAll(ctx context.Context, electionReq *vote.GetAllElectionReq) (*vote.GetAllElectionRes, error)
	GetCandidateVotes(ctx context.Context, electionReq *vote.GetCandidateVotesReq) (*vote.GetCandidateVotesRes, error)
	ValidElectionDate(ctx context.Context, id *string) (bool, error)
}

type CandidateI interface {
	Create(ctx context.Context, candidateReq *vote.CandidateCreate) (*vote.Candidate, error)
	Update(ctx context.Context, candidateReq *vote.CandidateUpdate) error
	Delete(ctx context.Context, candidateReq *vote.CandidateDelete) error
	GetById(ctx context.Context, candidateReq *vote.CandidateById) (*vote.Candidate, error)
	GetAll(ctx context.Context, candidateReq *vote.GetAllCandidateReq) (*vote.GetAllCandidateRes, error)
}

type VoteI interface {
	Create(ctx context.Context, voteReq *vote.VoteCreate) (*vote.Vote, error)
	GetAll(ctx context.Context, voteReq *vote.GetAllVoteReq) (*vote.GetAllVoteRes, error)
	GetById(ctx context.Context, voteReq *vote.VoteById) error
}

type PublicVoteI interface {
	Create(ctx context.Context, publicVote *vote.PublicVoteCreate) (*vote.PublicVote, error)
	GetByPublicVoteId(ctx context.Context, req *vote.PublicVoteById) (*vote.PublicVote, error)
	GetAll(ctx context.Context, publicVoteReq *vote.GetAllPublicVoteReq) (*vote.GetAllPublicVoteRes, error)
}
