package rules

import (
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"guardrail/gitscan/internal/report"
)

type RuleInfo struct {
	Description string
	Severity    string
	RuleID      string
	Type        string
}

type RuleInterface interface {
	Process(file *object.File) (report.Location, bool, error)
	GetRuleInfo() RuleInfo
}

var rules []RuleInterface

func init() {
	rules = make([]RuleInterface, 0)
	rules = append(rules, SecretKey{
		Description: "private / public key detected",
		Severity:    "HIGH",
		RuleID:      "G401",
		Type:        "sast",
	})
}

func GetMetaData(ruleID string) report.Metadata {
	for _, v := range rules {
		ruleInfo := v.GetRuleInfo()
		if ruleInfo.RuleID == ruleID {
			return report.Metadata{
				Description: ruleInfo.Description,
				Severity:    ruleInfo.Severity,
			}
		}
	}

	return report.Metadata{}
}

func Process(c *object.Commit) ([]report.Finding, error) {
	findings := make([]report.Finding, 0)
	fIter, err := c.Files()
	if err != nil {
		return findings, err
	}
	detected := make(map[RuleInterface][]report.Location)

	//detect each file with all rules
	err = fIter.ForEach(func(file *object.File) error {
		for _, rule := range rules {
			location, found, err := rule.Process(file)
			if err != nil {
				return err
			}
			if !found {
				continue
			}
			//add detected location into existing list or create a new one if it was a first time
			if _, ok := detected[rule]; ok {
				detected[rule] = append(detected[rule], location)
			} else {
				l := make([]report.Location, 0)
				l = append(l, location)
				detected[rule] = l
			}
		}
		return nil
	})

	if err != nil {
		return findings, err
	}

	//convert the detected info to finding report
	for rule, locations := range detected {
		ruleInfo := rule.GetRuleInfo()
		finding := report.Finding{
			Type:     ruleInfo.Type,
			RuleId:   ruleInfo.RuleID,
			Location: locations,
			Commit:   c.Hash.String(),
			Metadata: report.Metadata{
				Description: ruleInfo.Description,
				Severity:    ruleInfo.Severity,
			},
		}
		findings = append(findings, finding)
	}

	return findings, nil
}
