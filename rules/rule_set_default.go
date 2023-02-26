package rules

// Custom rules for each request depend on pre define rules set,
func (r Rules) defaultSet() Interface {
	r.Add(NewSecretKey())

	return r
}
