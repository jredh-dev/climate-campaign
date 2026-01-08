package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templates *template.Template

func init() {
	var err error
	templates, err = template.ParseGlob(filepath.Join("internal", "templates", "*.html"))
	if err != nil {
		log.Printf("Warning: failed to parse templates: %v", err)
	}
}

// Home renders the landing page
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := map[string]interface{}{
		"Title": "The Ice Cubes Have Melted",
	}

	if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// WhoIsRich renders the economic tiers page
func WhoIsRich(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Who Is Rich?",
	}

	if err := templates.ExecuteTemplate(w, "who-is-rich.html", data); err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// OilReality renders the oil data page
func OilReality(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "The Oil Reality",
	}

	if err := templates.ExecuteTemplate(w, "oil-reality.html", data); err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// TakeAction renders the action page
func TakeAction(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Turn Down The Thermostat",
	}

	if err := templates.ExecuteTemplate(w, "take-action.html", data); err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Health check endpoint
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"healthy"}`))
}
