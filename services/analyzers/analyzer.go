package analyzers

import "go-web-analyzer/models"

// IAnalyzer interface
type IAnalyzer interface {
	Analyze(analyzerInput *models.AnalyzerInput, arManager *AnalyzerResultManager)
}
