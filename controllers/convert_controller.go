package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/h2non/bimg"
	"image-converter/services"
)

type ImageController struct {
	service services.ImageService
}

func NewImageController(service services.ImageService) *ImageController {
	return &ImageController{service: service}
}

func (c *ImageController) ConvertToJPEG(ctx *gin.Context) {
	formart := ctx.Query("format")
	var imgFormat bimg.ImageType
	switch formart {
	case "webp":
		imgFormat = bimg.WEBP

	case "jpg", "jpeg":
		imgFormat = bimg.JPEG

	case "png":
		imgFormat = bimg.PNG

	case "tiff":
		imgFormat = bimg.TIFF

	case "avif":
		imgFormat = bimg.AVIF

	case "svg":
		imgFormat = bimg.SVG

	case "pdf":
		imgFormat = bimg.PDF

	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid File Formart"})

	}
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
	}

	files := form.File["images"]
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	converted, err := c.service.InitializeJobs(files, imgFormat)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Images converted successfully", "images": converted})
}
