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

// Given: An HTML documents containing both accessible and inaccessible internal and external links
// When: The LinksAnalyzer's Analyze method is invoked on this document
// Then: It should correctly detect all the links whether internal or external and detect whether they are accessible, and return the counts
func TestHtmlLinksAnalyzer_Analyze(t *testing.T) {

	// External mock HTTP server
	mockServerExternal := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/release" || r.URL.Path == "/tags" {
			w.WriteHeader(http.StatusOK) // HTTP 200
		} else if r.URL.Path == "/news" {
			w.WriteHeader(http.StatusMultipleChoices) // HTTP 300
		} else if r.URL.Path == "/events" {
			w.WriteHeader(http.StatusNotFound) // HTTP 404
		} else {
			w.WriteHeader(http.StatusInternalServerError) // HTTP 500
		}
	}))
	defer mockServerExternal.Close()

	// Target/Intermal mock HTTP server
	mockServerInternal := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/home" || r.URL.Path == "/aboutus" {
			w.WriteHeader(http.StatusOK) // HTTP 200
		} else if r.URL.Path == "/contactus" {
			w.WriteHeader(http.StatusNotFound) // HTTP 404
		} else {
			w.WriteHeader(http.StatusInternalServerError) // HTTP 500
		}
	}))
	defer mockServerInternal.Close()

	// Input Document -> Assuming as it is from mockServerInternal
	htmlContent := `<!doctype>
		<html>
			<head></head>
			<body>
				<div class="">
					<a href=""></a> 											<!-- Empty link -->
					<a href="#"></a> 											<!-- Valid, but not a link -->
					<span></span>

					<a href="` + mockServerInternal.URL + `/home"></a> 			<!-- Internal -->
					<a href="/aboutus"></a> 									<!-- Internal -->
					<a href="` + mockServerInternal.URL + `/aboutus"></a>		<!-- Internal --> <!-- duplicate -->
					<a href="/contactus"></a> 									<!-- Internal --> <!-- Inaccessible -->
				</div>
				<a href="` + mockServerExternal.URL + `/release"></a> 			<!-- External -->
				<a href="` + mockServerExternal.URL + `/tags"></a>				<!-- External -->
				<a href="` + mockServerExternal.URL + `/news"></a>				<!-- External -->
				<a href="` + mockServerExternal.URL + `/events"></a> 			<!-- External --> <!-- Inaccessible -->
				<a href="` + mockServerExternal.URL + `/release"></a> 			<!-- External --> <!-- duplicate -->
				<a href="` + mockServerExternal.URL + `/deals"></a> 			<!-- External --> <!-- Inaccessible -->
				
				<a href=""></a> 												<!-- Empty link -->
				<a href="emailto:test@test123.comm"></a>						<!-- NonHyperlink -->
			</body>
		</html>`

	//Expected Result
	expected := models.Links{
		External: models.LinkCount{
			Total:        5,
			Inaccessible: 2,
		},
		Internal: models.LinkCount{
			Total:        3,
			Inaccessible: 1,
		},
		EmptyLinks: models.LinkCount{
			Total:        2,
			Inaccessible: 0,
		},
		NonHyperLinks: models.LinkCount{
			Total:        1,
			Inaccessible: 0,
		},
	}

	htmlDoc, err := html.Parse(strings.NewReader(htmlContent))
	require.NoError(t, err, "Failed to prepare test input")

	// Derive Host
	parsedURL, err := url.Parse(mockServerInternal.URL)
	require.NoError(t, err, "Failed to prepare test input")

	// Prepare Analyzer
	analyzer := LinksAnalyzer{}
	arm := &AnalyzerResultManager{}
	analyzerInput := &models.AnalyzerInput{
		HtmlDoc: htmlDoc,
		BaseUrl: mockServerInternal.URL,
		Host:    parsedURL.Host,
	}

	// Invoke
	analyzer.Analyze(analyzerInput, arm)

	// Verify
	actual := arm.GetAnalyzerResult().Links
	assert.Equal(t, expected, actual, "Expected: %v", expected)
}
