package service

import (
	"image-processor/models"
	"sync"

	"github.com/google/uuid"
)

type JobService struct {
	jobs map[string]*models.Job
	mu   sync.RWMutex
}

func NewJobService() *JobService {
	return &JobService{
		jobs: make(map[string]*models.Job),
	}
}

func (s *JobService) CreateJob(req models.SubmitJobRequest) string {
	jobID := uuid.New().String()
	
	job := &models.Job{
		ID:     jobID,
		Status: models.StatusOngoing,
		Visits: req.Visits,
		Count:  req.Count,
	}

	s.mu.Lock()
	s.jobs[jobID] = job
	s.mu.Unlock()

	return jobID
}

func (s *JobService) GetJob(jobID string) (*models.Job, bool) {
	s.mu.RLock()
	job, exists := s.jobs[jobID]
	s.mu.RUnlock()
	return job, exists
}

func (s *JobService) UpdateJobStatus(jobID string, status models.JobStatus, errors []models.JobError) {
	s.mu.Lock()
	if job, exists := s.jobs[jobID]; exists {
		job.Status = status
		job.Errors = errors
	}
	s.mu.Unlock()
}