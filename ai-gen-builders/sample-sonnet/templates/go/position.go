
type Position struct {
	Offset         int
	furthestFailed int
	expected       []Expected
}

func (p *Position) furthestFailure() int {
	if p.furthestFailed == 0 {
		return -1
	}
	return p.furthestFailed
}

type Expected struct {
	Name  string
	Value string
}
