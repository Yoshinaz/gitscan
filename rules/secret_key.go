package rules

import (
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"guardrail/gitscan/internal/report"
	"strconv"
	"strings"
)

type SecretKey struct {
	Description string
	Severity    string
	RuleID      string
	Type        string
}

func (r SecretKey) Process(file *object.File) (report.Location, bool, error) {
	lines, err := file.Lines()
	if err != nil {
		return report.Location{}, false, err
	}
	position := make([]report.Begin, 0)
	for i, v := range lines {
		if strings.Index(v, "public_key") != -1 {
			position = append(position, report.Begin{Line: strconv.Itoa(i)})
		}

		if strings.Index(v, "private_key") != -1 {
			position = append(position, report.Begin{Line: strconv.Itoa(i)})
		}
	}

	return report.Location{
		Path:      file.Name,
		Positions: report.Position{Begin: position},
	}, len(position) != 0, nil
}

func (r SecretKey) GetRuleInfo() RuleInfo {
	return RuleInfo{
		Description: r.Description,
		Severity:    r.Severity,
		RuleID:      r.RuleID,
		Type:        r.Type,
	}
}
