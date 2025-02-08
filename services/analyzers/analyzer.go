package analyzers

import "github.com/sura2k/go-web-analyzer/models"

// IAnalyzer interface
type IAnalyzer interface {
	Analyze(analyzerInput *models.AnalyzerInput, arManager *AnalyzerResultManager)
}
