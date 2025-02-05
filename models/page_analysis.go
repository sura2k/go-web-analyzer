package models

// PageAnalysis struct for passing data to the template
type PageAnalysis struct {
	TargetUrl        string
	HtmlVersion      string
	PageTitle        string
	Headings         Headings
	Links            Links
	HasLoginForm     bool
	RedirectsToLogin bool
	Status           bool
	StatusMessage    string
}

// Headings struct to store heading count by level
type Headings struct {
	NumOfH1 int
	NumOfH2 int
	NumOfH3 int
	NumOfH4 int
	NumOfH5 int
	NumOfH6 int
}

// Links struct to store internal, external, inaccessible links count
type Links struct {
	NumOfIntLinks             int
	NumOfExtLinks             int
	NumOfIntLinksInaccessible int
	NumOfExtLinksInaccessible int
}
