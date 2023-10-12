package app

import (
	"encoding/json"
	"net/http"
	"time"
)

type TimeDTO struct {
	CurrentTime time.Time `json:"current_time"`
}

func getTime(w http.ResponseWriter, r *http.Request) {
	now := time.Now().UTC()

	timeResponse := TimeDTO{CurrentTime: now}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timeResponse)
}
