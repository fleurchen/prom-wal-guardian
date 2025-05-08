package checker

import (
	"fmt"
	"log"
	"time"

	"prom-wal-guardian/internal/utils"
)

// CheckResult holds the result of a WAL check
type CheckResult struct {
	TotalSize    int64
	OldestAge    time.Duration
	NeedsCleanup bool
	Error        error
}

// Checker handles WAL checking operations
type Checker struct {
	walPath string
	maxSize int64
	maxAge  int64
}

// NewChecker creates a new Checker instance
func NewChecker(walPath string, maxSize, maxAge int64) *Checker {
	return &Checker{
		walPath: walPath,
		maxSize: maxSize,
		maxAge:  maxAge,
	}
}

// Check performs a WAL check and returns the result
func (c *Checker) Check() CheckResult {
	result := CheckResult{}

	// Check total size
	size, err := utils.GetDirSize(c.walPath)
	if err != nil {
		result.Error = fmt.Errorf("failed to get WAL size: %v", err)
		return result
	}
	result.TotalSize = size

	// Check oldest file age
	age, err := utils.GetOldestFileAge(c.walPath)
	if err != nil {
		result.Error = fmt.Errorf("failed to get oldest file age: %v", err)
		return result
	}
	result.OldestAge = age

	// Determine if cleanup is needed
	result.NeedsCleanup = size > c.maxSize || (c.maxAge > 0 && age > time.Duration(c.maxAge)*time.Hour)

	log.Printf("WAL check results:")
	log.Printf("  Total size: %s", utils.FormatBytes(size))
	log.Printf("  Oldest file age: %v", age.Round(time.Minute))
	log.Printf("  Max size: %s", utils.FormatBytes(c.maxSize))
	if c.maxAge > 0 {
		log.Printf("  Max age: %v", time.Duration(c.maxAge)*time.Hour)
	}
	log.Printf("  Cleanup needed: %v", result.NeedsCleanup)

	return result
} 