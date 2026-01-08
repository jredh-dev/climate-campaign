package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jredh-dev/climate-campaign/internal/handlers"
)

var (
	version   = "dev"
	commit    = "none"
	buildDate = "unknown"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	// Static files
	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/who-is-rich", handlers.WhoIsRich)
	mux.HandleFunc("/the-oil-reality", handlers.OilReality)
	mux.HandleFunc("/turn-down-the-thermostat", handlers.TakeAction)
	mux.HandleFunc("/health", handlers.Health)

	addr := ":" + port
	log.Printf("Climate Campaign server v%s (commit: %s, built: %s)", version, commit, buildDate)
	log.Printf("Starting server on %s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"version":"%s","commit":"%s","buildDate":"%s"}`, version, commit, buildDate)
}
