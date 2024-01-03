package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func (h *Handler) GetMicroservice(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	ctx := r.Context()

	microservice, err := h.db.GetMicroservice(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("Error getting microservice")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if microservice == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(microservice)
    if err != nil {
        log.Error().Err(err).Msg("Error encoding response body")
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
