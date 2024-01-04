package handler

import (
	"encoding/json"
	"net/http"

	lg "github.com/rs/zerolog/log"
)

func (h *Handler) ListLogs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	logs, err := h.db.ListAllLogs(ctx)
	if err != nil {
		lg.Error().Err(err).Msg("Error listing logs")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(logs)
	if err != nil {
		lg.Error().Err(err).Msg("Error encoding response body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}