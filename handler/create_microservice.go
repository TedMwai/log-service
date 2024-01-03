package handler

import (
	"encoding/json"
	"log-management/domain"
	"net/http"

	"github.com/rs/zerolog/log"
)

type microserviceRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h *Handler) CreateMicroservice(w http.ResponseWriter, r *http.Request) {
	var microserviceReq microserviceRequest
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&microserviceReq)
	if err != nil {
		log.Error().Err(err).Msg("Error decoding request body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	microservice := domain.NewMicroservice(microserviceReq.Name, microserviceReq.Description)

	microservice, err = h.db.CreateMicroservice(ctx, microservice)
	if err != nil {
		log.Error().Err(err).Msg("Error creating microservice")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)

    err = json.NewEncoder(w).Encode(microservice)
    if err != nil {
        log.Error().Err(err).Msg("Error encoding response body")
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
