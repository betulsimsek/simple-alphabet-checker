# Simple Alphabet Checker API

A simple Go HTTP API that checks if names start with letters in the first half of the alphabet (A-M).

## API Endpoint

**GET** `/hello-world?name={name}`

### Behavior

- **First half of alphabet (A-M, a-m)**: Returns `200 OK` with `{"message": "Hello {name}"}`
- **Second half of alphabet (N-Z, n-z)**: Returns `400 Bad Request` with `{"error": "Invalid Input"}`
- **Missing/empty name**: Returns `400 Bad Request` with `{"error": "Invalid Input"}`
- **Non-alphabetic first character**: Returns `400 Bad Request` with `{"error": "Invalid Input"}`

### Examples

```bash
# Valid request (first half of alphabet)
curl "http://localhost:8080/hello-world?name=Alice"
# Response: 200 OK - {"message":"Hello Alice"}

# Invalid request (second half of alphabet)
curl "http://localhost:8080/hello-world?name=Zane"
# Response: 400 Bad Request - {"error":"Invalid Input"}

# Invalid request (missing name)
curl "http://localhost:8080/hello-world"
# Response: 400 Bad Request - {"error":"Invalid Input"}
```

### Production API Testing

**🌟 Live Production API:** `https://2en31mhkrj.execute-api.eu-west-1.amazonaws.com/prod/hello-world`

```bash
# Test production API
curl "https://2en31mhkrj.execute-api.eu-west-1.amazonaws.com/prod/hello-world?name=Alice"
# Response: {"message":"Hello Alice"}
```

## How to Run

### Prerequisites
- Go 1.18 or later

### Running the Application

1. Clone the repository and navigate to the project directory:
```bash
cd simple-alphabet-checker
```

2. Initialize Go modules (if not already done):
```bash
go mod tidy
```

3. Start the server:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

### Running Tests

Run all unit tests:
```bash
go test
```

Run tests with verbose output:
```bash
go test -v
```

Run tests with coverage:
```bash
go test -cover
```

## Assumptions

1. **Case insensitive**: Both uppercase and lowercase letters are handled (Alice and alice both work)
2. **First character only**: Only the first character of the name is checked
3. **Alphabetic characters**: Names starting with numbers or special characters return "Invalid Input"
4. **Whitespace handling**: Leading and trailing whitespace is trimmed from the name parameter
5. **JSON responses**: All responses are in JSON format with appropriate Content-Type headers
6. **Port**: The server runs on port 8080 by default

## AWS Lambda Deployment

This API is also available as a serverless AWS Lambda function! 

📋 **See [AWS-DEPLOYMENT.md](AWS-DEPLOYMENT.md) for complete deployment guide.**

### Quick Deploy
```bash
# 1. Setup AWS credentials
aws configure

# 2. Create IAM role
./scripts/create-iam-role.sh

# 3. Deploy to Lambda
./build.sh && aws lambda create-function --function-name simple-alphabet-checker --runtime provided.al2023 --role arn:aws:iam::YOUR_ACCOUNT_ID:role/lambda-execution-role --handler bootstrap --zip-file fileb://lambda-deployment.zip --region YOUR_REGION
```

## API Testing with Postman

📋 **Import Collection:** `Simple-Alphabet-Checker-API.postman_collection.json`

### Collection Features:
- ✅ **13 comprehensive test cases** with automated assertions
- ✅ **Environment variables** for easy switching between local/production
- ✅ **Automated tests** for status codes and response validation
- ✅ **Edge case testing** (empty params, special characters, etc.)
- ✅ **HTTP method validation**

### How to Use:
1. **Import** the collection file into Postman
2. **Set environment variable** `baseUrl` to:
   - Local: `http://localhost:8080` 
   - Production: `https://2en31mhkrj.execute-api.eu-west-1.amazonaws.com/prod`
3. **Run collection** to execute all tests automatically

## Project Structure

```
.
├── main.go                                      # Main HTTP server application
├── main_test.go                                # Comprehensive unit tests
├── cmd/lambda/main.go                          # AWS Lambda version
├── go.mod                                      # Go module definition
├── build.sh                                   # Lambda build script
├── deploy.sh                                  # SAM deployment script
├── deploy-simple.sh                           # Simple AWS CLI deployment
├── scripts/                                   # Utility scripts
│   └── create-iam-role.sh
├── .env                                       # Environment configuration
├── .env.example                              # Environment template
├── template.yaml                             # SAM template
├── test-event.json                           # Lambda test payload
├── Simple-Alphabet-Checker-API.postman_collection.json  # Postman test collection
├── AWS-DEPLOYMENT.md                         # Detailed deployment guide
└── README.md                                # This file
```

## Environment Configuration

Copy `.env.example` to `.env` and configure:

```bash
# AWS Configuration
AWS_REGION=your-preferred-region
FUNCTION_NAME=simple-alphabet-checker
STACK_NAME=simple-alphabet-checker-stack

# Lambda Configuration  
LAMBDA_TIMEOUT=10
LAMBDA_MEMORY=128
LAMBDA_RUNTIME=provided.al2023
```

**Security Note:** Never put AWS credentials in `.env` file. Use `aws configure` instead.

## Development vs Production

### Local Development
```bash
go run main.go  # Runs on localhost:8080
```

### AWS Lambda Production
```bash
./deploy.sh     # Deploys to AWS Lambda + API Gateway
```

Both versions share the same business logic but have different entry points for different environments.