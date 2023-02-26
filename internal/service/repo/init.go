package repo

import "github.com/gitscan/rules"

func (r *Repo) Init(name, url string, rules rules.Interface) {
	r.Name = name
	r.URL = url
	r.Rules = rules
}
