package postgres

import (
	"context"
	"fmt"
	"log"
	"testing"
	vote "vote/genproto"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

// Create a test database connection pool
func newTestPublicVote(t *testing.T) *PublicVote {
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
	return &PublicVote{Db: db}
}

// Create a test public vote object
func createTestPublicVote() *vote.PublicVoteCreate {
	return &vote.PublicVoteCreate{
		ElectionId: "d2042066-6bee-4a3a-a4b6-52646e237773", // Replace with a valid election ID
		PublicId:   uuid.NewString(),
	}
}

func TestCreatePublicVote(t *testing.T) {
	stgPublicVote := newTestPublicVote(t)

	testPublicVote := createTestPublicVote()

	publicVoteRes, err := stgPublicVote.Create(context.TODO(), testPublicVote)
	if err != nil {
		t.Fatalf("Error creating public vote: %v", err)
	}

	assert.NotEmpty(t, publicVoteRes.Id)
	assert.Equal(t, testPublicVote.ElectionId, publicVoteRes.ElectionId)
	assert.Equal(t, testPublicVote.PublicId, publicVoteRes.PublicId)
}

func TestGetPublicVoteById(t *testing.T) {
	stgPublicVote := newTestPublicVote(t)

	testPublicVote := createTestPublicVote()
	createdPublicVote, err := stgPublicVote.Create(context.TODO(), testPublicVote)
	if err != nil {
		t.Fatalf("Error creating public vote: %v", err)
	}

	publicVote, err := stgPublicVote.GetByPublicVoteId(context.TODO(), &vote.PublicVoteById{Id: createdPublicVote.Id})
	if err != nil {
		t.Fatalf("Error getting public vote by ID: %v", err)
	}

	assert.Equal(t, createdPublicVote.Id, publicVote.Id)
	assert.Equal(t, createdPublicVote.ElectionId, publicVote.ElectionId)
	assert.Equal(t, createdPublicVote.PublicId, publicVote.PublicId)
}

func TestGetAllPublicVotes(t *testing.T) {
	stgPublicVote := newTestPublicVote(t)

	// Create multiple public votes for the same election
	electionID := "d2042066-6bee-4a3a-a4b6-52646e237773"
	for i := 0; i < 3; i++ {
		testPublicVote := &vote.PublicVoteCreate{
			ElectionId: electionID,
			PublicId:   uuid.NewString(),
		}
		_, err := stgPublicVote.Create(context.TODO(), testPublicVote)
		if err != nil {
			t.Fatalf("Error creating public vote: %v", err)
		}
	}

	publicVotesRes, err := stgPublicVote.GetAll(context.TODO(), &vote.GetAllPublicVoteReq{ElectionId: electionID})
	if err != nil {
		t.Fatalf("Error getting all public votes: %v", err)
	}

	assert.GreaterOrEqual(t, publicVotesRes.Count, int32(3)) // At least 3 public votes created

	for _, publicVote := range publicVotesRes.PublicVotes {
		log.Println(publicVote)
	}
}
