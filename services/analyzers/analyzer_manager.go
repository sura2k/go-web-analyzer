package analyzers

import (
	"log"

	"github.com/sura2k/go-web-analyzer/models"
)

// AnalyzerManager manages the analyzer workflow
type AnalyzerManager struct {
	targetUrl string
}

// NewAnalyzerManager initializes the AnalyzerManager with targetUrl
func NewAnalyzerManager(targetUrl string) *AnalyzerManager {
	return &AnalyzerManager{targetUrl: targetUrl}
}

// Start executes the analyzer workflow
func (am *AnalyzerManager) Start() (*models.AnalyzerResult, error) {
	log.Println("AnalyzerManager: Started. targetUrl: ", am.targetUrl)

	// Invoke AnalyzerPreprocessor
	analyzerPreprocessor := NewAnalyzerPreprocessor(am.targetUrl)
	analyzerInput, err := analyzerPreprocessor.ExecutePreprocessor()
	if err != nil {
		return nil, err
	}

	// Invoke AnalyzerExecutor
	analyzerExecutor := NewAnalyzerExecutor(analyzerInput)
	analyzerResult, err := analyzerExecutor.ExecuteAnalyzers()
	if err != nil {
		return nil, err
	}

	log.Println("AnalyzerManager: Completed")

	return analyzerResult, nil
}
