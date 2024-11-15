package service

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"image-processor/models"
	"math/rand"
	"net/http"
	"time"
)

type ImageProcessor struct {
	jobService *JobService
}

func NewImageProcessor() *ImageProcessor {
	return &ImageProcessor{}
}

func (p *ImageProcessor) ProcessJob(jobID string, visits []models.Visit) {
	var errors []models.JobError

	for _, visit := range visits {
		// Process each image URL
		for _, imageURL := range visit.ImageURLs {
			if err := p.processImage(imageURL); err != nil {
				errors = append(errors, models.JobError{
					StoreID: visit.StoreID,
					Error:   err.Error(),
				})
				p.jobService.UpdateJobStatus(jobID, models.StatusFailed, errors)
				return
			}
		}
	}

	p.jobService.UpdateJobStatus(jobID, models.StatusCompleted, nil)
}

func (p *ImageProcessor) processImage(url string) error {
	// Download image
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode image
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return err
	}

	// Calculate perimeter
	bounds := img.Bounds()
	perimeter := 2 * (bounds.Dx() + bounds.Dy())

	// Random sleep to simulate GPU processing
	sleepTime := 100 + rand.Intn(301) // 100-400ms
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)

	return nil
}