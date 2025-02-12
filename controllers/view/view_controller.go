package view

import (
	"html/template"
	"log"
	"net/http"

	"github.com/sura2k/go-web-analyzer/models"
	"github.com/sura2k/go-web-analyzer/services/analyzers"
)

// AnalyzerViewHandler
// Endpoint: /
// Methods: GET, POST
func AnalyzerViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getAnalyzerViewHandler(w)
	} else if r.Method == http.MethodPost {
		postAnalyzerViewHandler(w, r)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

// Return Analyzer View
func getAnalyzerViewHandler(w http.ResponseWriter) {
	renderTemplate(w, &models.AnalyzerResponse{})
}

// Return Analyzer View and Model
func postAnalyzerViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Unsupported request method", http.StatusMethodNotAllowed)
		return
	}
	targetUrl := r.FormValue("targetUrl")
	log.Println("AnalyzerViewHandler: Analyze request received. targetUrl:", targetUrl)

	// Invoke AnalyzerManager
	analyzerManager := analyzers.NewAnalyzerManager(targetUrl)
	analyzerResult, err := analyzerManager.Start()

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
// Parameters:
//   - w: The HTTP response writer.
//   - analyzeResponse: A pointer to AnalyzeResponse struct containing analysis details
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
