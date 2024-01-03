package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func (h *Handler) ListMicroservices(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	microservices, err := h.db.ListMicroservices(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error listing microservices")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(microservices)
	if err != nil {
		log.Error().Err(err).Msg("Error encoding response body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}