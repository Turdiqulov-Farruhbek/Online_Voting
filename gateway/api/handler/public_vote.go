package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/online_voting_service/gateway/genproto/vote"
)

// @Summary Create a new public vote
// @Description Creates a new public vote.
// @Tags Public Votes
// @Accept json
// @Produce json
// @Param public_vote body vote.PublicVoteCreate true "Public vote data"
// @Success 200 {object} vote.PublicVoteRes
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /public_vote [post]
func (h *HandlerStruct) CreatePublicVoteHandler(c *gin.Context) {
	var (
		publicVoteReq vote.PublicVoteCreate
		err           error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err = c.BindJSON(&publicVoteReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Binding data: " + err.Error()})
		return
	}

	publicVoteRes, err := h.PublicVote.Create(ctx, &publicVoteReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create public vote: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"public_vote": publicVoteRes})
}

// @Summary Get a public vote by public vote ID
// @Description Retrieves a public vote by its ID.
// @Tags Public Votes
// @Accept json
// @Produce json
// @Param id query string true "Public vote ID"
// @Success 200 {object} vote.PublicVoteRes
// @Failure 400 {object} string "Invalid request parameters"
// @Failure 404 {object} string "Resource not found"
// @Failure 500 {object} string "Internal server error"
// @Router /public_vote/public/id [get]
func (h *HandlerStruct) GetPublicVoteByPublicIdHandler(c *gin.Context) {
	var (
		publicVoteReq vote.PublicVoteById
		err           error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	publicVoteReq.Id = c.Query("id")

	publicVoteRes, err := h.PublicVote.GetByIdPublic(ctx, &publicVoteReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get public vote: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"public_vote": publicVoteRes})
}

// @Summary Get a public vote by vote ID
// @Description Retrieves a public vote by its ID.
// @Tags Public Votes
// @Accept json
// @Produce json
// @Param id query string true "Vote ID"
// @Success 200 {object} vote.VoteRes
// @Failure 400 {object} string "Invalid request parameters"
// @Failure 404 {object} string "Resource not found"
// @Failure 500 {object} string "Internal server error"
// @Router /public_vote/vote/id [get]
func (h *HandlerStruct) GetPublicVoteByVoteIdHandler(c *gin.Context) {
	var (
		voteReq vote.VoteById
		err     error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	voteReq.Id = c.Query("id")

	VoteRes, err := h.PublicVote.GetByIdVote(ctx, &voteReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get vote: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"vote": VoteRes})
}

// @Summary Get all public votes
// @Description Retrieves all public votes.
// @Tags Public Votes
// @Accept json
// @Produce json
// @Param election_id query string false "Election ID (optional)"
// @Param public_id query string false "Public ID (optional)"
// @Success 200 {object} vote.GetAllPublicVoteRes
// @Failure 400 {object} string "Invalid request parameters"
// @Failure 500 {object} string "Internal server error"
// @Router /public_vote/public/all [get]
func (h *HandlerStruct) GetAllPublicVotesHandler(c *gin.Context) {
	var (
		publicVoteReq vote.GetAllPublicVoteReq
		err           error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 50*time.Second)
	defer cancel()

	publicVoteReq.ElectionId = c.Query("election_id")
	publicVoteReq.PublicId = c.Query("public_id")

	publicVoteRes, err := h.PublicVote.GetAllPublic(ctx, &publicVoteReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get public votes: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"public_votes": publicVoteRes})
}

// @Summary Get all votes
// @Description Retrieves all votes.
// @Tags Public Votes
// @Accept json
// @Produce json
// @Param candidate_id query string false "Candidate ID (optional)"
// @Success 200 {object} vote.GetAllVoteRes
// @Failure 400 {object} string "Invalid request parameters"
// @Failure 500 {object} string "Internal server error"
// @Router /public_vote/vote/all [get]
func (h *HandlerStruct) GetAllVotesHandler(c *gin.Context) {
	var (
		voteReq vote.GetAllVoteReq
		err     error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 50*time.Second)
	defer cancel()

	voteReq.CandidateId = c.Query("candidate_id")

	voteRes, err := h.PublicVote.GetAllVote(ctx, &voteReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get votes: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"votes": voteRes})
}
