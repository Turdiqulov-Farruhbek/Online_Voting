package postgres

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	party "github.com/myfirstgo/online_voting_service/public_service/genproto"
)

func newTestParty(t *testing.T) *PartyDb {

	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		"postgres",  // Replace with your database username
		"root",      // Replace with your database password
		"localhost", // Replace with your database host
		5432,        // Replace with your database port
		"vote_db",   // Replace with your database name
	)

	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return &PartyDb{Db: db}
}

func CreateTestParty() *party.PartyCreate {
	return &party.PartyCreate{
		Id:          uuid.NewString(),
		Name:        "Party A",
		Slogan:      "Slogan A",
		OpenedDate:  "2024-01-01",
		Description: "Description A",
	}

}

func TestCreateParty(t *testing.T) {
	part := newTestParty(t)
	req := CreateTestParty()

	err := part.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to create party record: %v", err)
	}

	res, err := part.GetById(context.Background(), &party.PartyById{Id: req.Id})
	if err != nil {
		t.Fatalf("Failed to get party record by ID: %v", err)
	}
	assert.Equal(t, req.Id, res.Id)
	assert.Equal(t, req.Name, res.Name)
	assert.Equal(t, req.Slogan, res.Slogan)
	assert.Equal(t, req.OpenedDate, res.OpenedDate)
	assert.Equal(t, req.Description, res.Description)

}

func TestDeleteParty(t *testing.T) {
	part := newTestParty(t)
	req := CreateTestParty()

	err := part.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to create party record: %v", err)
	}

	err = part.Delete(context.Background(), &party.PartyDelete{Id: req.Id})
	if err != nil {
		t.Fatalf("Failed to delete party record: %v", err)
	}

	res, err := part.GetById(context.Background(), &party.PartyById{Id: req.Id})
	if err != nil {
		t.Fatalf("Failed to get party record by ID: %v", err)
	}
	assert.Equal(t, int64(0), res.DeletedAt)

}

func TestUpdateParty(t *testing.T) {
	part := newTestParty(t)
	req := CreateTestParty()

	err := part.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to create party record: %v", err)
	}

	res, err := part.GetById(context.Background(), &party.PartyById{Id: req.Id})
	if err != nil {
		t.Fatalf("Failed to get party record by ID: %v", err)
	}
	assert.Equal(t, req.Id, res.Id)
	assert.Equal(t, req.Name, res.Name)
	assert.Equal(t, req.Slogan, res.Slogan)
	assert.Equal(t, req.OpenedDate, res.OpenedDate)
}

func TestGetAllParties(t *testing.T) {
	part := newTestParty(t)
	req := CreateTestParty()

	err := part.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to create party record: %v", err)
	}

	res, err := part.GetAll(context.Background(), &party.GetAllPartyRequest{})
	if err != nil {
		t.Fatalf("Failed to get party record by ID: %v", err)
	}
	assert.Equal(t, req.Id, res.Parties[0].Id)
	assert.Equal(t, req.Name, res.Parties[0].Name)
	assert.Equal(t, req.Slogan, res.Parties[0].Slogan)
	assert.Equal(t, req.OpenedDate, res.Parties[0].OpenedDate)
}

func TestGetPartyById(t *testing.T) {
	part := newTestParty(t)
	req := CreateTestParty()

	err := part.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to create party record: %v", err)
	}

	res, err := part.GetById(context.Background(), &party.PartyById{Id: req.Id})
	if err != nil {
		t.Fatalf("Failed to get party record by ID: %v", err)
	}
	assert.Equal(t, req.Id, res.Id)
	assert.Equal(t, req.Name, res.Name)
	assert.Equal(t, req.Slogan, res.Slogan)
	assert.Equal(t, req.OpenedDate, res.OpenedDate)
}
