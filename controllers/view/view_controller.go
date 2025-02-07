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
	renderTemplate(w, &models.AnalyzerResponse{})
}

// Return Analyzer View and Model
func postAnalyzerViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Unsupported request method", http.StatusMethodNotAllowed)
		return
	}
	targetUrl := r.FormValue("targetUrl")
	log.Println("Analyze request received. targetUrl:", targetUrl)

	analyzerResult, err := analyzers.StartAnalyzer(targetUrl)
	analyzerResponse := &models.AnalyzerResponse{Processed: true}
	if err != nil {
		analyzerResponse.Status = false
		analyzerResponse.Message = err.Error()
	} else {
		analyzerResponse.Status = true
		analyzerResponse.Data = *analyzerResult
	}

	renderTemplate(w, analyzerResponse)
}

// Reusable function to render the tempate
/*
	Parameters:
	- w: The HTTP response writer.
	- analyzeResponse: A pointer to AnalyzeResponse struct containing analysis details. A pointer is used to minimize the effort of copying objects
*/
func renderTemplate(w http.ResponseWriter, analyzerResponse *models.AnalyzerResponse) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Println("Template parsing failed. error:", err)
		http.Error(w, "Internal server error. Please try again later!", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, analyzerResponse)
	if err != nil {
		log.Println("Template execution failed. error:", err)
		http.Error(w, "Internal server error. Please try again later!", http.StatusInternalServerError)
		return
	}
}
