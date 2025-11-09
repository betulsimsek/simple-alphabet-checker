#!/bin/bash

# Create IAM role for Lambda function
set -e

# Load environment variables
if [ -f .env ]; then
    export $(cat .env | grep -v '^#' | xargs)
fi

ROLE_NAME="lambda-execution-role"

echo "🔐 Creating IAM role for Lambda..."

# Create trust policy
cat > trust-policy.json << EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF

# Create role
aws iam create-role \
    --role-name $ROLE_NAME \
    --assume-role-policy-document file://trust-policy.json \
    --description "Execution role for simple-alphabet-checker Lambda function" || true

# Attach basic execution policy
aws iam attach-role-policy \
    --role-name $ROLE_NAME \
    --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole

# Get account ID
ACCOUNT_ID=$(aws sts get-caller-identity --query Account --output text)

echo "✅ IAM role created successfully!"
echo "📋 Role ARN: arn:aws:iam::$ACCOUNT_ID:role/$ROLE_NAME"

# Clean up
rm trust-policy.json