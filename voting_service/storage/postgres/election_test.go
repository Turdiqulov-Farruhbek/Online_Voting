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
func newTestElection(t *testing.T) *ElectionDb {

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
	return &ElectionDb{Db: db}
}

// Create a test user object
func createTestElection() *vote.ElectionCreate {
	return &vote.ElectionCreate{
		Name:     "Test Election",
		OpenDate: "2024-06-07 15:04:05",
		EndDate:  "2024-06-08 15:04:05",
	}
}

func TestCreateElection(t *testing.T) {
	stgElection := newTestElection(t)

	testElection := createTestElection()

	electionRes, err := stgElection.Create(context.TODO(), testElection)
	if err != nil {
		t.Fatalf("Error creating election: %v", err)
	}

	assert.NotEmpty(t, electionRes.Id)
	assert.Equal(t, testElection.Name, electionRes.Name)
	assert.Equal(t, testElection.OpenDate, electionRes.OpenDate[:19]) // Trim time zone from open date
	assert.Equal(t, testElection.EndDate, electionRes.EndDate[:19])   // Trim time zone from end date
}

func TestGetElectionById(t *testing.T) {
	stgElection := newTestElection(t)

	testElection := createTestElection()
	createdElection, err := stgElection.Create(context.TODO(), testElection)
	if err != nil {
		t.Fatalf("Error creating election: %v", err)
	}

	election, err := stgElection.GetById(context.TODO(), &vote.ElectionById{Id: createdElection.Id})
	if err != nil {
		t.Fatalf("Error getting election by ID: %v", err)
	}
	log.Println(election.OpenDate, createdElection.OpenDate)
	assert.Equal(t, createdElection.Id, election.Id)
	assert.Equal(t, createdElection.Name, election.Name)
	assert.Equal(t, createdElection.OpenDate, election.OpenDate)
	assert.Equal(t, createdElection.EndDate, election.EndDate)
}

func TestDeleteElection(t *testing.T) {
	stgElection := newTestElection(t)

	testElection := createTestElection()
	createdElection, err := stgElection.Create(context.TODO(), testElection)
	if err != nil {
		t.Fatalf("Error creating election: %v", err)
	}

	err = stgElection.Delete(context.TODO(), &vote.ElectionDelete{Id: createdElection.Id})
	if err != nil {
		t.Fatalf("Error deleting election: %v", err)
	}

	_, err = stgElection.GetById(context.TODO(), &vote.ElectionById{Id: createdElection.Id})
	assert.Equal(t, ErrElectionNotFound, err)
}

func TestUpdateElection(t *testing.T) {
	stgElection := newTestElection(t)

	testElection := createTestElection()
	createdElection, err := stgElection.Create(context.TODO(), testElection)
	if err != nil {
		t.Fatalf("Error creating election: %v", err)
	}

	updatedElection := &vote.ElectionUpdate{
		Id:       createdElection.Id,
		Name:     "Updated Election",
		OpenDate: "2024-06-10 10:00:00",
		EndDate:  "2024-06-11 10:00:00",
	}

	err = stgElection.Update(context.TODO(), updatedElection)
	if err != nil {
		t.Fatalf("Error updating election: %v", err)
	}

	election, err := stgElection.GetById(context.TODO(), &vote.ElectionById{Id: createdElection.Id})
	if err != nil {
		t.Fatalf("Error getting election by ID: %v", err)
	}

	assert.Equal(t, updatedElection.Name, election.Name)
	assert.Equal(t, updatedElection.OpenDate, election.OpenDate[:19])
	assert.Equal(t, updatedElection.EndDate, election.EndDate[:19])
}

func TestGetAllElections(t *testing.T) {
	stgElection := newTestElection(t)

	// Create multiple elections
	for i := 0; i < 3; i++ {
		testElection := &vote.ElectionCreate{
			Name:     fmt.Sprintf("Election %d", i+1),
			OpenDate: "2024-06-07 15:04:05",
			EndDate:  "2024-06-08 15:04:05",
		}
		_, err := stgElection.Create(context.TODO(), testElection)
		if err != nil {
			t.Fatalf("Error creating election: %v", err)
		}
	}

	electionsRes, err := stgElection.GetAll(context.TODO(), &vote.GetAllElectionReq{})
	if err != nil {
		t.Fatalf("Error getting all elections: %v", err)
	}

	assert.GreaterOrEqual(t, electionsRes.Count, int64(3)) // At least 3 elections created

	// Check if all elections are retrieved
	for _, election := range electionsRes.Elections {
		log.Println(election)
	}
}

func TestGetCandidateVotes(t *testing.T) {
	stgElection := newTestElection(t)

	testElection := createTestElection()
	createdElection, err := stgElection.Create(context.TODO(), testElection)
	if err != nil {
		t.Fatalf("Error creating election: %v", err)
	}

	// Assuming you have a way to create candidates and votes associated with the election
	// (This part would need to be implemented in your candidate and vote services)
	// ... create candidates and votes for the election ...

	candidateVotesRes, err := stgElection.GetCandidateVotes(context.TODO(), &vote.GetCandidateVotesReq{Id: createdElection.Id})
	if err != nil {
		t.Fatalf("Error getting candidate votes: %v", err)
	}

	// Assert that the result contains at least one candidate with votes (assuming you created candidates)
	assert.GreaterOrEqual(t, len(candidateVotesRes.CandidateRes), 1)
}
