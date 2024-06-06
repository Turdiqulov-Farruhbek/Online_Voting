package main

import (
	"log"

	"github.com/online_voting_service/gateway/api"
	_ "github.com/online_voting_service/gateway/docs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// @title Online Voting System API
// @version 1.0
// @description API for managing online voting system resources
// @host localhost:8080
// @BasePath /api/v1
// @in header
// @name Authorization
func main() {
	// Connect to the Vote service
	connVote, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Vote service: %v", err)
	}
	defer connVote.Close()

	// Create and run the Gin router
	engine := api.NewGin(connVote)
	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
