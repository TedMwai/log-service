package handler

import (
	"bytes"
	"encoding/json"
	"log-management/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestHandler_UpdateLog(t *testing.T) {
	handler := &Handler{
		db: &mocks.MockDAO{},
	}

	// Create a new HTTP request with a sample update log request
	updateLogReq := UpdateLogStruct{
		ID:      "123",
		Message: "Updated log message",
	}
	reqBody, _ := json.Marshal(updateLogReq)
	req, err := http.NewRequest("PUT", "/log", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a new router and register the handler
	r := chi.NewRouter()
	r.HandleFunc("/log", handler.UpdateLog)

	// Serve the request using the router
	r.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check the response body
	var updatedLog UpdateLogStruct
	err = json.Unmarshal(rr.Body.Bytes(), &updatedLog)
	if err != nil {
		t.Fatal(err)
	}

	// Example assertions for checking the updated log
	expectedID := "123"
	expectedMessage := "Updated log message"
	if updatedLog.ID != expectedID {
		t.Errorf("handler returned unexpected log ID: got %v want %v", updatedLog.ID, expectedID)
	}

	if updatedLog.Message != expectedMessage {
		t.Errorf("handler returned unexpected log message: got %v want %v", updatedLog.Message, expectedMessage)
	}
}
