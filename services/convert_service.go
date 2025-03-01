package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/h2non/bimg"
)

type ImageService interface {
	InitializeJobs(file []*multipart.FileHeader, format bimg.ImageType) ([]string, error)
}

// BimgService struct implementing ImageService
type BimgService struct{}

// NewBimgService returns an instance of BimgService
func NewBimgService() ImageService {
	return &BimgService{}
}

type Job struct {
	Number int
	File   *multipart.FileHeader
	Format bimg.ImageType
}

func ConvertFilesWoker(jobs <-chan Job, results chan<- string, wg *sync.WaitGroup) error {
	defer wg.Done()

	// loop through the uploaded files
	for file := range jobs {
		// open image for reading
		src, err := file.File.Open()
		if err != nil {
			return fmt.Errorf("failed to open file: %v", err)
		}
		defer src.Close()
		// read file and conert to bytes
		fileBytes, err := io.ReadAll(src)
		if err != nil {
			return err
		}
		// convert the bytes to the specified file formart
		webp_image, err := bimg.NewImage(fileBytes).Convert(file.Format)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		ext := getExtension(file.Format)

		outputpath := filepath.Join("output", file.File.Filename+"_converted."+ext)
		if err := os.WriteFile(outputpath, webp_image, 0o644); err != nil {
			return fmt.Errorf("failed to save image: %v", err)
		}

		results <- outputpath

	}

	return nil
}

func (s *BimgService) InitializeJobs(images []*multipart.FileHeader, format bimg.ImageType) ([]string, error) {
	// number of wokers in respect to your machines CPU Cores.
	tasks := len(images)
	numwokers := runtime.NumCPU()
	// number of jobs to be completed i.e number of images to be converted.
	jobs := make(chan Job, tasks)
	results := make(chan string, tasks)
	var wg sync.WaitGroup

	// start wokers to convert images
	for i := 0; i < numwokers; i++ {
		wg.Add(1)
		go ConvertFilesWoker(jobs, results, &wg)
	}
	// send tasks to jobs
	for i, file := range images {
		jobs <- Job{Number: i, File: file, Format: format}
	}
	close(jobs)
	wg.Wait()

	close(results)

	var convertedPaths []string
	for path := range results {
		convertedPaths = append(convertedPaths, path)
	}

	return convertedPaths, nil
}

// getExtension maps bimg.ImageType to a file extension
func getExtension(format bimg.ImageType) string {
	switch format {
	case bimg.WEBP:
		return "webp"
	case bimg.JPEG:
		return "jpg"
	case bimg.PNG:
		return "png"
	case bimg.TIFF:
		return "tiff"
	case bimg.AVIF:
		return "avif"
	case bimg.SVG:
		return "svg"
	default:
		return "img"
	}
}
