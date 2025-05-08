#!/bin/bash

# Default values
WAL_PATH="/prometheus/data/wal"
MAX_SIZE="5GB"
DRY_RUN="false"

# Parse command line arguments
while [[ $# -gt 0 ]]; do
  case $1 in
    --path)
      WAL_PATH="$2"
      shift 2
      ;;
    --max-size)
      MAX_SIZE="$2"
      shift 2
      ;;
    --dry-run)
      DRY_RUN="true"
      shift
      ;;
    *)
      echo "Unknown option: $1"
      exit 1
      ;;
  esac
done

# Run the guardian
/prom-wal-guardian --path "$WAL_PATH" --max-size "$MAX_SIZE" --dry-run "$DRY_RUN"

# Check exit code
case $? in
  0)
    echo "WAL check passed, no cleanup needed"
    exit 0
    ;;
  1)
    echo "WAL cleanup performed successfully"
    exit 0
    ;;
  *)
    echo "WAL check/cleanup failed"
    exit 1
    ;;
esac 