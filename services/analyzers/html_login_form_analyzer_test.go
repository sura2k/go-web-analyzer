package analyzers

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sura2k/go-web-analyzer/models"
	"golang.org/x/net/html"
)

// Given: An HTML document containing a "form" with an input "text", "password" and a "submit" fields
// When: The LoginFormAnalyzer's Analyze method is invoked on this document
// Then: It should correctly detect that the page contains a login form
func TestHtmlLoginFormAnalyzer_Analyze__When_FormAndInputTextPasswordSubmitAvailable(t *testing.T) {
	htmlContent := `<!doctype>
		<html>
			<head></head>
			<body>
				<div class="header">
				</div>
				<div>
					<form>
						<input type="text"></input>
						<input type="password"></input>
						<input type="submit"></input>
					</form>
				</div>
			</body>
		</html>`

	// Expected Result
	expected := true

	htmlDoc, err := html.Parse(strings.NewReader(htmlContent))
	require.NoError(t, err, "Failed to prepare test input")

	// Prepare Analyzer
	analyzer := LoginFormAnalyzer{}
	arm := &AnalyzerResultManager{}
	analyzerInput := &models.AnalyzerInput{
		HtmlDoc: htmlDoc,
	}

	// Invoke
	analyzer.Analyze(analyzerInput, arm)

	// Verify
	actual := arm.GetAnalyzerResult().HasLoginForm
	assert.Equal(t, expected, actual, "Expected: %v", expected)
}

// Given: An HTML document containing a "form" with an input "email", "password" and a "submit" fields
// When: The LoginFormAnalyzer's Analyze method is invoked on this document
// Then: It should correctly detect that the page contains a login form
func TestHtmlLoginFormAnalyzer_Analyze__When_FormAndInputEmailPasswordSubmitAvailable(t *testing.T) {
	htmlContent := `<!doctype>
		<html>
			<head></head>
			<body>
				<div class="header">
				</div>
				<div>
					<form>
						<input type="email"></input>
						<input type="password"></input>
						<input type="submit"></input>
					</form>
				</div>
			</body>
		</html>`

	// Expected Result
	expected := true

	htmlDoc, err := html.Parse(strings.NewReader(htmlContent))
	require.NoError(t, err, "Failed to prepare test input")

	// Prepare Analyzer
	analyzer := LoginFormAnalyzer{}
	arm := &AnalyzerResultManager{}
	analyzerInput := &models.AnalyzerInput{
		HtmlDoc: htmlDoc,
	}

	// Invoke
	analyzer.Analyze(analyzerInput, arm)

	// Verify
	actual := arm.GetAnalyzerResult().HasLoginForm
	assert.Equal(t, expected, actual, "Expected: %v", expected)
}

// Given: An HTML document containing a "form" with an input "email", "password" fields and and "submit" button
// When: The LoginFormAnalyzer's Analyze method is invoked on this document
// Then: It should correctly detect that the page contains a login form
func TestHtmlLoginFormAnalyzer_Analyze__When_FormAndInputEmailPasswordButtonSubmitAvailable(t *testing.T) {
	htmlContent := `<!doctype>
		<html>
			<head></head>
			<body>
				<div class="header">
				</div>
				<div>
					<form>
						<input type="email"></input>
						<input type="password"></input>
						<button type="submit"></button>
					</form>
				</div>
			</body>
		</html>`

	// Expected Result
	expected := true

	htmlDoc, err := html.Parse(strings.NewReader(htmlContent))
	require.NoError(t, err, "Failed to prepare test input")

	// Prepare Analyzer
	analyzer := LoginFormAnalyzer{}
	arm := &AnalyzerResultManager{}
	analyzerInput := &models.AnalyzerInput{
		HtmlDoc: htmlDoc,
	}

	// Invoke
	analyzer.Analyze(analyzerInput, arm)

	// Verify
	actual := arm.GetAnalyzerResult().HasLoginForm
	assert.Equal(t, expected, actual, "Expected: %v", expected)
}

// Given: An HTML document containing a "form" but no "submit" field/button
// When: The LoginFormAnalyzer's Analyze method is invoked on this document
// Then: It should correctly detect that the page DOES NOT contain a login form
func TestHtmlLoginFormAnalyzer_Analyze__When_FormHasNoSubmit(t *testing.T) {
	htmlContent := `<!doctype>
		<html>
			<head></head>
			<body>
				<div class="header">
				</div>
				<div>
					<form>
						<input type="text"></input>
						<input type="email"></input>
						<input type="password"></input>
					</form>
				</div>
			</body>
		</html>`

	// Expected Result
	expected := false

	htmlDoc, err := html.Parse(strings.NewReader(htmlContent))
	require.NoError(t, err, "Failed to prepare test input")

	// Prepare Analyzer
	analyzer := LoginFormAnalyzer{}
	arm := &AnalyzerResultManager{}
	analyzerInput := &models.AnalyzerInput{
		HtmlDoc: htmlDoc,
	}

	// Invoke
	analyzer.Analyze(analyzerInput, arm)

	// Verify
	actual := arm.GetAnalyzerResult().HasLoginForm
	assert.Equal(t, expected, actual, "Expected: %v", expected)
}

// Given: An HTML document containing a "form" but no input "password" field
// When: The LoginFormAnalyzer's Analyze method is invoked on this document
// Then: It should correctly detect that the page DOES NOT contain a login form
func TestHtmlLoginFormAnalyzer_Analyze__When_FormHasNoPassword(t *testing.T) {
	htmlContent := `<!doctype>
		<html>
			<head></head>
			<body>
				<div class="header">
				</div>
				<div>
					<form>
						<input type="text"></input>
						<input type="email"></input>
						<input type="submit"></input>
					</form>
				</div>
			</body>
		</html>`

	// Expected Result
	expected := false

	htmlDoc, err := html.Parse(strings.NewReader(htmlContent))
	require.NoError(t, err, "Failed to prepare test input")

	// Prepare Analyzer
	analyzer := LoginFormAnalyzer{}
	arm := &AnalyzerResultManager{}
	analyzerInput := &models.AnalyzerInput{
		HtmlDoc: htmlDoc,
	}

	// Invoke
	analyzer.Analyze(analyzerInput, arm)

	// Verify
	actual := arm.GetAnalyzerResult().HasLoginForm
	assert.Equal(t, expected, actual, "Expected: %v", expected)
}

// Given: An HTML document containing a "form" but no input "text" or "email" field
// When: The LoginFormAnalyzer's Analyze method is invoked on this document
// Then: It should correctly detect that the page DOES NOT contain a login form
func TestHtmlLoginFormAnalyzer_Analyze__When_FormHasNoTextOrEmail(t *testing.T) {
	htmlContent := `<!doctype>
		<html>
			<head></head>
			<body>
				<div class="header">
				</div>
				<div>
					<form>
						<input type="password"></input>
						<input type="submit"></input>
					</form>
				</div>
			</body>
		</html>`

	// Expected Result
	expected := false

	htmlDoc, err := html.Parse(strings.NewReader(htmlContent))
	require.NoError(t, err, "Failed to prepare test input")

	// Prepare Analyzer
	analyzer := LoginFormAnalyzer{}
	arm := &AnalyzerResultManager{}
	analyzerInput := &models.AnalyzerInput{
		HtmlDoc: htmlDoc,
	}

	// Invoke
	analyzer.Analyze(analyzerInput, arm)

	// Verify
	actual := arm.GetAnalyzerResult().HasLoginForm
	assert.Equal(t, expected, actual, "Expected: %v", expected)
}

// Given: An HTML document containing all the necessary input fields, but not inside a "form"
// When: The LoginFormAnalyzer's Analyze method is invoked on this document
// Then: It should correctly detect that the page DOES NOT contain a login form
func TestHtmlLoginFormAnalyzer_Analyze__When_AllFieldsAreAvailableWithoutForm(t *testing.T) {
	htmlContent := `<!doctype>
		<html>
			<head></head>
			<body>
				<div class="header">
					<form></form>
				</div>
				<div>
					<input type="text"></input>
					<input type="password"></input>
					<input type="submit"></input>
				</div>
			</body>
		</html>`

	// Expected Result
	expected := false

	htmlDoc, err := html.Parse(strings.NewReader(htmlContent))
	require.NoError(t, err, "Failed to prepare test input")

	// Prepare Analyzer
	analyzer := LoginFormAnalyzer{}
	arm := &AnalyzerResultManager{}
	analyzerInput := &models.AnalyzerInput{
		HtmlDoc: htmlDoc,
	}

	// Invoke
	analyzer.Analyze(analyzerInput, arm)

	// Verify
	actual := arm.GetAnalyzerResult().HasLoginForm
	assert.Equal(t, expected, actual, "Expected: %v", expected)
}

// Given: An HTML document containing multiple "form"s but proper login form
// When: The LoginFormAnalyzer's Analyze method is invoked on this document
// Then: It should correctly detect that the page DOES NOT contain a login form
func TestHtmlLoginFormAnalyzer_Analyze__When_MultipleFormsAvailableWithoutAnyLoginForm(t *testing.T) {
	htmlContent := `<!doctype>
		<html>
			<head></head>
			<body>
				<div class="header">
					<form>
						<input type="text"></input>
						<input type="submit"></input>
					</form>
				</div>
				<form></form>
				<div>
					<form>
						<input type="text"></input>
						<input type="email"></input>
						<input type="password"></input>
					</form>
				</div>
			</body>
		</html>`

	// Expected Result
	expected := false

	htmlDoc, err := html.Parse(strings.NewReader(htmlContent))
	require.NoError(t, err, "Failed to prepare test input")

	// Prepare Analyzer
	analyzer := LoginFormAnalyzer{}
	arm := &AnalyzerResultManager{}
	analyzerInput := &models.AnalyzerInput{
		HtmlDoc: htmlDoc,
	}

	// Invoke
	analyzer.Analyze(analyzerInput, arm)

	// Verify
	actual := arm.GetAnalyzerResult().HasLoginForm
	assert.Equal(t, expected, actual, "Expected: %v", expected)
}

// Given: An HTML document containing multiple "form"s and at least one form is a valid login form
// When: The LoginFormAnalyzer's Analyze method is invoked on this document
// Then: It should correctly detect that the page contains a login form
func TestHtmlLoginFormAnalyzer_Analyze__When_MultipleFormsAvailableWithAtleastOneLoginForm(t *testing.T) {
	htmlContent := `<!doctype>
		<html>
			<head></head>
			<body>
				<div class="header">
					<form>
						<input type="text"></input>
						<input type="text"></input>
					</form>
				</div>
				<form></form>
				<div>
					<form>
						<input type="text"></input>
						<input type="password"></input>
						<button type="submit"></input>
					</form>
				</div>
			</body>
		</html>`

	// Expected Result
	expected := true

	htmlDoc, err := html.Parse(strings.NewReader(htmlContent))
	require.NoError(t, err, "Failed to prepare test input")

	// Prepare Analyzer
	analyzer := LoginFormAnalyzer{}
	arm := &AnalyzerResultManager{}
	analyzerInput := &models.AnalyzerInput{
		HtmlDoc: htmlDoc,
	}

	// Invoke
	analyzer.Analyze(analyzerInput, arm)

	// Verify
	actual := arm.GetAnalyzerResult().HasLoginForm
	assert.Equal(t, expected, actual, "Expected: %v", expected)
}
