package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/online_voting_service/gateway/genproto/vote"
)

// @Summary Create a new candidate
// @Description Creates a new candidate for an election.
// @Tags Candidates
// @Accept json
// @Produce json
// @Param candidate body vote.CandidateCreate true "Candidate data"
// @Success 200 {object} vote.Candidate
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /candidate [post]
func (h *HandlerStruct) CreateCandidateHandler(c *gin.Context) {
	var (
		candidateReq vote.CandidateCreate
		err          error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err = c.BindJSON(&candidateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Binding data: " + err.Error()})
		return
	}

	candidateRes, err := h.Candidate.Create(ctx, &candidateReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create candidate: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"candidate": candidateRes})
}

// @Summary Update an existing candidate
// @Description Updates an existing candidate.
// @Tags Candidates
// @Accept json
// @Produce json
// @Param candidate body vote.CandidateUpdate true "Candidate data"
// @Success 200 {object} vote.Void
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Resource not found"
// @Failure 500 {object} string "Internal server error"
// @Router /candidate [put]
func (h *HandlerStruct) UpdateCandidateHandler(c *gin.Context) {
	var (
		candidateReq vote.CandidateUpdate
		err          error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err = c.BindJSON(&candidateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Binding data: " + err.Error()})
		return
	}

	_, err = h.Candidate.Update(ctx, &candidateReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update candidate: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Candidate updated successfully"})
}

// @Summary Delete an existing candidate
// @Description Deletes an existing candidate by ID.
// @Tags Candidates
// @Accept json
// @Produce json
// @Param candidate body vote.CandidateDelete true "Candidate data"
// @Success 200 {object} vote.Void
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Resource not found"
// @Failure 500 {object} string "Internal server error"
// @Router /candidate [delete]
func (h *HandlerStruct) DeleteCandidateHandler(c *gin.Context) {
	var (
		candidateReq vote.CandidateDelete
		err          error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err = c.BindJSON(&candidateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Binding data: " + err.Error()})
		return
	}

	_, err = h.Candidate.Delete(ctx, &candidateReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete candidate: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Candidate deleted successfully"})
}

// @Summary Get a candidate by its ID
// @Description Retrieves a candidate by its ID.
// @Tags Candidates
// @Accept json
// @Produce json
// @Param id query string true "Candidate ID"
// @Success 200 {object} vote.Candidate
// @Failure 400 {object} string "Invalid request parameters"
// @Failure 404 {object} string "Resource not found"
// @Failure 500 {object} string "Internal server error"
// @Router /candidate/id [get]
func (h *HandlerStruct) GetCandidateByIdHandler(c *gin.Context) {
	var (
		candidateReq vote.CandidateById
		err          error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	candidateReq.Id = c.Query("id")

	candidateRes, err := h.Candidate.GetById(ctx, &candidateReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get candidate: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"candidate": candidateRes})
}

// @Summary Get all candidates
// @Description Retrieves all candidates associated with an election.
// @Tags Candidates
// @Accept json
// @Produce json
// @Param election_id query string true "Election ID"
// @Param public_id query string false "Public ID (optional)"
// @Success 200 {object} vote.GetAllCandidateRes
// @Failure 400 {object} string "Invalid request parameters"
// @Failure 500 {object} string "Internal server error"
// @Router /candidate/all [get]
func (h *HandlerStruct) GetAllCandidatesHandler(c *gin.Context) {
	var (
		candidateReq vote.GetAllCandidateReq
		err          error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 50*time.Second)
	defer cancel()

	candidateReq.ElectionId = c.Query("election_id")
	candidateReq.PublicId = c.Query("public_id")

	candidateRes, err := h.Candidate.GetAll(ctx, &candidateReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get candidates: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"candidates": candidateRes})
}
