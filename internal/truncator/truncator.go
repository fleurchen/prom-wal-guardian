package truncator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// Truncator handles WAL cleanup operations
type Truncator struct {
	walPath string
	dryRun  bool
}

// NewTruncator creates a new Truncator instance
func NewTruncator(walPath string, dryRun bool) *Truncator {
	return &Truncator{
		walPath: walPath,
		dryRun:  dryRun,
	}
}

// Truncate removes old WAL segments until the total size is below maxSize
func (t *Truncator) Truncate(maxSize int64) error {
	// Get all WAL files
	files, err := t.getWALFiles()
	if err != nil {
		return err
	}

	// Sort files by modification time (oldest first)
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().Before(files[j].ModTime())
	})

	// Calculate total size
	var totalSize int64
	for _, file := range files {
		totalSize += file.Size()
	}

	// Remove oldest files until we're under maxSize
	for _, file := range files {
		if totalSize <= maxSize {
			break
		}

		filePath := filepath.Join(t.walPath, file.Name())
		log.Printf("Removing old WAL segment: %s (size: %d bytes, age: %v)",
			file.Name(), file.Size(), time.Since(file.ModTime()).Round(time.Minute))

		if !t.dryRun {
			if err := os.Remove(filePath); err != nil {
				return fmt.Errorf("failed to remove %s: %v", filePath, err)
			}
		}

		totalSize -= file.Size()
	}

	return nil
}

// getWALFiles returns all WAL files in the directory
func (t *Truncator) getWALFiles() ([]os.FileInfo, error) {
	var files []os.FileInfo

	err := filepath.Walk(t.walPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, info)
		}
		return nil
	})

	return files, err
} 