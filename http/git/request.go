package git

type Request struct {
	Name          string `json:"name"`
	URL           string `json:"url"`
	RulesSet      string `json:"rules_set"`
	ScanAllCommit bool   `json:"scan_all_commit"`
}
