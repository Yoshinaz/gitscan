package status

type Info int

const (
	QUEUED Info = iota
	INPROGRESS
	SUCCESS
	FAILED
)

func (t Info) String() string {
	return []string{
		"QUEUED",
		"INPROGRESS",
		"SUCCESS",
		"FAILED",
	}[t]
}
