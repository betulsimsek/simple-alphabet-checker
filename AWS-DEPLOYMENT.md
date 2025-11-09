# AWS Lambda Deployment Guide

This guide contains the actual deployment steps for the Simple Alphabet Checker API to AWS Lambda.

## ✅ SUCCESSFUL DEPLOYMENT SUMMARY

**🌟 Live API:** `https://YOUR_API_ID.execute-api.YOUR_REGION.amazonaws.com/prod/hello-world`

**📋 Test Commands:**
```bash
# Valid name test (A-M)
curl "https://YOUR_API_ID.execute-api.YOUR_REGION.amazonaws.com/prod/hello-world?name=Alice"
# Response: {"message":"Hello Alice"}

# Invalid name test (N-Z) 
curl "https://YOUR_API_ID.execute-api.YOUR_REGION.amazonaws.com/prod/hello-world?name=Zane"
# Response: {"error":"Invalid Input"}

# Missing parameter test
curl "https://YOUR_API_ID.execute-api.YOUR_REGION.amazonaws.com/prod/hello-world"
# Response: {"error":"Invalid Input"}
```

## Prerequisites

### 1. AWS CLI Installation
```bash
# macOS
brew install awscli

# Linux
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
```

### 2. AWS Credentials Setup
```bash
aws configure
```
Required information:
- AWS Access Key ID
- AWS Secret Access Key  
- Default region: `eu-west-1` (or your preferred region)
- Output format: `json`

### 3. Environment File
```bash
# Edit .env file
AWS_REGION=your-preferred-region
FUNCTION_NAME=simple-alphabet-checker
STACK_NAME=simple-alphabet-checker-stack
```

## ACTUAL DEPLOYMENT STEPS (Executed)

### 1. Environment Preparation
```bash
# Check credentials
aws sts get-caller-identity
# Account ID: YOUR_ACCOUNT_ID
```

### 2. IAM Role Creation
```bash
# Create IAM execution role
./scripts/create-iam-role.sh
# Role ARN: arn:aws:iam::YOUR_ACCOUNT_ID:role/lambda-execution-role
```

### 3. Lambda Function Build & Create
```bash
# Build
./build.sh
# Deployment package: lambda-deployment.zip (~6MB)

# Create Lambda function
aws lambda create-function \
    --function-name simple-alphabet-checker \
    --runtime provided.al2023 \
    --role arn:aws:iam::YOUR_ACCOUNT_ID:role/lambda-execution-role \
    --handler bootstrap \
    --zip-file fileb://lambda-deployment.zip \
    --region YOUR_REGION
```

### 4. API Gateway Setup (Manual)
```bash
# 1. Create REST API
aws apigateway create-rest-api \
    --name "simple-alphabet-checker-api" \
    --region YOUR_REGION
# API ID: YOUR_API_ID

# 2. Create Resource
aws apigateway create-resource \
    --rest-api-id YOUR_API_ID \
    --parent-id YOUR_ROOT_RESOURCE_ID \
    --path-part "hello-world" \
    --region YOUR_REGION
# Resource ID: YOUR_RESOURCE_ID

# 3. Add GET method
aws apigateway put-method \
    --rest-api-id YOUR_API_ID \
    --resource-id YOUR_RESOURCE_ID \
    --http-method GET \
    --authorization-type "NONE" \
    --region YOUR_REGION

# 4. Lambda integration
aws apigateway put-integration \
    --rest-api-id YOUR_API_ID \
    --resource-id YOUR_RESOURCE_ID \
    --http-method GET \
    --type AWS_PROXY \
    --integration-http-method POST \
    --uri "arn:aws:apigateway:YOUR_REGION:lambda:path/2015-03-31/functions/arn:aws:lambda:YOUR_REGION:YOUR_ACCOUNT_ID:function:simple-alphabet-checker/invocations" \
    --region YOUR_REGION

# 5. Grant Lambda permission
aws lambda add-permission \
    --function-name simple-alphabet-checker \
    --statement-id apigateway-invoke \
    --action lambda:InvokeFunction \
    --principal apigateway.amazonaws.com \
    --source-arn "arn:aws:execute-api:YOUR_REGION:YOUR_ACCOUNT_ID:YOUR_API_ID/*/GET/hello-world" \
    --region YOUR_REGION

# 6. Deploy API
aws apigateway create-deployment \
    --rest-api-id YOUR_API_ID \
    --stage-name prod \
    --region YOUR_REGION
```

## Alternative: Automated Deploy (Future Use)

For one-command deployment:

### Option 1: Script Deploy
```bash
./deploy-simple.sh  # For function updates
```

### Option 2: SAM Deploy
```bash
./deploy.sh  # Full infrastructure
```

## ✨ REAL DEPLOYMENT DIFFERENCES

This deployment differs from standard examples:

### 1. **Manual API Gateway Setup**
- Manual setup with AWS CLI instead of SAM/CloudFormation
- Each step executed individually
- Real-time feedback and debugging capability

### 2. **Specific Resource Configuration**
```bash
# Real IDs used (not templates):
API_ID="YOUR_API_ID"
ROOT_RESOURCE_ID="YOUR_ROOT_RESOURCE_ID"  
RESOURCE_ID="YOUR_RESOURCE_ID"
ACCOUNT_ID="YOUR_ACCOUNT_ID"
```

### 3. **Region Choice**
- `eu-west-1` selected (Frankfurt) instead of `us-east-1`
- Lower latency for European users

### 4. **Environment-Driven Configuration**
- `.env` file for configuration management
- Production-ready environment variables

### 5. **Security-First Approach**
- No credentials in .env file (security)
- IAM role created with separate script
- Principle of least privilege

## Test Results (Real)

### API Test Results ✅
```bash
# ✅ Alice (A-M range)
curl "https://YOUR_API_ID.execute-api.YOUR_REGION.amazonaws.com/prod/hello-world?name=Alice"
→ {"message":"Hello Alice"}

# ✅ Zane (N-Z range) 
curl "https://YOUR_API_ID.execute-api.YOUR_REGION.amazonaws.com/prod/hello-world?name=Zane"
→ {"error":"Invalid Input"}

# ✅ Missing parameter
curl "https://YOUR_API_ID.execute-api.YOUR_REGION.amazonaws.com/prod/hello-world"
→ {"error":"Invalid Input"}
```

### Lambda Function Direct Test
```bash
aws lambda invoke \
    --function-name simple-alphabet-checker \
    --payload '{"httpMethod":"GET","queryStringParameters":{"name":"Alice"}}' \
    --region YOUR_REGION \
    response.json
```

## Production Resources (Active)

### AWS Resources
| Resource Type | Name/ID | ARN/URL |
|---------------|---------|---------|
| **Lambda Function** | `simple-alphabet-checker` | `arn:aws:lambda:YOUR_REGION:YOUR_ACCOUNT_ID:function:simple-alphabet-checker` |
| **API Gateway** | `simple-alphabet-checker-api` | `https://YOUR_API_ID.execute-api.YOUR_REGION.amazonaws.com` |
| **IAM Role** | `lambda-execution-role` | `arn:aws:iam::YOUR_ACCOUNT_ID:role/lambda-execution-role` |
| **Log Group** | `/aws/lambda/simple-alphabet-checker` | CloudWatch Logs |

### Function Configuration
```json
{
  "FunctionName": "simple-alphabet-checker",
  "Runtime": "provided.al2023", 
  "MemorySize": 128,
  "Timeout": 3,
  "Architecture": "x86_64",
  "CodeSize": "~6MB"
}
```

## Updates (Function Code)

```bash
# After code changes
./build.sh

# Function update
aws lambda update-function-code \
    --function-name simple-alphabet-checker \
    --zip-file fileb://lambda-deployment.zip \
    --region YOUR_REGION
```

## Cleanup (Deletion)

```bash
# 1. Delete API Gateway
aws apigateway delete-rest-api --rest-api-id YOUR_API_ID --region YOUR_REGION

# 2. Delete Lambda function  
aws lambda delete-function --function-name simple-alphabet-checker --region YOUR_REGION

# 3. Delete IAM role
aws iam detach-role-policy --role-name lambda-execution-role --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
aws iam delete-role --role-name lambda-execution-role
```

## Cost Analysis (Real)

**Lambda:**
- Request count: ~10/day test → Within free tier
- Memory: 128MB × 3s timeout → ~$0.00001 per invocation
- **Monthly estimate:** $0.01

**API Gateway:** 
- Request count: ~10/day → $0.0001
- **Monthly estimate:** $0.003

**TOTAL: ~$0.013/month** (very cost-effective)

## Monitoring & Logs

```bash
# CloudWatch logs
aws logs describe-log-groups --region YOUR_REGION | grep simple-alphabet

# Function metrics
aws cloudwatch get-metric-statistics \
    --namespace AWS/Lambda \
    --metric-name Invocations \
    --dimensions Name=FunctionName,Value=simple-alphabet-checker \
    --start-time 2025-11-09T18:00:00Z \
    --end-time 2025-11-09T22:00:00Z \
    --period 3600 \
    --statistics Sum \
    --region YOUR_REGION
```