package main

func getUnits(convType string) map[string]Unit {
	switch convType {
	case "length":
		return lengthUnits
	case "weight":
		return weightUnits
	case "temperature":
		return tempUnits
	default:
		return nil
	}
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
