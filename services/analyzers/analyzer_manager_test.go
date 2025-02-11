package analyzers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Given: A URL
// When: The AnalyzerManager's Start method is invoked on this URL
// Then: It should act as a orchestrator and delegate the work to Preprocessor and AnayzerExecutor, then return what executor returns
func TestAnalyzerManager_Start(t *testing.T) {

	// Document for the targetUrl
	htmlContent := `<!doctype html>
	<html>
		<head><title>Title1</title></head>
		<body>
			<a href="/internal"></a>
			<a href="/external"></a>
			<h1>H1</h1>
			<h5>H5</h5>
			<form>
				<input type="text"></input>
				<input type="password"></input>
				<input type="submit"></input>
			</form>
		</body>
	</html>`

	// Target/Internal mock HTTP server
	mockServerInternal := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/page.html" {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, htmlContent)
		} else if r.URL.Path == "/internal" {
			w.WriteHeader(http.StatusOK) // HTTP 200
		} else {
			w.WriteHeader(http.StatusInternalServerError) // HTTP 500
		}
	}))
	defer mockServerInternal.Close()

	// Invoke Preprocessor
	analyzerManager := NewAnalyzerManager(mockServerInternal.URL + "/page.html")
	analyzerResult, err := analyzerManager.Start()

	// Verify
	require.NoError(t, err, "Fail to invoke preprocessor")
	assert.NotEmpty(t, analyzerResult, "analyzerResult should not be empty")
}
