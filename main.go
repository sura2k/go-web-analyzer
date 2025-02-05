package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

// PageAnalysis struct for passing data to the template
type PageAnalysis struct {
	TargetUrl        string
	HtmlVersion      string
	PageTitle        string
	Headings         Headings
	Links            Links
	HasLoginForm     bool
	RedirectsToLogin bool
	Status           bool
	StatusMessage    string
}

// Headings struct to store heading count by level
type Headings struct {
	NumOfH1 int
	NumOfH2 int
	NumOfH3 int
	NumOfH4 int
	NumOfH5 int
	NumOfH6 int
}

// Links struct to store internal, external, inaccessible links count
type Links struct {
	NumOfIntLinks             int
	NumOfExtLinks             int
	NumOfIntLinksInaccessible int
	NumOfExtLinksInaccessible int
}

func main() {
	//Handler Mappings
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/analyze", analyzeHandler)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

// HomePageHandler - Returns the analyser page
func homePageHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, &PageAnalysis{})
}

// AnalyzeHandler - Processes the submitted URL and returns the analysis
func analyzeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	targetUrl := r.FormValue("targetUrl")
	fmt.Println("Analyze request received. targetUrl:", targetUrl)

	data := PageAnalysis{
		TargetUrl:     targetUrl,
		Status:        true,
		StatusMessage: "",
	}

	if !isValidUrl(targetUrl) {
		data.Status = false
		data.StatusMessage = "Provided URL is invalid"
		renderTemplate(w, &data)
		return
	}

	data.HtmlVersion = "HTML5"
	data.PageTitle = "Demo"
	data.Headings = Headings{
		NumOfH1: 0,
		NumOfH2: 0,
		NumOfH3: 0,
		NumOfH4: 0,
		NumOfH5: 0,
		NumOfH6: 0,
	}
	data.Links = Links{
		NumOfIntLinks:             0,
		NumOfExtLinks:             0,
		NumOfIntLinksInaccessible: 0,
		NumOfExtLinksInaccessible: 0,
	}
	data.HasLoginForm = true
	data.RedirectsToLogin = false
	data.Status = true
	data.StatusMessage = "TODO"

	renderTemplate(w, &data)
}

// Reusable function to render the tempate
/*
	Parameters:
	- w: The HTTP response writer.
	- data: A pointer to PageAnalysis struct containing analysis details. A pointer is used to minimize the effort of copying objects
*/
func renderTemplate(w http.ResponseWriter, data *PageAnalysis) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		fmt.Println("Template parsing error:", err)
		http.Error(w, "Internal server error. Please try again later!", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("Template execution error:", err)
		http.Error(w, "Internal server error. Please try again later!", http.StatusInternalServerError)
		return
	}
}

// Check whether URL is in valid format
func isValidUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
