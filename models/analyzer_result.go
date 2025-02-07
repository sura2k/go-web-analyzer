package models

// Main AnalyzerResult struct
type AnalyzerResult struct {
	TargetUrl     string
	HtmlVersion   string
	PageTitle     string
	Headings      Headings
	Links         Links
	HasLoginForm  bool
	Status        bool
	StatusMessage string
}

// Headings struct to store heading count by level
type Headings struct {
	H1Count int
	H2Count int
	H3Count int
	H4Count int
	H5Count int
	H6Count int
}

// Links struct to store internal, external links counts
type Links struct {
	External      LinkCount
	Internal      LinkCount
	EmptyLinks    LinkCount
	NonHyperLinks LinkCount
}

// Link struct to store total and inaccessible links counts
type LinkCount struct {
	Total        int
	Inaccessible int
}
