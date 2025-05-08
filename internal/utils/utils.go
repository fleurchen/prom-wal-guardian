package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

// GetDirSize returns the total size of a directory in bytes
func GetDirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}

// GetOldestFileAge returns the age of the oldest file in the directory
func GetOldestFileAge(path string) (time.Duration, error) {
	var oldestTime time.Time
	first := true

	err := filepath.Walk(path, func(_ string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if first {
				oldestTime = info.ModTime()
				first = false
			} else if info.ModTime().Before(oldestTime) {
				oldestTime = info.ModTime()
			}
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	if first {
		return 0, fmt.Errorf("no files found in directory")
	}

	return time.Since(oldestTime), nil
}

// FormatBytes converts bytes to human readable format
func FormatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
} 