package rules

import (
	"github.com/gitscan/internal/report"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"strconv"
	"strings"
)

type SecretKey struct {
	Description string
	Severity    string
	RuleID      string
	Type        string
}

func NewSecretKey() RuleInfoInterface {
	return SecretKey{
		Description: "private / public key detected",
		Severity:    "HIGH",
		RuleID:      "G401",
		Type:        "sast",
	}
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

func (r SecretKey) GetRuleInfo() ruleInfo {
	return ruleInfo{
		Description: r.Description,
		Severity:    r.Severity,
		RuleID:      r.RuleID,
		Type:        r.Type,
	}
}
