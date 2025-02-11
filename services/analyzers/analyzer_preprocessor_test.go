package analyzers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

// Given: A URL
// When: The AnalyzerPreprocessor's ExecutePreprocessor method is invoked on this URL
// Then: It should correctly check whether URL is accessible, fetch the HTML content, derive base URL and Host, and return the data
func TestAnalyzerPreprocessor_ExecutePreprocessor(t *testing.T) {

	//Expected Page Title
	expectedPageTitle := "Test Title"

	// Mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/release.html" {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.01 Transitional//EN\"><html><head><title>"+expectedPageTitle+"</title></head><body><span>H1</span></body></html>")
		} else {
			w.WriteHeader(http.StatusInternalServerError) // HTTP 500
		}
	}))
	defer mockServer.Close()

	// Derive Host
	parsedURL, err := url.Parse(mockServer.URL)
	require.NoError(t, err, "Fail to prepare test input")

	// Expected Base URL
	expectedBaseUrl := mockServer.URL
	expectedHost := parsedURL.Host

	//Target URL (URL sent by the user)
	targetUrl := mockServer.URL + "/release.html"

	// Invoke Preprocessor
	analyzerPreprocessor := NewAnalyzerPreprocessor(targetUrl)
	analyzerInput, err := analyzerPreprocessor.ExecutePreprocessor()

	// Verify
	require.NoError(t, err, "Fail to invoke preprocessor")
	assert.NotEmpty(t, analyzerInput, "analyzerInput should not be empty")
	assert.Equal(t, targetUrl, analyzerInput.TargetUrl, "Expected: %v", targetUrl)
	assert.Equal(t, expectedBaseUrl, analyzerInput.BaseUrl, "Expected: %v", expectedBaseUrl)
	assert.Equal(t, expectedHost, analyzerInput.Host, "Expected: %v", expectedHost)
	assert.NotEmpty(t, analyzerInput.HtmlDoc, "analyzerInput.HtmlDoc should not be empty")
	assert.IsType(t, &html.Node{}, analyzerInput.HtmlDoc, "analyzerInput.HtmlDoc should be type of *html.Node")
	assert.Equal(t, expectedPageTitle, analyzerInput.HtmlDoc.FirstChild.NextSibling.FirstChild.FirstChild.FirstChild.Data, "Expected Title: %v", expectedPageTitle)
}

// Given: A URL which is not accessible
// When: The AnalyzerPreprocessor's ExecutePreprocessor method is invoked on this URL
// Then: It should correctly check whether URL is accessible, and if not, return with an error
func TestAnalyzerPreprocessor_ExecutePreprocessor_URL404(t *testing.T) {
	// Mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/release.html" {
			w.WriteHeader(http.StatusNotFound) // HTTP 404
		} else {
			w.WriteHeader(http.StatusInternalServerError) // HTTP 500
		}
	}))
	defer mockServer.Close()

	//Target URL (URL sent by the user)
	targetUrl := mockServer.URL + "/release.html"

	// Invoke Preprocessor
	analyzerPreprocessor := NewAnalyzerPreprocessor(targetUrl)
	_, err := analyzerPreprocessor.ExecutePreprocessor()

	// Verify
	assert.NotEmpty(t, err, "err should not be empty")
}
