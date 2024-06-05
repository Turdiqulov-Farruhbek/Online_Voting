package main

import (
	"log"
	"net"
	vote "vote/genproto"
	"vote/service"
	"vote/storage"

	"google.golang.org/grpc"
)

func main() {

	db, err := storage.DBConn()
	if err != nil {
		log.Fatal(err)
	}

	liss, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	vote.RegisterElectionServiceServer(s, service.NewElectionService(db))
	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
