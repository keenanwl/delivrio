#!/bin/sh

# Start Docker service
service docker start

# Wait for Docker to be ready
/usr/local/bin/wait-for-docker.sh

# Execute the provided command
exec "$@"
