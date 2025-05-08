package main

import (
	"log"
	"os"

	"prom-wal-guardian/internal/checker"
	"prom-wal-guardian/internal/config"
	"prom-wal-guardian/internal/truncator"
)

func main() {
	// Load configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create checker
	chk := checker.NewChecker(cfg.WALPath, cfg.MaxSize, cfg.MaxAge)

	// Perform WAL check
	result := chk.Check()
	if result.Error != nil {
		log.Fatalf("WAL check failed: %v", result.Error)
	}

	// If cleanup is needed and not in dry-run mode
	if result.NeedsCleanup && !cfg.DryRun {
		log.Printf("WAL size exceeds threshold, starting cleanup...")
		
		// Create truncator
		tr := truncator.NewTruncator(cfg.WALPath, cfg.DryRun)
		
		// Perform cleanup
		if err := tr.Truncate(cfg.MaxSize); err != nil {
			log.Fatalf("WAL cleanup failed: %v", err)
		}
		
		log.Printf("WAL cleanup completed successfully")
	} else if result.NeedsCleanup {
		log.Printf("WAL size exceeds threshold, but running in dry-run mode - no changes made")
	} else {
		log.Printf("WAL size is within limits, no cleanup needed")
	}

	// Exit with appropriate code
	if result.NeedsCleanup {
		os.Exit(1) // Indicate that cleanup was needed
	}
	os.Exit(0) // Everything is fine
} 