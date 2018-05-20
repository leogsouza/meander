package meander

import (
	"errors"
	"strings"
)

// ParseCost returns the cost which corresponds to string parameter
func ParseCost(s string) Cost {
	return costStrings[s]
}

// CostRange represents the range of costs
type CostRange struct {
	From Cost
	To   Cost
}

func (r CostRange) String() string {
	return r.From.String() + "..." + r.To.String()
}

// ParseCostRange parses the costRange string into CostRange type
func ParseCostRange(s string) (CostRange, error) {
	var r CostRange
	segs := strings.Split(s, "...")
	if len(segs) != 2 {
		return r, errors.New("invalid cost range")
	}
	r.From = ParseCost(segs[0])
	r.To = ParseCost(segs[1])
	return r, nil
}
