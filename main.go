// Sample 7 of the bring-your-code test matrix: Go + chi. The platform must find: backend, golang, port 8080 (PORT),
// health at /health, no database, OPENWEATHER_KEY as a required secret.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/maria-orkosi/byoc-go-chi/internal/weather"
)

func main() {
	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	r.Get("/forecast/{city}", func(w http.ResponseWriter, req *http.Request) {
		json.NewEncoder(w).Encode(weather.Forecast(chi.URLParam(req, "city"), os.Getenv("OPENWEATHER_KEY")))
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("forecast api on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
