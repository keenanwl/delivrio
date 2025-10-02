#!/bin/sh

# Wait for the Docker daemon to be available
until docker info > /dev/null 2>&1; do
  echo "Waiting for Docker to start..."
  sleep 1
done

echo "Docker is up and running!"
