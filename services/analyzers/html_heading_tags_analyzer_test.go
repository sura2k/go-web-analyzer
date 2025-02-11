package analyzers

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sura2k/go-web-analyzer/models"
	"golang.org/x/net/html"
)

// Given: An HTML documents containing all the types of heading tags
// When: The HeadingTagsAnalyzer's Analyze method is invoked on this document
// Then: It should correctly detect all the heading tags, count them and return
func TestHtmlHeadingTagsAnalyzer_Analyze(t *testing.T) {
	htmlContent := `<!doctype>
		<html>
			<head></head>
			<body>
				<div class="header">
					<h2>test2</h2>
					<h1>test1</h1>
				</div>
				<h3>test3</h3>
				<h2>test2.2</h2>
				<h4>test4</h4>
				<h3>test3.2</h3>
				<div>
					<h6>test6</h6>
					<h3>test3.3</h3>
					<h6>test6.2</h6>
				</div>
				<h3>test3.4</h3>
				<h2>test2.3</h2>
			</body>
		</html>`

	// Expected Result
	expected := models.Headings{
		H1Count: 1,
		H2Count: 3,
		H3Count: 4,
		H4Count: 1,
		H5Count: 0,
		H6Count: 2,
	}

	htmlDoc, err := html.Parse(strings.NewReader(htmlContent))
	require.NoError(t, err, "Failed to prepare test input")

	// Prepare Analyzer
	analyzer := HeadingTagsAnalyzer{}
	arm := &AnalyzerResultManager{}
	analyzerInput := &models.AnalyzerInput{
		HtmlDoc: htmlDoc,
	}

	// Invoke
	analyzer.Analyze(analyzerInput, arm)

	// Verify
	actual := arm.GetAnalyzerResult().Headings
	assert.Equal(t, expected, actual, "Expected: %v", expected)
}
