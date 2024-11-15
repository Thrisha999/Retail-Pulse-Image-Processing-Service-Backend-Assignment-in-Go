package main

import (
	"image-processor/handlers"
	"image-processor/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize services
	jobService := service.NewJobService()
	imageProcessor := service.NewImageProcessor()

	// Initialize handler with services
	handler := handlers.NewHandler(jobService, imageProcessor)

	// Setup router
	router := gin.Default()

	// Routes
	router.POST("/api/submit", handler.SubmitJob)
	router.GET("/api/status", handler.GetJobStatus)

	// Start server
	log.Fatal(router.Run(":8080"))
}