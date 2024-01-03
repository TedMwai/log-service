package handler

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("pong"))
	if err != nil {
		log.Error().Err(err).Msg("Error writing response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
