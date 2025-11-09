#!/bin/bash

# Simple deploy script using AWS CLI directly
set -e

# Load environment variables
if [ -f .env ]; then
    export $(cat .env | grep -v '^#' | xargs)
fi

echo "🚀 Building and deploying to AWS Lambda..."

# Build
GOOS=linux GOARCH=amd64 go build -o bootstrap ./cmd/lambda
zip lambda-deployment.zip bootstrap

# Use environment variable or default
FUNCTION_NAME=${FUNCTION_NAME:-"simple-alphabet-checker"}

echo "📦 Uploading to Lambda function: $FUNCTION_NAME"

# Update function code
aws lambda update-function-code \
    --function-name $FUNCTION_NAME \
    --zip-file fileb://lambda-deployment.zip

echo "🎉 Deployment complete!"

# Clean up
rm bootstrap lambda-deployment.zip

echo "📋 Test your function with:"
echo "aws lambda invoke --function-name $FUNCTION_NAME --payload '{\"httpMethod\":\"GET\",\"queryStringParameters\":{\"name\":\"Alice\"}}' response.json"