package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/h2non/bimg"
)

type ImageService interface {
	ConvertToJPEG(file []*multipart.FileHeader) ([]string, error)
}

// BimgService struct implementing ImageService
type BimgService struct{}

// NewBimgService returns an instance of BimgService
func NewBimgService() ImageService {
	return &BimgService{}
}

func (s *BimgService) ConvertToJPEG(files []*multipart.FileHeader) ([]string, error) {
	var convertedPaths []string

	// loop through the uploaded files
	for _, file := range files {
		// open image for reading
		src, err := file.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %v", err)
		}
		defer src.Close()
		// read file and conert to bytes
		fileBytes, err := io.ReadAll(src)
		if err != nil {
			return nil, err
		}
		// convert the bytes to webp file formart
		webp_image, err := bimg.NewImage(fileBytes).Convert(bimg.WEBP)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		output := filepath.Join("output", file.Filename+"_converted.webp")
		if err := os.WriteFile(output, webp_image, 0o644); err != nil {
			return nil, fmt.Errorf("failed to save image: %v", err)
		}

		convertedPaths = append(convertedPaths, output)
	}

	return convertedPaths, nil
}
