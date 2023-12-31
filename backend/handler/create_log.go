package handler

import (
	"encoding/json"
	"log-management/config"
	"log-management/domain"
	"net/http"
	"strings"

	"github.com/resend/resend-go/v2"
	lg "github.com/rs/zerolog/log"
)

type logRequest struct {
	MicroserviceID string `json:"microservice_id"`
	Level          string `json:"level"`
	Message        string `json:"message"`
}

type LogResponse struct {
	ID               string `json:"id"`
	MicroserviceName string `json:"microservice_name"`
	Level            string `json:"level"`
	Message          string `json:"message"`
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

	// check if log level is fatal
	if strings.ToLower(log.LogLevel) == "fatal" {
		// send email
		client := resend.NewClient(config.RESEND_API_KEY)

		params := &resend.SendEmailRequest{
			From:    "Acme <onboarding@resend.dev>",
			To:      []string{config.DEV_EMAIL},
			Html:    "<strong>We have received a FATAL LOG</strong>",
			Subject: "FATAL LOG RECEIVED 🚩🚩",
		}

		sent, err := client.Emails.Send(params)
		if err != nil {
			lg.Error().Err(err).Msg("Error sending email")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		lg.Info().Msgf("Email sent: %s", sent.Id)
	}

	logResp := LogResponse{
		ID:               log.ID,
		MicroserviceName: microservice.Name,
		Level:            log.LogLevel,
		Message:          log.Message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(logResp)
	if err != nil {
		lg.Error().Err(err).Msg("Error encoding response body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
