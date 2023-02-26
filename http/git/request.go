package git

type Request struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	RulesSet  string `json:"rules_set"`
	AllCommit string `json:"all_commit"`
}
