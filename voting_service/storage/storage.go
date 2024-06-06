package storage

import (
	"context"

	vote "vote/genproto"
)

type StorageI interface {
	Election() ElectionI
	Candidate() CandidateI
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

type PublicVoteI interface {
	Create(ctx context.Context, in *vote.PublicVoteCreate) (*vote.PublicVoteRes, error)
	// Get a public vote by its ID
	GetByIdPublic(ctx context.Context, in *vote.PublicVoteById) (*vote.PublicVoteRes, error)
	// Get a vote by its ID
	GetByIdVote(ctx context.Context, in *vote.VoteById) (*vote.VoteRes, error)
	// Get all public votes
	GetAllPublic(ctx context.Context, in *vote.GetAllPublicVoteReq) (*vote.GetAllPublicVoteRes, error)
	// Get all public votes
	GetAllVote(ctx context.Context, in *vote.GetAllVoteReq) (*vote.GetAllVoteRes, error)
}
