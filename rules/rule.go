package rules

import (
	"github.com/gitscan/internal/report"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

type ruleInfo struct {
	Description string
	Severity    string
	RuleID      string
	Type        string
}

type RuleInfoInterface interface {
	Process(file *object.File) (report.Location, bool, error)
	GetRuleInfo() ruleInfo
}
