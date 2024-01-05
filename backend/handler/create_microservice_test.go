package handler

import (
	"bytes"
	"encoding/json"
	"log-management/domain"
	"log-management/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestHandler_CreateMicroservice(t *testing.T) {
	handler := &Handler{
		db: &mocks.MockDAO{},
	}

	// Create a sample request body
	microserviceReq := microserviceRequest{
		Name:        "Test Microservice",
		Description: "Test Description",
	}

	reqBody, err := json.Marshal(microserviceReq)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the sample request body
	req, err := http.NewRequest("POST", "/microservice", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a new router and register the handler
	r := chi.NewRouter()
	r.HandleFunc("/microservice", handler.CreateMicroservice)

	// Serve the request using the router
	r.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check the response body
	var microserviceRes domain.Microservice
	err = json.Unmarshal(rr.Body.Bytes(), &microserviceRes)
	if err != nil {
		t.Fatal(err)
	}

	expected := microserviceReq.Name
	if microserviceRes.Name != expected {
		t.Errorf("handler returned unexpected name: got %v want %v", microserviceRes.Name, expected)
	}

	expected = microserviceReq.Description
	if microserviceRes.Description != expected {
		t.Errorf("handler returned unexpected description: got %v want %v", microserviceRes.Description, expected)
	}
}
