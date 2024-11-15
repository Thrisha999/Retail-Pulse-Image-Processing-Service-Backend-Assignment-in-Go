package models

import "time"

type JobStatus string

const (
	StatusOngoing   JobStatus = "ongoing"
	StatusCompleted JobStatus = "completed"
	StatusFailed    JobStatus = "failed"
)

type Job struct {
	ID      string
	Status  JobStatus
	Visits  []Visit
	Count   int
	Errors  []JobError
}

type Visit struct {
	StoreID   string   `json:"store_id"`
	ImageURLs []string `json:"image_url"`
	VisitTime string   `json:"visit_time"`
}

type JobError struct {
	StoreID string `json:"store_id"`
	Error   string `json:"error"`
}

type SubmitJobRequest struct {
	Count  int     `json:"count"`
	Visits []Visit `json:"visits"`
}

type SubmitJobResponse struct {
	JobID string `json:"job_id"`
}

type JobStatusResponse struct {
	Status JobStatus  `json:"status"`
	JobID  string    `json:"job_id"`
	Errors []JobError `json:"error,omitempty"`
}