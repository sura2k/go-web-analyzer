package models

// Controller level model for view rendering and/or JSON rendering
type AnalyzerResponse struct {
	Processed bool
	Status    bool
	Message   string
	Data      AnalyzerResult
}
