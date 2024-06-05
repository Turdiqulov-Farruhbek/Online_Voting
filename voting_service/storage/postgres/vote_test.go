package postgres

import (
	"context"
	"fmt"
	"log"
	"testing"
	vote "vote/genproto"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

// Create a test database connection pool
func newTestVote(t *testing.T) *Vote {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		"sayyidmuhammad",
		"root",
		"localhost",
		5432,
		"vote",
	)

	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return &Vote{Db: db}
}

// Create a test vote object
func createTestVote(candidateID string) *vote.VoteCreate {
	return &vote.VoteCreate{
		CandidateId: candidateID,
	}
}

func TestCreateVote(t *testing.T) {
	stgVote := newTestVote(t)

	candidateID := "d63cf968-c2dd-49ab-82d2-443beac8dd50"

	testVote := createTestVote(candidateID)

	voteRes, err := stgVote.Create(context.TODO(), testVote)
	if err != nil {
		t.Fatalf("Error creating vote: %v", err)
	}

	assert.NotEmpty(t, voteRes.Id)
	assert.Equal(t, testVote.CandidateId, voteRes.CandidateId)
}

func TestGetAllVotes(t *testing.T) {
	stgVote := newTestVote(t)

	candidateID := "d63cf968-c2dd-49ab-82d2-443beac8dd50"

	// Create multiple votes for the same candidate
	for i := 0; i < 3; i++ {
		testVote := createTestVote(candidateID)
		_, err := stgVote.Create(context.TODO(), testVote)
		if err != nil {
			t.Fatalf("Error creating vote: %v", err)
		}
	}

	votesRes, err := stgVote.GetAll(context.TODO(), &vote.GetAllVoteReq{CandidateId: candidateID})
	if err != nil {
		t.Fatalf("Error getting all votes: %v", err)
	}

	assert.GreaterOrEqual(t, votesRes.Count, int32(3))

	for _, vote := range votesRes.Votes {
		log.Println(vote)
	}
}

func TestGetVoteById(t *testing.T) {
	stgVote := newTestVote(t)

	candidateID := "d63cf968-c2dd-49ab-82d2-443beac8dd50"
	testVote := createTestVote(candidateID)

	createdVote, err := stgVote.Create(context.TODO(), testVote)
	if err != nil {
		t.Fatalf("Error creating vote: %v", err)
	}

	err = stgVote.GetById(context.TODO(), &vote.VoteById{Id: createdVote.Id})
	if err != nil {
		t.Fatalf("Error getting vote by ID: %v", err)
	}
}
