package handler

import (
	"github.com/online_voting_service/gateway/genproto/public"
	"github.com/online_voting_service/gateway/genproto/vote"
	"google.golang.org/grpc"
)

type HandlerStruct struct {
	Public     public.PublicServiceClient
	Party      public.PartyServiceClient
	Candidate  vote.CandidateServiceClient
	Election   vote.ElectionServiceClient
	PublicVote vote.PublicVoteServiceClient
}

func NewHandler(connVote *grpc.ClientConn) *HandlerStruct {
	return &HandlerStruct{
		// Public:     public.NewPublicServiceClient(connPublic),
		// Party:      public.NewPartyServiceClient(connPublic),
		Candidate:  vote.NewCandidateServiceClient(connVote),
		Election:   vote.NewElectionServiceClient(connVote),
		PublicVote: vote.NewPublicVoteServiceClient(connVote),
	}
}
