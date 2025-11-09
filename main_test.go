package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorldHandler(t *testing.T) {
	tests := []struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedBody   Response
	}{
		{
			name:           "Valid name starting with A",
			queryParam:     "name=Alice",
			expectedStatus: http.StatusOK,
			expectedBody:   Response{Message: "Hello Alice"},
		},
		{
			name:           "Valid name starting with M",
			queryParam:     "name=Mary",
			expectedStatus: http.StatusOK,
			expectedBody:   Response{Message: "Hello Mary"},
		},
		{
			name:           "Valid name starting with lowercase a",
			queryParam:     "name=alice",
			expectedStatus: http.StatusOK,
			expectedBody:   Response{Message: "Hello alice"},
		},
		{
			name:           "Valid name starting with lowercase m",
			queryParam:     "name=mike",
			expectedStatus: http.StatusOK,
			expectedBody:   Response{Message: "Hello mike"},
		},
		{
			name:           "Invalid name starting with N",
			queryParam:     "name=Nancy",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   Response{Error: "Invalid Input"},
		},
		{
			name:           "Invalid name starting with Z",
			queryParam:     "name=Zane",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   Response{Error: "Invalid Input"},
		},
		{
			name:           "Invalid name starting with lowercase n",
			queryParam:     "name=nancy",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   Response{Error: "Invalid Input"},
		},
		{
			name:           "Invalid name starting with lowercase z",
			queryParam:     "name=zoe",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   Response{Error: "Invalid Input"},
		},
		{
			name:           "Empty name parameter",
			queryParam:     "name=",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   Response{Error: "Invalid Input"},
		},
		{
			name:           "Missing name parameter",
			queryParam:     "",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   Response{Error: "Invalid Input"},
		},
		{
			name:           "Name with only spaces",
			queryParam:     "name=   ",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   Response{Error: "Invalid Input"},
		},
		{
			name:           "Name starting with number",
			queryParam:     "name=1John",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   Response{Error: "Invalid Input"},
		},
		{
			name:           "Name starting with special character",
			queryParam:     "name=@Alice",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   Response{Error: "Invalid Input"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/hello-world?"+tt.queryParam, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(helloWorldHandler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			var response Response
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatal(err)
			}

			if response != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v", response, tt.expectedBody)
			}

			contentType := rr.Header().Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("handler returned wrong content type: got %v want %v", contentType, "application/json")
			}
		})
	}
}