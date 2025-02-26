package main

import (
	"github.com/gin-gonic/gin"
	"image-converter/controllers"
	"image-converter/services"
)

func main() {
	r := gin.Default()

	// Initialize service and controller
	imageService := services.NewBimgService()
	imageController := controllers.NewImageController(imageService)

	// Define routes
	r.POST("/convert", imageController.ConvertToJPEG)

	// Start server
	r.Run(":8080")
}
