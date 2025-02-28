package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, OPTIONS",
		AllowHeaders: "Content-Type, Accept",
	}))

	app.Get("/api/pathfinder2e/v1/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/api/pathfinder2e/v1/distribution", func(c *fiber.Ctx) error {
		modifier, err := strconv.Atoi(c.Query("modifier"))

		if err != nil {
			return fiber.ErrBadRequest
		}

		dc, err := strconv.Atoi(c.Query("dc"))

		if err != nil {
			return fiber.ErrBadRequest
		}

		result := check(dc, modifier)

		return c.JSON(result)
	})

	http.HandleFunc("/api/pathfinder2e/v1/distribution", corsMiddleware(handleDistribution))

	log.Fatal(app.Listen(":4000"))
}

func handleDistribution(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	// Get the request body
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// Convert the request body into a struct
	var data map[string]any
	err = json.Unmarshal(body, &data)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Println(data)

	// Return a response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Distribution handled successfully"))
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

func check(dc int, modifier int) map[string]int {
	criticalFailures := 0
	failures := 0
	successes := 0
	criticalSuccesses := 0

	for i := 1; i <= 20; i++ {
		if i+modifier <= dc-10 {
			criticalFailures++
		}
		if i+modifier < dc && i+modifier > dc-10 {
			failures++
		}
		if i+modifier >= dc && i+modifier < dc+10 {
			successes++
		}
		if i+modifier >= dc+10 {
			criticalSuccesses++
		}
	}

	return map[string]int{
		"critical_failures":  criticalFailures,
		"failures":           failures,
		"successes":          successes,
		"critical_successes": criticalSuccesses,
	}
}
