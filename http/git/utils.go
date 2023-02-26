package git

import "github.com/gitscan/rules"

//Custom rules for each request depend on pre define rules set,
func getRules(rulesID string) rules.Interface {
	//current has only one set
	r := rules.New()
	r.Add(rules.NewSecretKey())
	return r
}
