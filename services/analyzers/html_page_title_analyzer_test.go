package analyzers

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sura2k/go-web-analyzer/models"
	"golang.org/x/net/html"
)

// Given: An HTML document containing a "title" element inside the "head" element
// When: The PageTitleAnalyzer's Analyze method is invoked on this document
// Then: It should correctly detect and extraxt the page title
func TestHtmlPageTitleAnalyzer_Analyze(t *testing.T) {
	htmlContent := "<!doctype> <html><head><meta></meta><title>Title 01</title></head><body></body></html>"
	expected := "Title 01"

	// Build Node tree
	htmlDoc, err := html.Parse(strings.NewReader(htmlContent))
	require.NoError(t, err, "Failed to prepare test input")

	// Prepare Analyzer
	analyzer := PageTitleAnalyzer{}
	arm := &AnalyzerResultManager{}
	analyzerInput := &models.AnalyzerInput{
		HtmlDoc: htmlDoc,
	}

	// Invoke
	analyzer.Analyze(analyzerInput, arm)

	// Verify
	actual := arm.GetAnalyzerResult().PageTitle
	assert.Equal(t, expected, actual, "Expected: %v", expected)
}

// Given: An HTML document containing a "head" element but no "title" element
// When: The PageTitleAnalyzer's Analyze method is invoked on this document
// Then: It should correctly detect the unavailability of the "title" element and return empty
func TestHtmlPageTitleAnalyzer_Analyze_When_No_TITLE_Elem(t *testing.T) {
	htmlContent := "<!doctype> <html><head></head><body></body></html>"
	expected := ""

	htmlDoc, err := html.Parse(strings.NewReader(htmlContent))
	require.NoError(t, err, "Failed to prepare test input")

	// Prepare Analyzer
	analyzer := PageTitleAnalyzer{}
	arm := &AnalyzerResultManager{}
	analyzerInput := &models.AnalyzerInput{
		HtmlDoc: htmlDoc,
	}

	// Invoke
	analyzer.Analyze(analyzerInput, arm)

	// Verify
	actual := arm.GetAnalyzerResult().PageTitle
	assert.Equal(t, expected, actual, "Expected: %v", expected)
}

// Given: An HTML document containing no "head" element, so no "title" element
// When: The PageTitleAnalyzer's Analyze method is invoked on this document
// Then: It should correctly detect the unavailability of the "head" element and return empty
func TestHtmlPageTitleAnalyzer_Analyze_When_No_HEAD_Elem(t *testing.T) {
	htmlContent := "<!doctype> <html><body></body></html>"
	expected := ""

	htmlDoc, err := html.Parse(strings.NewReader(htmlContent))
	require.NoError(t, err, "Failed to prepare test input")

	// Prepare Analyzer
	analyzer := PageTitleAnalyzer{}
	arm := &AnalyzerResultManager{}
	analyzerInput := &models.AnalyzerInput{
		HtmlDoc: htmlDoc,
	}

	// Invoke
	analyzer.Analyze(analyzerInput, arm)

	// Verify
	actual := arm.GetAnalyzerResult().PageTitle
	assert.Equal(t, expected, actual, "Expected: %v", expected)
}
