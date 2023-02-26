package reqError

type Req int

const (
	URL_ERROR Req = iota
	HEAD_ERROR
	INTERNAL_ERROR
)

func (t Req) String() string {
	return []string{
		"URL ERROR",
		"COULD NOT RETRIEVE HEAD",
		"INTERNAL ERROR",
	}[t]
}
