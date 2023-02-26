package repo

type Request struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	RulesID string `json:"rules_id"`
}
