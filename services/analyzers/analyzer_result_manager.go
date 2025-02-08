package analyzers

import (
	"sync"

	"github.com/sura2k/go-web-analyzer/models"
)

// AnalyzerResultManager manages the concurrent writes
type AnalyzerResultManager struct {
	lock           sync.Mutex
	analyzerResult models.AnalyzerResult
}

// Following Setters are to update AnalyzerResult fields safely

func (arm *AnalyzerResultManager) SetTargetUrl(targetUrl string) {
	arm.lock.Lock()         //Aquire Lock
	defer arm.lock.Unlock() //Release Lock when exiting the function
	arm.analyzerResult.TargetUrl = targetUrl
}

func (arm *AnalyzerResultManager) SetHtmlVersion(htmlVersion string) {
	arm.lock.Lock()
	defer arm.lock.Unlock()
	arm.analyzerResult.HtmlVersion = htmlVersion
}

func (arm *AnalyzerResultManager) SetPageTitle(pageTitle string) {
	arm.lock.Lock()
	defer arm.lock.Unlock()
	arm.analyzerResult.PageTitle = pageTitle
}

func (arm *AnalyzerResultManager) SetHeadings(headings *models.Headings) {
	arm.lock.Lock()
	defer arm.lock.Unlock()
	arm.analyzerResult.Headings = *headings
}

func (arm *AnalyzerResultManager) SetLinks(links *models.Links) {
	arm.lock.Lock()
	defer arm.lock.Unlock()
	arm.analyzerResult.Links = *links
}

func (arm *AnalyzerResultManager) SetHasLoginForm(hasLoginForm bool) {
	arm.lock.Lock()
	defer arm.lock.Unlock()
	arm.analyzerResult.HasLoginForm = hasLoginForm
}

// GetAnalyzerResult returns the parallaly updated AnalyzerResult safely
func (arm *AnalyzerResultManager) GetAnalyzerResult() *models.AnalyzerResult {
	arm.lock.Lock()
	defer arm.lock.Unlock()
	return &arm.analyzerResult
}
