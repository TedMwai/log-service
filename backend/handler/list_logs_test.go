package handler

import (
	"encoding/json"
	"log-management/domain"
	"log-management/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
)

func TestHandler_ListLogs(t *testing.T) {
	handler := &Handler{
		db: &mocks.MockDAO{},
	}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/logs", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a new router and register the handler
	r := chi.NewRouter()
	r.HandleFunc("/logs", handler.ListLogs)

	// Serve the request using the router
	r.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	var logs []domain.Log
	err = json.Unmarshal(rr.Body.Bytes(), &logs)
	if err != nil {
		t.Fatal(err)
	}

	// Example assertions for checking the logs
	expectedLogs := []domain.Log{
		{
			Base: domain.Base{
				ID:        "123",
				CreatedBy: "SYSTEM",
				UpdatedBy: "SYSTEM",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			MicroserviceID: "123",
			LogLevel:       "info",
			Message:        "Sample log message",
		},
		{
			Base: domain.Base{
				ID:        "456",
				CreatedBy: "SYSTEM",
				UpdatedBy: "SYSTEM",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			MicroserviceID: "123",
			LogLevel:       "info",
			Message:        "Sample log message",
		},
	}
	if len(logs) != len(expectedLogs) {
		t.Errorf("handler returned unexpected number of logs: got %d want %d", len(logs), len(expectedLogs))
	}

	for i, log := range logs {
		if log.ID != expectedLogs[i].ID {
			t.Errorf("handler returned unexpected log ID at index %d: got %s want %s", i, log.ID, expectedLogs[i].ID)
		}

		if log.Message != expectedLogs[i].Message {
			t.Errorf("handler returned unexpected log message at index %d: got %s want %s", i, log.Message, expectedLogs[i].Message)
		}
	}
}
