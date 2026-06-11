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

var lengthUnits = map[string]Unit{
	"metres": {
		ToBase:   func(n float64) float64 { return n },
		FromBase: func(n float64) float64 { return n },
	},
	"centimetres": {
		ToBase:   func(n float64) float64 { return n * 100 },
		FromBase: func(n float64) float64 { return n / 100 },
	},
	"kilometres": {
		ToBase:   func(n float64) float64 { return n * 1000 },
		FromBase: func(n float64) float64 { return n / 1000 },
	},
}
var weightUnits = map[string]Unit{
	"kilograms": {
		ToBase:   func(n float64) float64 { return n },
		FromBase: func(n float64) float64 { return n },
	},
	"grams": {
		ToBase:   func(n float64) float64 { return n / 1000 },
		FromBase: func(n float64) float64 { return n * 1000 },
	},
	"pounds": {
		ToBase:   func(n float64) float64 { return n * 0.453592 },
		FromBase: func(n float64) float64 { return n * 2.20462 },
	},
}
var tempUnits = map[string]Unit{
	"celsius": {
		ToBase:   func(n float64) float64 { return n },
		FromBase: func(n float64) float64 { return n },
	},
	"kelvin": {
		ToBase:   func(n float64) float64 { return n - 273.15 },
		FromBase: func(n float64) float64 { return n + 273.15 },
	},
	"fahrenheit": {
		ToBase:   func(n float64) float64 { return (n - 32) * 5 / 9 },
		FromBase: func(n float64) float64 { return (n * 9 / 5) + 32 },
	},
}
