package handler

import (
	"encoding/json"
	"net/http"

	lg "github.com/rs/zerolog/log"
)

type getLogRequest struct {
	ID string `json:"id"`
}

func (h *Handler) GetLog(w http.ResponseWriter, r *http.Request) {
	var getLogReq getLogRequest
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&getLogReq)
	if err != nil {
		lg.Error().Err(err).Msg("Error decoding request body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log, err := h.db.GetLog(ctx, getLogReq.ID)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// TODO: Send also microservice name
	err = json.NewEncoder(w).Encode(log)
	if err != nil {
		lg.Error().Err(err).Msg("Error encoding response body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}