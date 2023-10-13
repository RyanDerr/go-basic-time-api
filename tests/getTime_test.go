package tests

import (
	"github.com/RyanDerr/go-basic-time-api/app"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTime(t *testing.T) {
	// Test case 1: No query parameter, expect current time in UTC.
	req1, err := http.NewRequest("GET", "/time", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder1 := httptest.NewRecorder()
	app.GetTime(recorder1, req1)
	if recorder1.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder1.Code)
	}

	// Test case 2: Valid query parameter for a timezone.
	req2, err := http.NewRequest("GET", "/time?tz=America/New_York", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder2 := httptest.NewRecorder()
	app.GetTime(recorder2, req2)
	if recorder2.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder2.Code)
	}

	// Test case 3: Invalid query parameter for a timezone.
	req3, err := http.NewRequest("GET", "/time?tz=InvalidTimeZone", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder3 := httptest.NewRecorder()
	app.GetTime(recorder3, req3)
	if recorder3.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, recorder3.Code)
	}
}
