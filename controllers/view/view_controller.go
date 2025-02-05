package view

import (
	"go-web-analyzer/models"
	"go-web-analyzer/services/analyzers"
	"html/template"
	"log"
	"net/http"
)

// ViewAnalyzeHandler
// Endpoint: /
// Methods: GET, POST
func AnalyzerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getHomePageHandler(w, r)
	} else if r.Method == http.MethodPost {
		doAnalyzeHandler(w, r)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

// Returns the analyser page
func getHomePageHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, analyzers.GetEmptyAnalyze())
}

// Processes the submitted URL and returns the analysis
func doAnalyzeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	targetUrl := r.FormValue("targetUrl")
	log.Println("Analyze request received. targetUrl:", targetUrl)

	data := analyzers.Analyze(targetUrl)
	renderTemplate(w, data)
}

// Reusable function to render the tempate
/*
	Parameters:
	- w: The HTTP response writer.
	- data: A pointer to PageAnalysis struct containing analysis details. A pointer is used to minimize the effort of copying objects
*/
func renderTemplate(w http.ResponseWriter, data *models.PageAnalysis) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Println("Template parsing error:", err)
		http.Error(w, "Internal server error. Please try again later!", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Template execution error:", err)
		http.Error(w, "Internal server error. Please try again later!", http.StatusInternalServerError)
		return
	}
}
