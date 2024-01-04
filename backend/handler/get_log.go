package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	lg "github.com/rs/zerolog/log"
)

func (h *Handler) GetLog(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()

	log, err := h.db.GetLog(ctx, id)
	if err != nil {
		lg.Error().Err(err).Msg("Error getting log")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if log == nil {
		lg.Error().Err(err).Msg("Log not found")
		http.Error(w, "Log not found", http.StatusNotFound)
		return
	}

	microservice, err := h.db.GetMicroservice(ctx, log.MicroserviceID)
	if err != nil {
		lg.Error().Err(err).Msg("Error getting microservice associated with the log")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logResp := LogResponse{
		ID:               log.ID,
		MicroserviceName: microservice.Name,
		Level:            log.LogLevel,
		Message:          log.Message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(logResp)
	if err != nil {
		lg.Error().Err(err).Msg("Error encoding response body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}