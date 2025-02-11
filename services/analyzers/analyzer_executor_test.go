package analyzers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sura2k/go-web-analyzer/models"
	"golang.org/x/net/html"
)

// Given: analyzerInput model
// When: The AnalyzerExecutor's ExecuteAnalyzers method is invoked on this URL
// Then: It should invoke all the analyzers parallaly and return with analyzerResult once all completed
func TestAnalyzerExecutor_ExecuteAnalyzers(t *testing.T) {

	// External mock HTTP server
	mockServerExternal := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/external" {
			w.WriteHeader(http.StatusOK) // HTTP 200
		} else {
			w.WriteHeader(http.StatusInternalServerError) // HTTP 500
		}
	}))
	defer mockServerExternal.Close()

	// Target/Internal mock HTTP server
	mockServerInternal := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/internal" {
			w.WriteHeader(http.StatusOK) // HTTP 200
		} else {
			w.WriteHeader(http.StatusInternalServerError) // HTTP 500
		}
	}))
	defer mockServerInternal.Close()

	// Input Document -> Assuming as it is from mockServerInternal
	htmlContent := `<!doctype html>
	<html>
		<head><title>Title1</title></head>
		<body>
			<a href="` + mockServerInternal.URL + `/internal"></a>
			<a href="` + mockServerExternal.URL + `/external"></a>
			<h1>H1</h1>
			<h5>H5</h5>
			<form>
				<input type="text"></input>
				<input type="password"></input>
				<input type="submit"></input>
			</form>
		</body>
	</html>`

	htmlDoc, err := html.Parse(strings.NewReader(htmlContent))
	require.NoError(t, err, "Failed to prepare test input")

	// Derive Host
	parsedURL, err := url.Parse(mockServerInternal.URL)
	require.NoError(t, err, "Fail to prepare test input")

	analyzerInput := &models.AnalyzerInput{
		HtmlDoc:   htmlDoc,
		BaseUrl:   mockServerInternal.URL,
		Host:      parsedURL.Host,
		TargetUrl: mockServerInternal.URL + "/internal.html",
	}

	// Invoke Preprocessor
	analyzerExecutor := NewAnalyzerExecutor(analyzerInput)
	analyzerResult, err := analyzerExecutor.ExecuteAnalyzers()

	// Verify
	require.NoError(t, err, "Fail to invoke preprocessor")
	assert.NotEmpty(t, analyzerResult, "analyzerResult should not be empty")
	assert.Equal(t, "Title1", analyzerResult.PageTitle, "Expected PageTitle: %v", "Title1")
	assert.Equal(t, true, analyzerResult.HasLoginForm, "Expected HasLoginForm: %v", true)
}
