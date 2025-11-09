#!/bin/bash

# Build script for AWS Lambda deployment
set -e

echo "Building Lambda function..."

# Build for Linux (Lambda runtime environment)
GOOS=linux GOARCH=amd64 go build -o bootstrap ./cmd/lambda

# Create deployment package
zip lambda-deployment.zip bootstrap

echo "Deployment package created: lambda-deployment.zip"
echo "File size: $(du -h lambda-deployment.zip | cut -f1)"

# Clean up
rm bootstrap

echo "Build complete!"