#!/bin/bash

# Deploy script for AWS Lambda using SAM
set -e

# Load environment variables
if [ -f .env ]; then
    export $(cat .env | grep -v '^#' | xargs)
fi

echo "🚀 Deploying Simple Alphabet Checker to AWS Lambda..."

# Check if AWS CLI is installed
if ! command -v aws &> /dev/null; then
    echo "❌ AWS CLI is not installed. Please install it first:"
    echo "brew install awscli"
    exit 1
fi

# Check if SAM CLI is installed
if ! command -v sam &> /dev/null; then
    echo "❌ SAM CLI is not installed. Please install it first:"
    echo "brew install aws-sam-cli"
    exit 1
fi

# Check AWS credentials
if ! aws sts get-caller-identity &> /dev/null; then
    echo "❌ AWS credentials not configured. Please run:"
    echo "aws configure"
    exit 1
fi

echo "✅ AWS CLI and SAM CLI are ready"

# Build the Lambda function
echo "🔨 Building Lambda function..."
./build.sh

# Deploy using SAM
echo "📦 Deploying to AWS..."
sam deploy --guided --stack-name ${STACK_NAME:-simple-alphabet-checker-stack}

echo "🎉 Deployment complete!"
echo "📋 To test your API, use the endpoint URL shown above with:"
echo "curl 'YOUR_API_ENDPOINT/hello-world?name=Alice'"