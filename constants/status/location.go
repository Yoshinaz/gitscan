package status

type Location int

const (
	ACTIVE Location = iota
	IGNORE
)

func (t Location) String() string {
	return []string{
		"ACTIVE",
		"IGNORE",
	}[t]
}
