package analyzers

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sura2k/go-web-analyzer/models"
	"golang.org/x/net/html"
)

// Given: List of HTML documents containing all valid <!DOCTYPE> directives
// When: The HtmlVersionAnalyzer's Analyze method is invoked on this document
// Then: It should correctly detect HTML Version and return the value
func TestHtmlVersionAnalyzer_Analyze_Known_Versions(t *testing.T) {
	samples := []struct {
		htmlContent string
		expected    string
	}{
		{"<html><head></head><body></body></html>", "HTML5"},
		{"<!doctype> <html><head></head><body></body></html>", "HTML5"},
		{"<!DOCTYPE html> <html><head></head><body></body></html>", "HTML5"},
		{"<!doctype HTML PUBLIC \"-//W3C//DTD HTML 4.01//EN\"> <html><head></head><body></body></html>", "HTML 4.01 Strict"},
		{"<!doctype HTML PUBLIC \"-//W3C//DTD HTML 4.01 Transitional//EN\"> <html><head></head><body></body></html>", "HTML 4.01 Transitional"},
		{"<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.01 Frameset//EN\"> <html><head></head><body></body></html>", "HTML 4.01 Frameset"},
		{"<!doctype HTML PUBLIC \"-//W3C//DTD XHTML 1.0 Strict//EN\"> <html><head></head><body></body></html>", "XHTML 1.0 Strict"},
		{"<!DOCTYPE HTML PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\"> <html><head></head><body></body></html>", "XHTML 1.0 Transitional"},
		{"<!doctype html PUBLIC \"-//W3C//DTD XHTML 1.0 Frameset//EN\"> <html><head></head><body></body></html>", "XHTML 1.0 Frameset"},
		{"<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.1//EN\"> <html><head></head><body></body></html>", "XHTML 1.1"},
	}

	// Loop through samples
	for _, sample := range samples {
		htmlDoc, err := html.Parse(strings.NewReader(sample.htmlContent))
		require.NoError(t, err, "Failed to prepare test input")

		// Prepare Analyzer
		analyzer := HtmlVersionAnalyzer{}
		arm := &AnalyzerResultManager{}
		analyzerInput := &models.AnalyzerInput{
			HtmlDoc: htmlDoc,
		}

		// Invoke
		analyzer.Analyze(analyzerInput, arm)

		// Verify
		actual := arm.GetAnalyzerResult().HtmlVersion
		assert.Equal(t, sample.expected, actual, "Expected: %v", sample.expected)
	}
}

// Given: An HTML documents containing <!DOCTYPE> with unknown PUBLIC attribute value
// When: The HtmlVersionAnalyzer's Analyze method is invoked on this document
// Then: It should correctly detect, but return with extra message saying that the version is Unknown
func TestHtmlVersionAnalyzer_Analyze_UnKnown_Versions(t *testing.T) {
	htmlContent := "<!doctype HTML PUBLIC \"-//DUMMY\"> <html><head></head><body></body></html>"
	expected := "Unknown Version: -//DUMMY"

	htmlDoc, err := html.Parse(strings.NewReader(htmlContent))
	require.NoError(t, err, "Failed to prepare test input")

	// Prepare Analyzer
	analyzer := HtmlVersionAnalyzer{}
	arm := &AnalyzerResultManager{}
	analyzerInput := &models.AnalyzerInput{
		HtmlDoc: htmlDoc,
	}

	// Invoke
	analyzer.Analyze(analyzerInput, arm)

	// Verify
	actual := arm.GetAnalyzerResult().HtmlVersion
	assert.Equal(t, expected, actual, "Expected: %v", expected)
}
