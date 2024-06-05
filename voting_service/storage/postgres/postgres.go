package postgres

import (
	"context"
	"fmt"
	"log/slog"
	"vote/config"
	"vote/storage"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	Db          *pgx.Conn
	CandidateS  storage.CandidateI
	ElectionS   storage.ElectionI
	PublicVoteS storage.PublicVoteI
	VoteS       storage.VoteI
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
	cd := NewCandidate(db)
	el := NewElection(db)
	pbv := NewPublicVote(db)
	vt := NewVote(db)
	return &Storage{
		CandidateS:  cd,
		ElectionS:   el,
		PublicVoteS: pbv,
		VoteS:       vt,
	}, err
}

func (s *Storage) Election() storage.ElectionI {
	if s.ElectionS == nil {
		s.ElectionS = NewElection(s.Db)
	}
	return s.ElectionS
}

func (s *Storage) Candidate() storage.CandidateI {
	if s.CandidateS == nil {
		s.CandidateS = NewCandidate(s.Db)
	}
	return s.CandidateS
}

func (s *Storage) Vote() storage.VoteI {
	if s.VoteS == nil {
		s.VoteS = NewVote(s.Db)
	}
	return s.VoteS
}

func (s *Storage) PublicVote() storage.PublicVoteI {
	if s.PublicVoteS == nil {
		s.PublicVoteS = NewPublicVote(s.Db)
	}
	return s.PublicVoteS
}
