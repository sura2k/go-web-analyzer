package analyzers

import (
	"log"
	"sync"

	"github.com/sura2k/go-web-analyzer/models"
)

// AnalyzerExecutor runs all analyzers parallaly
type AnalyzerExecutor struct {
	analyzerInput *models.AnalyzerInput
}

// NewAnalyzerExecutor initializes the AnalyzerExecutor with analyzerInput
func NewAnalyzerExecutor(analyzerInput *models.AnalyzerInput) *AnalyzerExecutor {
	return &AnalyzerExecutor{analyzerInput: analyzerInput}
}

// ExecuteAnalyzers executes all analyzers parallaly
func (ae *AnalyzerExecutor) ExecuteAnalyzers() (*models.AnalyzerResult, error) {
	log.Println("AnalyzerExecutor: Started. targetUrl: ", ae.analyzerInput.TargetUrl)

	// List of available analyzers which are to be executed parallaly
	analyzers := []IAnalyzer{
		LinksAnalyzer{},
		HeadingTagsAnalyzer{},
		LoginFormAnalyzer{},
		PageTitleAnalyzer{},
		HtmlVersionAnalyzer{},
	}

	// Ready to start all analyzers parallaly

	arm := &AnalyzerResultManager{}
	arm.SetTargetUrl(ae.analyzerInput.TargetUrl) //Set the targetUrl before starting the analyzers

	var wg sync.WaitGroup

	// ----------
	// IMPORTANT:
	// ----------
	//	Set all available analyzers count to WaitGroup before the loop,
	// 	to ensure all goroutines/threads are considered before the WaitGroup start.
	//
	// 	If wg.Add(1) is added inside the goroutine block, there can be an problem where
	// 	wg.Wait() might execute before wg.Add(1), and then it will lead to a deadlock,
	// 	since WorkGroup itself might think that there are no goroutines to wait for.
	// ----------
	wg.Add(len(analyzers))

	for _, analyzer := range analyzers {
		go func(a IAnalyzer) {
			defer wg.Done() // Each goroutine notifies to WaitGroup that it has completed its task
			a.Analyze(ae.analyzerInput, arm)
		}(analyzer)
	}

	// Blocks the main thread until all goroutines are completed
	wg.Wait()

	log.Println("AnalyzerExecutor: Completed")

	return arm.GetAnalyzerResult(), nil
}
