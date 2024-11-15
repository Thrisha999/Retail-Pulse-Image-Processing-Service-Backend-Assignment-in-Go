package handlers

import (
	"image-processor/models"
	"image-processor/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	jobService     *service.JobService
	imageProcessor *service.ImageProcessor
}

func NewHandler(js *service.JobService, ip *service.ImageProcessor) *Handler {
	return &Handler{
		jobService:     js,
		imageProcessor: ip,
	}
}

func (h *Handler) SubmitJob(c *gin.Context) {
	var req models.SubmitJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate request
	if req.Count != len(req.Visits) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "count does not match number of visits"})
		return
	}

	// Create and process job
	jobID := h.jobService.CreateJob(req)
	go h.imageProcessor.ProcessJob(jobID, req.Visits)

	c.JSON(http.StatusCreated, models.SubmitJobResponse{JobID: jobID})
}

func (h *Handler) GetJobStatus(c *gin.Context) {
	jobID := c.Query("jobid")
	if jobID == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	job, exists := h.jobService.GetJob(jobID)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	response := models.JobStatusResponse{
		Status: job.Status,
		JobID:  job.ID,
	}

	if job.Status == models.StatusFailed {
		response.Errors = job.Errors
	}

	c.JSON(http.StatusOK, response)
}