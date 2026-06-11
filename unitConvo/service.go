package main

import (
	"fmt"
)

func convertGeneric(n float64, from, to string, units map[string]Unit) (float64, error) {
	if from == to {
		return n, nil
	}
	fromUnit, ok := units[from]
	if !ok {
		return 0, fmt.Errorf("invalid unit: %s does not exist", from)
	}
	toUnit, ok := units[to]
	if !ok {
		return 0, fmt.Errorf("invalid unit: %s does not exist", to)
	}
	base := fromUnit.ToBase(n)
	return toUnit.FromBase(base), nil
}
