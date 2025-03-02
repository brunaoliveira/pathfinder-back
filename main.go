package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Result struct {
	CriticalFailures  int `json:"critical_failures"`
	Failures          int `json:"failures"`
	Successes         int `json:"successes"`
	CriticalSuccesses int `json:"critical_successes"`
}

func main() {

	fmt.Printf("1")

	http.HandleFunc("/api/pathfinder2e/v1/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/api/pathfinder2e/v1/distribution", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		modifier, err := strconv.Atoi(r.URL.Query().Get("modifier"))

		if err != nil {
			http.Error(w, "Invalid modifier", http.StatusBadRequest)
			return
		}

		dc, err := strconv.Atoi(r.URL.Query().Get("dc"))

		if err != nil {
			http.Error(w, "Invalid dc", http.StatusBadRequest)
			return
		}

		result := calculateDegrees(dc, modifier)

		json.NewEncoder(w).Encode(result)
	}))

	log.Fatal(http.ListenAndServe(":4000", nil))
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func calculateDegrees(dc int, modifier int) map[string]int {
	var result Result

	result.CriticalFailures = max(0, min(20, dc-10-modifier))
	result.Failures = max(0, min(20, dc-modifier-1)) - result.CriticalFailures
	result.Successes = max(0, min(20, dc+10-modifier-1)) - result.Failures - result.CriticalFailures
	result.CriticalSuccesses = 20 - result.CriticalFailures - result.Failures - result.Successes

	result = ajustNaturalOne(modifier, dc, result)
	result = adjustNaturalTwenty(modifier, dc, result)

	return map[string]int{
		"critical_failures":  result.CriticalFailures,
		"failures":           result.Failures,
		"successes":          result.Successes,
		"critical_successes": result.CriticalSuccesses,
	}
}

func ajustNaturalOne(modifier int, dc int, result Result) Result {
	if modifier+1 >= dc+10 { // critical success -> success
		result.CriticalSuccesses--
		result.Successes++
	} else if modifier+1 >= dc { // success -> failure
		result.Successes--
		result.Failures++
	} else if modifier+1 >= dc-10 { // failure -> critical failure
		result.Failures--
		result.CriticalFailures++
	}

	return result
}

func adjustNaturalTwenty(modifier int, dc int, result Result) Result {

	if modifier+20 <= dc-10 { // critical failures -> failures
		result.CriticalFailures--
		result.Failures++
	} else if modifier+20 < dc { // failures -> successes
		result.Failures--
		result.Successes++
	} else if modifier+20 < dc+10 { // successes -> critical successes
		result.Successes--
		result.CriticalSuccesses++
	}

	return result
}
