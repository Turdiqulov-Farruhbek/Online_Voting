package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/online_voting_service/gateway/genproto/vote"
)

// @Summary Create a new election
// @Description Creates a new election.
// @Tags Elections
// @Accept json
// @Produce json
// @Param election body vote.ElectionCreate true "Election data"
// @Success 200 {object} vote.Election
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /election [post]
func (h *HandlerStruct) CreateElectionHandler(c *gin.Context) {
	var (
		electionReq vote.ElectionCreate
		err         error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err = c.BindJSON(&electionReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Binding data: " + err.Error()})
		return
	}

	electionRes, err := h.Election.Create(ctx, &electionReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create election: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"election": electionRes})
}

// @Summary Update an existing election
// @Description Updates an existing election.
// @Tags Elections
// @Accept json
// @Produce json
// @Param election body vote.ElectionUpdate true "Election data"
// @Success 200 {object} vote.Void
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Resource not found"
// @Failure 500 {object} string "Internal server error"
// @Router /election [put]
func (h *HandlerStruct) UpdateElectionHandler(c *gin.Context) {
	var (
		electionReq vote.ElectionUpdate
		err         error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err = c.BindJSON(&electionReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Binding data: " + err.Error()})
		return
	}

	_, err = h.Election.Update(ctx, &electionReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update election: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Election updated successfully"})
}

// @Summary Delete an existing election
// @Description Deletes an existing election by ID.
// @Tags Elections
// @Accept json
// @Produce json
// @Param election body vote.ElectionDelete true "Election data"
// @Success 200 {object} vote.Void
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Resource not found"
// @Failure 500 {object} string "Internal server error"
// @Router /election [delete]
func (h *HandlerStruct) DeleteElectionHandler(c *gin.Context) {
	var (
		electionReq vote.ElectionDelete
		err         error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err = c.BindJSON(&electionReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Binding data: " + err.Error()})
		return
	}

	_, err = h.Election.Delete(ctx, &electionReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete election: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Election deleted successfully"})
}

// @Summary Get an election by its ID
// @Description Retrieves an election by its ID.
// @Tags Elections
// @Accept json
// @Produce json
// @Param id query string true "Election ID"
// @Success 200 {object} vote.Election
// @Failure 400 {object} string "Invalid request parameters"
// @Failure 404 {object} string "Resource not found"
// @Failure 500 {object} string "Internal server error"
// @Router /election/id [get]
func (h *HandlerStruct) GetElectionByIdHandler(c *gin.Context) {
	var (
		electionReq vote.ElectionById
		err         error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	electionReq.Id = c.Query("id")

	electionRes, err := h.Election.GetById(ctx, &electionReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get election: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"election": electionRes})
}

// @Summary Get all elections
// @Description Retrieves all elections.
// @Tags Elections
// @Accept json
// @Produce json
// @Param name query string false "Election name (optional)"
// @Param open_date query string false "Open date (optional)"
// @Param end_date query string false "End date (optional)"
// @Success 200 {object} vote.GetAllElectionRes
// @Failure 400 {object} string "Invalid request parameters"
// @Failure 500 {object} string "Internal server error"
// @Router /election/all [get]
func (h *HandlerStruct) GetAllElectionsHandler(c *gin.Context) {
	var (
		electionReq vote.GetAllElectionReq
		err         error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 50*time.Second)
	defer cancel()

	electionReq.Name = c.Query("name")
	electionReq.OpenDate = c.Query("open_date")
	electionReq.EndDate = c.Query("end_date")

	electionRes, err := h.Election.GetAll(ctx, &electionReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get elections: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"elections": electionRes})
}

// @Summary Get election results
// @Description Retrieves the results of an election by its ID.
// @Tags Elections
// @Accept json
// @Produce json
// @Param id query string true "Election ID"
// @Success 200 {object} vote.GetCandidateVotesRes
// @Failure 400 {object} string "Invalid request parameters"
// @Failure 500 {object} string "Internal server error"
// @Router /election/results [get]
func (h *HandlerStruct) GetCandidateVotesHandler(c *gin.Context) {
	var (
		electionReq vote.GetCandidateVotesReq
		err         error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	electionReq.Id = c.Query("id")

	electionRes, err := h.Election.GetCandidateVoes(ctx, &electionReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get election results: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": electionRes})
}
