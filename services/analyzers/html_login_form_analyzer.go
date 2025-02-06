package analyzers

import (
	"golang.org/x/net/html"
)

// Returns whether give htmlDoc contains a login form
// This function assumes that if the following conditions are true, then HTML page has a login form
//   - <form> element
//   - <input type="password"> inside the same <form>
//   - <input type="text"> or <input type="email"> inside the same <form>
//   - <button type="submit"> or <input type="submit"> inside the same <form>
func HasLoginForm(analyzerInfo *AnalyzerInfo) bool {

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

				// Once all the necessary form-login elements are found,
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
				return false // If form-login elements are not found, returns false otherwise
			}

			//Start the traversing through <form> children once <form> is detected
			if traverseFormElem(node) {
				return true
			}
		}

		// Loop through the document tree
		// Note:
		//		If the root is <!doctype>, then loop will skip for the <!doctype>
		// 		since there are no child elements in <!doctype>
		for elem := node.FirstChild; elem != nil; elem = elem.NextSibling {
			if traverse(elem) {
				return true
			}
		}
		return false
	}

	//Start traversing from the root
	return traverse(analyzerInfo.htmlDoc)
}
