package postgres

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	public "github.com/myfirstgo/online_voting_service/public_service/genproto"
)

func newTestPublic(t *testing.T) *PublicDb {

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
	return &PublicDb{Db: db}
}

func CreateTestPublic() *public.PublicCreate {
	return &public.PublicCreate{
		Id:       uuid.NewString(),
		Name:     "Public A",
		LastName: "LastName A",
		Phone:    "Phone A",
		Email:    "Email A",
		Birthday: "2024-06-07",
        Gender:   "f",
        PartyId:  "b384b482-b97c-4494-a6c6-e10b4ab11c0f",
	}
}

func TestCreatePublic(t *testing.T) {
	pub := newTestPublic(t)
	req := CreateTestPublic()

	err := pub.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to create public record: %v", err)	
	}

	res, err := pub.GetById(context.Background(), &public.PublicById{Id: req.Id})
	if err != nil {
		t.Fatalf("Failed to get public record by ID: %v", err)
	}
	assert.Equal(t, req.Id, res.Id)
	assert.Equal(t, req.Name, res.Name)
	assert.Equal(t, req.LastName, res.LastName)
	assert.Equal(t, req.Phone, res.Phone)
	assert.Equal(t, req.Email, res.Email)
}

func TestUpdatePublic(t *testing.T) {
	pub := newTestPublic(t)
    req := CreateTestPublic()

    err := pub.Create(context.Background(), req)
    if err!= nil {
        t.Fatalf("Failed to create public record: %v", err)
    }

    res, err := pub.GetById(context.Background(), &public.PublicById{Id: req.Id})
    if err!= nil {
        t.Fatalf("Failed to get public record by ID: %v", err)
    }
    assert.Equal(t, req.Id, res.Id)
    assert.Equal(t, req.Name, res.Name)
    assert.Equal(t, req.LastName, res.LastName)
    assert.Equal(t, req.Phone, res.Phone)
    assert.Equal(t, req.Email, res.Email)
}


func TestDeletePublic(t *testing.T) {
	pub := newTestPublic(t)
    req := CreateTestPublic()

    err := pub.Create(context.Background(), req)
    if err!= nil {
        t.Fatalf("Failed to create public record: %v", err)
    }

    err = pub.Delete(context.Background(), &public.PublicDelete{Id: req.Id})
    if err!= nil {
        t.Fatalf("Failed to delete public record: %v", err)
    }

    res, err := pub.GetById(context.Background(), &public.PublicById{Id: req.Id})
    if err!= nil {
        t.Fatalf("Failed to get public record by ID: %v", err)
    }
    assert.Equal(t, req.Id, res.Id)
    assert.Equal(t, req.Name, res.Name)
}


func TestGetPublicById(t *testing.T) {
	pub := newTestPublic(t)
    req := CreateTestPublic()

    err := pub.Create(context.Background(), req)
    if err!= nil {
        t.Fatalf("Failed to create public record: %v", err)
    }

    res, err := pub.GetById(context.Background(), &public.PublicById{Id: req.Id})
    if err!= nil {
        t.Fatalf("Failed to get public record by ID: %v", err)
    }
    assert.Equal(t, req.Id, res.Id)
    assert.Equal(t, req.Name, res.Name)
    assert.Equal(t, req.LastName, res.LastName)
    assert.Equal(t, req.Phone, res.Phone)
    assert.Equal(t, req.Email, res.Email)
}
