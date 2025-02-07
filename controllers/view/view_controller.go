package view

import (
	"go-web-analyzer/models"
	"go-web-analyzer/services/analyzers"
	"html/template"
	"log"
	"net/http"
)

// AnalyzerViewHandler
// Endpoint: /
// Methods: GET, POST
func AnalyzerViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getAnalyzerViewHandler(w, r)
	} else if r.Method == http.MethodPost {
		postAnalyzerViewHandler(w, r)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

// Return Analyzer View
func getAnalyzerViewHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, &models.AnalyzerResult{})
}

// Return Analyzer View and Model
func postAnalyzerViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Unsupported request method", http.StatusMethodNotAllowed)
		return
	}
	targetUrl := r.FormValue("targetUrl")
	log.Println("Analyze request received. targetUrl:", targetUrl)

	analyzerResult := analyzers.StartAnalyzer(targetUrl)
	renderTemplate(w, analyzerResult)
}

// Reusable function to render the tempate
/*
	Parameters:
	- w: The HTTP response writer.
	- data: A pointer to PageAnalysis struct containing analysis details. A pointer is used to minimize the effort of copying objects
*/
func renderTemplate(w http.ResponseWriter, analyzerResult *models.AnalyzerResult) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Println("Template parsing failed. error:", err)
		http.Error(w, "Internal server error. Please try again later!", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, analyzerResult)
	if err != nil {
		log.Println("Template execution failed. error:", err)
		http.Error(w, "Internal server error. Please try again later!", http.StatusInternalServerError)
		return
	}
}
