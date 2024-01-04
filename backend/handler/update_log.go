package handler

import (
	"encoding/json"
	"net/http"

	lg "github.com/rs/zerolog/log"
)

type UpdateLogStruct struct {
	ID          string `json:"id"`
	Message     string `json:"message"`
}

func (h *Handler) UpdateLog(w http.ResponseWriter, r *http.Request) {
	var logReq UpdateLogStruct
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&logReq)
	if err != nil {
		lg.Error().Err(err).Msg("Error decoding request body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log, err := h.db.GetLog(ctx, logReq.ID)
	if err != nil {
		lg.Error().Err(err).Msg("Error getting log")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Message = logReq.Message

	log, err = h.db.UpdateLog(ctx, log)
	if err != nil {
		lg.Error().Err(err).Msg("Error updating log")
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