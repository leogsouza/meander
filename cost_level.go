package meander

// Cost represents the costs from a place
type Cost int8

// Cost Enumerator
const (
	_ Cost = iota
	Cost1
	Cost2
	Cost3
	Cost4
	Cost5
)

var costStrings = map[string]Cost{
	"$":     Cost1,
	"$$":    Cost2,
	"$$$":   Cost3,
	"$$$$":  Cost4,
	"$$$$$": Cost5,
}

func (c Cost) String() string {
	for s, v := range costStrings {
		if c == v {
			return s
		}
	}

	return "invalid"
}
