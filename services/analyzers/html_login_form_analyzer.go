package analyzers

import (
	"go-web-analyzer/models"
	"log"

	"golang.org/x/net/html"
)

// LoginFormAnalyzer struct
type LoginFormAnalyzer struct{}

// Analyze method updates the relevant field
func (a LoginFormAnalyzer) Analyze(analyzerInput *models.AnalyzerInput, arm *AnalyzerResultManager) {
	log.Println("LoginFormAnalyzer: Started")

	hasLoginForm := hasLoginForm(analyzerInput)
	arm.SetHasLoginForm(hasLoginForm)

	log.Println("LoginFormAnalyzer: Completed")
}

// Analyze and returns whether a login form is available
// This function assumes that if the following conditions are true, then HTML page has a login form
//   - <form> element
//   - <input type="password"> inside the same <form>
//   - <input type="text"> or <input type="email"> inside the same <form>
//   - <button type="submit"> or <input type="submit"> inside the same <form>
//
// Assumptions
//   - Considered hidden login form is valid login form
func hasLoginForm(analyzerInput *models.AnalyzerInput) bool {

	// Recursive function whcih traverse through document tree
	var traverse func(*html.Node) bool
	traverse = func(node *html.Node) bool {
		if node.Type == html.ElementNode && node.Data == "form" {
			hasInputPassword := false
			hasInputText := false
			hasSubmit := false

			// Recursive function whcih traverse through <form> element's children
			var traverseFormElem func(*html.Node) bool
			traverseFormElem = func(nodeFormElem *html.Node) bool {
				if nodeFormElem.Type == html.ElementNode && nodeFormElem.Data == "input" {
					for _, attr := range nodeFormElem.Attr {
						if attr.Key == "type" {
							if attr.Val == "password" {
								hasInputPassword = true
							} else if attr.Val == "text" || attr.Val == "email" {
								hasInputText = true
							} else if attr.Val == "submit" {
								hasSubmit = true
							}
						}
					}
				} else if nodeFormElem.Type == html.ElementNode && nodeFormElem.Data == "button" {
					for _, attr := range nodeFormElem.Attr {
						if attr.Key == "type" && attr.Val == "submit" {
							hasSubmit = true
						}
					}
				}

				// Once all the necessary login form elements are found inside a <form>,
				// no need further recursions, return immediate with true
				if hasInputPassword && hasInputText && hasSubmit {
					return true
				}

				// Loop through <form> element's children
				for formElem := nodeFormElem.FirstChild; formElem != nil; formElem = formElem.NextSibling {
					if traverseFormElem(formElem) {
						return true
					}
				}
				return false // Otherwise return false, if the login-form elements are not found in a <form>
			}

			//Start the traversing through <form> children till a <form> is detected
			if traverseFormElem(node) {
				return true
			}
		}

		// Loop through the document tree
		// Note:
		//		If the root is <!DOCTYPE>, then recursive travel will skip in the 2nd attempt
		// 		since there are no child elements in <!DOCTYPE>
		//		Then it will start the recursive travel starting from <html>
		for elem := node.FirstChild; elem != nil; elem = elem.NextSibling {
			if traverse(elem) {
				return true
			}
		}
		return false
	}

	//Start traversing from the root
	return traverse(analyzerInput.HtmlDoc)
}
