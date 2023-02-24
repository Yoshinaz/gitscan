package report

import "time"

type Info struct {
	Name        string
	URL         string
	Status      string
	Description string
	EnqueuedAt  time.Time
	StartedAt   time.Time
	FinishedAt  time.Time
	CreatedAt   time.Time
	Findings    []Finding
}

type Finding struct {
	Type     string
	RuleId   string
	Location []Location
	Commit   string
	Metadata Metadata
}

type Location struct {
	Path      string
	Positions Position
}

type Position struct {
	Begin []Begin
}

type Begin struct {
	Line string
}

type Metadata struct {
	Description string
	Severity    string
}
