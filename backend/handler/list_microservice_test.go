package handler

import (
	"encoding/json"
	"log-management/domain"
	"log-management/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestHandler_ListMicroservices(t *testing.T) {
	handler := &Handler{
		db: &mocks.MockDAO{},
	}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/microservices", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a new router and register the handler
	r := chi.NewRouter()
	r.HandleFunc("/microservices", handler.ListMicroservices)

	// Serve the request using the router
	r.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	var microservices []domain.Microservice
	err = json.Unmarshal(rr.Body.Bytes(), &microservices)
	if err != nil {
		t.Errorf("handler returned invalid JSON response: %v", err)
	}

	// Example assertion for checking the number of microservices
	expectedNumMicroservices := 2
	if len(microservices) != expectedNumMicroservices {
		t.Errorf("handler returned unexpected number of microservices: got %v want %v", len(microservices), expectedNumMicroservices)
	}
}
