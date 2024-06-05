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
func newTestCandidate(t *testing.T) *CandidateDb {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		"sayyidmuhammad", // Replace with your database username
		"root",           // Replace with your database password
		"localhost",      // Replace with your database host
		5432,             // Replace with your database port
		"vote",           // Replace with your database name
	)

	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return &CandidateDb{Db: db}
}

// Create a test candidate object
func createTestCandidate() *vote.CandidateCreate {
	return &vote.CandidateCreate{
		ElectionId: "d2042066-6bee-4a3a-a4b6-52646e237773",
		PublicId:   uuid.NewString(),
	}
}

func TestCreateCandidate(t *testing.T) {
	stgCandidate := newTestCandidate(t)

	testCandidate := createTestCandidate()

	candidateRes, err := stgCandidate.Create(context.TODO(), testCandidate)
	if err != nil {
		t.Fatalf("Error creating candidate: %v", err)
	}

	assert.NotEmpty(t, candidateRes.Id)
	assert.Equal(t, testCandidate.ElectionId, candidateRes.ElectionId)
	assert.Equal(t, testCandidate.PublicId, candidateRes.PublicId)
}

func TestGetCandidateById(t *testing.T) {
	stgCandidate := newTestCandidate(t)

	testCandidate := createTestCandidate()
	createdCandidate, err := stgCandidate.Create(context.TODO(), testCandidate)
	if err != nil {
		t.Fatalf("Error creating candidate: %v", err)
	}

	candidate, err := stgCandidate.GetById(context.TODO(), &vote.CandidateById{Id: createdCandidate.Id})
	if err != nil {
		t.Fatalf("Error getting candidate by ID: %v", err)
	}

	assert.Equal(t, createdCandidate.Id, candidate.Id)
	assert.Equal(t, createdCandidate.ElectionId, candidate.ElectionId)
	assert.Equal(t, createdCandidate.PublicId, candidate.PublicId)
}

func TestDeleteCandidate(t *testing.T) {
	stgCandidate := newTestCandidate(t)

	testCandidate := createTestCandidate()
	createdCandidate, err := stgCandidate.Create(context.TODO(), testCandidate)
	if err != nil {
		t.Fatalf("Error creating candidate: %v", err)
	}

	err = stgCandidate.Delete(context.TODO(), &vote.CandidateDelete{Id: createdCandidate.Id})
	if err != nil {
		t.Fatalf("Error deleting candidate: %v", err)
	}

	_, err = stgCandidate.GetById(context.TODO(), &vote.CandidateById{Id: createdCandidate.Id})
	assert.Equal(t, ErrCandidateNotFound, err)
}

func TestUpdateCandidate(t *testing.T) {
	stgCandidate := newTestCandidate(t)

	testCandidate := createTestCandidate()
	createdCandidate, err := stgCandidate.Create(context.TODO(), testCandidate)
	if err != nil {
		t.Fatalf("Error creating candidate: %v", err)
	}

	updatedCandidate := &vote.CandidateUpdate{
		Id:         createdCandidate.Id,
		ElectionId: "d2042066-6bee-4a3a-a4b6-52646e237773", // Replace with a different valid election ID
		PublicId:   uuid.NewString(),
	}

	err = stgCandidate.Update(context.TODO(), updatedCandidate)
	if err != nil {
		t.Fatalf("Error updating candidate: %v", err)
	}

	candidate, err := stgCandidate.GetById(context.TODO(), &vote.CandidateById{Id: createdCandidate.Id})
	if err != nil {
		t.Fatalf("Error getting candidate by ID: %v", err)
	}

	assert.Equal(t, updatedCandidate.ElectionId, candidate.ElectionId)
	assert.Equal(t, updatedCandidate.PublicId, candidate.PublicId)
}

func TestGetAllCandidates(t *testing.T) {
	stgCandidate := newTestCandidate(t)

	// Create multiple candidates for the same election
	electionID := "d2042066-6bee-4a3a-a4b6-52646e237773" // Replace with a valid election ID

	for i := 0; i < 3; i++ {
		testCandidate := &vote.CandidateCreate{
			ElectionId: electionID,
			PublicId:   uuid.NewString(),
		}
		_, err := stgCandidate.Create(context.TODO(), testCandidate)
		if err != nil {
			t.Fatalf("Error creating candidate: %v", err)
		}
	}

	candidatesRes, err := stgCandidate.GetAll(context.TODO(), &vote.GetAllCandidateReq{ElectionId: electionID})
	if err != nil {
		t.Fatalf("Error getting all candidates: %v", err)
	}

	assert.GreaterOrEqual(t, candidatesRes.Count, int64(3)) // At least 3 candidates created

	// Check if all candidates are retrieved (you can add more specific assertions here)
	for _, candidate := range candidatesRes.Candidates {
		log.Println(candidate)
	}
}
