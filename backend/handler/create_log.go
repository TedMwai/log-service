package handler

import (
	"encoding/json"
	"log-management/domain"
	"net/http"

	lg "github.com/rs/zerolog/log"
)

type logRequest struct {
	MicroserviceID string `json:"microservice_id"`
	Level          string `json:"level"`
	Message        string `json:"message"`
}

func (h *Handler) CreateLog(w http.ResponseWriter, r *http.Request) {
	var logReq logRequest
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&logReq)
	if err != nil {
		lg.Error().Err(err).Msg("Error decoding request body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	microservice, err := h.db.GetMicroservice(ctx, logReq.MicroserviceID)
	if err != nil {
		lg.Error().Err(err).Msg("Error getting microservice")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if microservice == nil {
		lg.Error().Err(err).Msg("Microservice not found")
		http.Error(w, "Microservice not found", http.StatusNotFound)
		return
	}

	log := domain.NewLog(logReq.MicroserviceID, logReq.Level, logReq.Message)

	log, err = h.db.CreateLog(ctx, log)
	if err != nil {
		lg.Error().Err(err).Msg("Error creating log")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(log)
	if err != nil {
		lg.Error().Err(err).Msg("Error encoding response body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}