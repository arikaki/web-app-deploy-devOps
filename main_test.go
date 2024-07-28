package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response using httptest
	rr := httptest.NewRecorder()
	handler := http.FileServer(http.Dir("/static/public"))

	// Serve the request
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect
	expected := `<!DOCTYPE html>` // Example snippet of what might be expected
	if !contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// Utility function to check if a string contains another string
func contains(str, substr string) bool {
	return len(str) >= len(substr) && str[:len(substr)] == substr
}
