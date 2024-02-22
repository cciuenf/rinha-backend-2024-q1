# Define variables
TAG=ghcr.io/cciuenf/rinha-backend-2024-q1

# Default target
all: build

# Build the Docker image
build:
	docker build -t $(TAG) .

# Run the application
run:
	docker compose down && docker build -t $(TAG) . && docker compose up

# Publish the Docker image
publish:
	docker build -t $(TAG) . && docker push $(TAG)

# Use .PHONY to specify that these are not file names
.PHONY: build run publish
