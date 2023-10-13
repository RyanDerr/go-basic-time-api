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
	var now time.Time
	query := r.URL.Query()
	loc := query.Get("tz")
	// If No Query Return Current Datetime of Machine
	if loc == "" {
		now = time.Now().UTC()
	} else if loc != "" {
		// Find TimeZone Based On Provided Location
		timezone, err := time.LoadLocation(loc)
		// If Timezone Retrieval Returns An Error Throw Exception, Else Get Time In Specified Timezone
		if err != nil {
			http.Error(w, "Timezone Does Not Exist", http.StatusNotFound)
			return
		} else {
			now = time.Now().In(timezone).UTC()
		}
	}
	// Return Formatted TimeDTO And Return As JSON Data
	timeResponse := TimeDTO{CurrentTime: now}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timeResponse)
}
