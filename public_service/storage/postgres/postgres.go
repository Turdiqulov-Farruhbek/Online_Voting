package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/myfirstgo/online_voting_service/public_service/config"
	"github.com/myfirstgo/online_voting_service/public_service/storage"
)

type Storage struct {
	Db      *pgx.Conn
	PartyS  storage.PartyI
	PublicS storage.PublicI
}

func DBConn() (*Storage, error) {
	var (
		db  *pgx.Conn
		err error
	)
	// Get postgres connection data from .env file
	cfg := config.Load()
	dbCon := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)

	// Connecting to postgres
	db, err = pgx.Connect(context.Background(), dbCon)
	if err != nil {
		slog.Warn("Unable to connect to database:", err)
	}
	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	partDb := NewParty(db)
	publicDb := NewPublic(db)
	return &Storage{
		Db:      db,
		PartyS:  partDb,
		PublicS: publicDb,
	}, err
}
func (s *Storage) Party() *storage.PartyI {
	if s.PartyS == nil {
		s.PartyS = NewParty(s.Db)
	}
	return &s.PartyS
}
func (s *Storage) Public() storage.PublicI {
	if s.PublicS == nil {
		s.PublicS = NewPublic(s.Db)
	}
	return s.PublicS
}
