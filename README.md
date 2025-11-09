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

## Project Structure

```
.
├── main.go        # Main application with HTTP server and handler
├── main_test.go   # Comprehensive unit tests
├── go.mod         # Go module definition
└── README.md      # This file
```