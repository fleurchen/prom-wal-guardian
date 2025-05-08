package tests

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"prom-wal-guardian/internal/checker"
)

func TestChecker(t *testing.T) {
	// Create temporary directory for test
	tempDir, err := os.MkdirTemp("", "wal-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create some test files
	files := []struct {
		name    string
		size    int64
		content string
	}{
		{"000000", 1024, "test data 1"},
		{"000001", 2048, "test data 2"},
		{"000002", 4096, "test data 3"},
	}

	for _, f := range files {
		path := filepath.Join(tempDir, f.name)
		if err := os.WriteFile(path, []byte(f.content), 0644); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	// Create checker with small max size
	chk := checker.NewChecker(tempDir, 5000, 0)

	// Perform check
	result := chk.Check()
	if result.Error != nil {
		t.Fatalf("Check failed: %v", result.Error)
	}

	// Verify results
	if !result.NeedsCleanup {
		t.Errorf("Expected cleanup to be needed, but it wasn't")
	}

	expectedSize := int64(7168) // 1024 + 2048 + 4096
	if result.TotalSize != expectedSize {
		t.Errorf("Expected total size %d, got %d", expectedSize, result.TotalSize)
	}
} 