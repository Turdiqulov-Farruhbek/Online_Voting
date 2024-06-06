package postgres

import (
	"context"
	"fmt"
	"testing"
	"time"

	vote "vote/genproto"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

// Create a test database connection pool for PublicVote
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
		ElectionId:  "d2042066-6bee-4a3a-a4b6-52646e237773",
		PublicId:    uuid.New().String(),
		CandidateId: "d63cf968-c2dd-49ab-82d2-443beac8dd50",
	}
}

// TestCreatePublicVote tests the Create method
func TestCreatePublicVote(t *testing.T) {
	stgPublicVote := newTestPublicVote(t)
	testPublicVote := createTestPublicVote()

	publicVoteRes, err := stgPublicVote.Create(context.Background(), testPublicVote)
	if err != nil {
		t.Fatalf("Error creating public vote: %v", err)
	}

	assert.NotEmpty(t, publicVoteRes.Id)
	assert.Equal(t, testPublicVote.ElectionId, publicVoteRes.ElectionId)
	assert.Equal(t, testPublicVote.PublicId, publicVoteRes.PublicId)
}

// TestGetByIdPublic tests the GetByIdPublic method
func TestGetByIdPublic(t *testing.T) {
	stgPublicVote := newTestPublicVote(t)
	testPublicVote := createTestPublicVote()

	createdPublicVote, err := stgPublicVote.Create(context.Background(), testPublicVote)
	if err != nil {
		t.Fatalf("Error creating public vote: %v", err)
	}

	publicVoteRes, err := stgPublicVote.GetByIdPublic(context.Background(), &vote.PublicVoteById{Id: createdPublicVote.Id})
	if err != nil {
		t.Fatalf("Error getting public vote by ID: %v", err)
	}

	assert.Equal(t, createdPublicVote.Id, publicVoteRes.Id)
	assert.Equal(t, createdPublicVote.ElectionId, publicVoteRes.ElectionId)
	assert.Equal(t, createdPublicVote.PublicId, publicVoteRes.PublicId)
}

// TestGetAllPublic tests the GetAllPublic method
func TestGetAllPublic(t *testing.T) {
	stgPublicVote := newTestPublicVote(t)
	testPublicVote := createTestPublicVote()

	_, err := stgPublicVote.Create(context.Background(), testPublicVote)
	if err != nil {
		t.Fatalf("Error creating public vote: %v", err)
	}

	publicVotes, err := stgPublicVote.GetAllPublic(context.Background(), &vote.GetAllPublicVoteReq{})
	if err != nil {
		t.Fatalf("Error getting all public votes: %v", err)
	}
	assert.GreaterOrEqual(t, len(publicVotes.PublicVotes), 1)

	publicVotesFiltered, err := stgPublicVote.GetAllPublic(context.Background(), &vote.GetAllPublicVoteReq{ElectionId: testPublicVote.ElectionId})
	if err != nil {
		t.Fatalf("Error getting all public votes with election ID filter: %v", err)
	}
	assert.GreaterOrEqual(t, len(publicVotesFiltered.PublicVotes), 1)
}

// TestGetAllVote tests the GetAllVote method
func TestGetAllVote(t *testing.T) {
	stgPublicVote := newTestPublicVote(t)
	testPublicVote := createTestPublicVote()

	_, err := stgPublicVote.Create(context.Background(), testPublicVote)
	if err != nil {
		t.Fatalf("Error creating public vote: %v", err)
	}

	time.Sleep(2 * time.Second)

	allVotes, err := stgPublicVote.GetAllVote(context.Background(), &vote.GetAllVoteReq{})
	if err != nil {
		t.Fatalf("Error getting all votes: %v", err)
	}
	assert.GreaterOrEqual(t, len(allVotes.Votes), 1)

	votesFiltered, err := stgPublicVote.GetAllVote(context.Background(), &vote.GetAllVoteReq{CandidateId: testPublicVote.CandidateId})
	if err != nil {
		t.Fatalf("Error getting all votes with candidate ID filter: %v", err)
	}
	assert.GreaterOrEqual(t, len(votesFiltered.Votes), 1)
}
