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

func TestHandler_GetMicroservice(t *testing.T) {
	handler := &Handler{
		db: &mocks.MockDAO{},
	}

	// Create a new HTTP request with a sample ID
	req, err := http.NewRequest("GET", "/microservice/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a new router and register the handler
	r := chi.NewRouter()
	r.HandleFunc("/microservice/{id}", handler.GetMicroservice)

	// Serve the request using the router
	r.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	var microservice domain.Microservice
	err = json.Unmarshal(rr.Body.Bytes(), &microservice)
	if err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	// Example assertion for checking the ID
	expectedID := "123"
	if microservice.ID != expectedID {
		t.Errorf("handler returned unexpected ID: got %v want %v", microservice.ID, expectedID)
	}

}
