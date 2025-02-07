package analyzers

import (
	"go-web-analyzer/models"
	"log"
	"sync"
)

// AnalyzerManager runs all analyzers parallaly
type AnalyzerManager struct {
	targetUrl string
}

// NewAnalyzerManager initializes the AnalyzerManager with targetUrl
func NewAnalyzerManager(targetUrl string) *AnalyzerManager {
	return &AnalyzerManager{targetUrl: targetUrl}
}

// ExecuteAnalyzers executes all analyzers parallaly
func (am *AnalyzerManager) ExecuteAnalyzers() (*models.AnalyzerResult, error) {
	log.Println("AnalyzerManager: Started. targetUrl: ", am.targetUrl)

	// List of available analyzers which are to be executed parallaly
	analyzers := []IAnalyzer{
		LinksAnalyzer{},
		HeadingTagsAnalyzer{},
		LoginFormAnalyzer{},
		PageTitleAnalyzer{},
		HtmlVersionAnalyzer{},
	}

	// Execute Preprocessor
	analyzerPreprocessor := NewAnalyzerPreprocessor(am.targetUrl)
	analyzerInput, err := analyzerPreprocessor.ExecutePreprocessor()
	if err != nil {
		return nil, err
	}

	//Starting all analyzers parallaly

	arm := &AnalyzerResultManager{}
	arm.SetTargetUrl(am.targetUrl) //Set the targetUrl before starting the analyzers

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
			a.Analyze(analyzerInput, arm)
		}(analyzer)
	}

	// Blocks the main thread until all goroutines are completed
	wg.Wait()

	log.Println("AnalyzerManager: Completed")

	return arm.GetAnalyzerResult(), nil
}
