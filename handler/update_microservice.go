package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type UpdateMicroserviceStruct struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h *Handler) UpdateMicroservice(w http.ResponseWriter, r *http.Request) {
	var microserviceReq UpdateMicroserviceStruct
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&microserviceReq)
	if err != nil {
		log.Error().Err(err).Msg("Error decoding request body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	microservice, err := h.db.GetMicroservice(ctx, microserviceReq.ID)
	if err != nil {
		log.Error().Err(err).Msg("Error getting microservice")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	microservice.Name = microserviceReq.Name
	microservice.Description = microserviceReq.Description

	microservice, err = h.db.UpdateMicroservice(ctx, microservice)
	if err != nil {
		log.Error().Err(err).Msg("Error updating microservice")
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
