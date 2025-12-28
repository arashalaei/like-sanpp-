package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/shared/contracts"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody requestTripPreview
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Printf("failed to decode request body:  %v", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// validation
	if reqBody.UserID == "" {
		http.Error(w, "userID is required", http.StatusBadRequest)
		return
	}
	jsonBody, _ := json.Marshal(reqBody)
	reader := bytes.NewReader(jsonBody)

	// TODO: CALL TRIP SERVICE
	resp, err := http.Post("http://trip-service:8083/perivew", "application/json", reader)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	var respBody any
	if err := json.NewDecoder(r.Body).Decode(&respBody); err != nil {
		http.Error(w, "failed to pars json data from trip service", http.StatusBadRequest)
		return
	}
	response := contracts.APIResponse{
		Data: respBody,
	}
	writeJSON(w, http.StatusAccepted, response)
}
