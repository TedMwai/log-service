package handler

import (
	"encoding/json"
	"log-management/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestHandler_GetLog(t *testing.T) {
	handler := &Handler{
		db: &mocks.MockDAO{},
	}

	// Create a new HTTP request with a sample log ID
	req, err := http.NewRequest("GET", "/log/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a new router and register the handler
	r := chi.NewRouter()
	r.HandleFunc("/log/{id}", handler.GetLog)

	// Serve the request using the router
	r.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	var logResp LogResponse
	err = json.Unmarshal(rr.Body.Bytes(), &logResp)
	if err != nil {
		t.Fatal(err)
	}

	// Example assertions for checking the log response
	expectedID := "123"
	expectedMicroserviceName := "Test Microservice"
	expectedLevel := "info"
	expectedMessage := "Sample log message"
	if logResp.ID != expectedID {
		t.Errorf("handler returned unexpected log ID: got %v want %v", logResp.ID, expectedID)
	}

	if logResp.MicroserviceName != expectedMicroserviceName {
		t.Errorf("handler returned unexpected microservice name: got %v want %v", logResp.MicroserviceName, expectedMicroserviceName)
	}

	if logResp.Level != expectedLevel {
		t.Errorf("handler returned unexpected log level: got %v want %v", logResp.Level, expectedLevel)
	}

	if logResp.Message != expectedMessage {
		t.Errorf("handler returned unexpected log message: got %v want %v", logResp.Message, expectedMessage)
	}
}
