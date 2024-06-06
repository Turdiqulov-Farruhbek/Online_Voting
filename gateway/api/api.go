package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/online_voting_service/gateway/api/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

func NewGin(connVote *grpc.ClientConn) *gin.Engine {

	// handler support database connection
	handler := handler.NewHandler(connVote)

	// Connecting to gin router
	router := gin.Default()
	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// Election routes
	election := router.Group("api/v1/election")
	{
		election.POST("", handler.CreateElectionHandler)            // POST /api/v1/election
		election.PUT("", handler.UpdateElectionHandler)             // PUT /api/v1/election
		election.DELETE("", handler.DeleteElectionHandler)          // DELETE /api/v1/election
		election.POST("/id", handler.GetElectionByIdHandler)        // GET /api/v1/election/id
		election.POST("/all", handler.GetAllElectionsHandler)       // GET /api/v1/election/all
		election.POST("/results", handler.GetCandidateVotesHandler) // POST /api/v1/election/results
	}

	candidate := router.Group("/api/v1/candidate")
	{
		candidate.POST("", handler.CreateCandidateHandler)      // POST /api/v1/candidate
		candidate.PUT("", handler.UpdateCandidateHandler)       // PUT /api/v1/candidate
		candidate.DELETE("", handler.DeleteCandidateHandler)    // DELETE /api/v1/candidate
		candidate.POST("/id", handler.GetCandidateByIdHandler)  // POST /api/v1/candidate/id
		candidate.POST("/all", handler.GetAllCandidatesHandler) // POST /api/v1/candidate/all
	}

	return router
}
