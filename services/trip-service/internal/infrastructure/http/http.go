package http

import (
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/service"
	"ride-sharing/shared/types"
)

type httpHandler struct {
	service service.Service
}

func NewHttpHandler(service service.Service) httpHandler {
	return httpHandler{service}
}

type requestTripPreview struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func (h *httpHandler) HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody requestTripPreview
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Printf("failed to decode request body:  %v", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	fare := domain.RideFareModel{
		UserID: "USR#43",
	}

	t, err := h.service.CreateTrip(r.Context(), fare)
	if err != nil {
		log.Panicln(err)
	}

	writeJSON(w, http.StatusOK, t)
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
