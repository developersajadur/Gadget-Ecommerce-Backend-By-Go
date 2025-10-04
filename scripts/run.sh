#!/bin/bash
set -e

# Load .env file if exists
if [ -f .env ]; then
  export $(cat .env | xargs)
fi

# echo "🚀 Starting server on port $PORT"
go run ./main.go
