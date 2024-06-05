package storage

import (
	"context"
	"fmt"
	"log/slog"
	"vote/config"
	"vote/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	Candidate  *postgres.CandidateDb
	Election   *postgres.ElectionDb
	PublicVote *postgres.PublicVote
	Vote       *postgres.Vote
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
	cd := postgres.NewCandidate(db)
	el := postgres.NewElection(db)
	pbv := postgres.NewPublicVote(db)
	vt := postgres.NewVote(db)
	return &Storage{
		Candidate:  cd,
		Election:   el,
		PublicVote: pbv,
		Vote:       vt,
	}, err
}
