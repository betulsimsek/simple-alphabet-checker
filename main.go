package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"unicode"
)

type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	name := strings.TrimSpace(r.URL.Query().Get("name"))
	
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "Invalid Input"})
		return
	}
	
	firstChar := unicode.ToLower(rune(name[0]))
	
	if firstChar < 'a' || firstChar > 'z' {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "Invalid Input"})
		return
	}
	
	if firstChar >= 'a' && firstChar <= 'm' {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{Message: fmt.Sprintf("Hello %s", name)})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "Invalid Input"})
	}
}

func main() {
	http.HandleFunc("/hello-world", helloWorldHandler)
	
	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}