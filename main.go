package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/brunaoliveira/pathfinder/services"
)

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

		result := services.CalculateDegrees(dc, modifier)

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
