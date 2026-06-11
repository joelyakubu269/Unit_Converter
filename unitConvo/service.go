package main

import (
	"fmt"
)

func convert(n float64, from, to, convtype string) (float64, error) {
	//convtype:= "length"
	switch convtype {
	case "length":
		return convertLength(n, from, to)
	case "weight":
		return convertWeight(n, from, to)
	case "temperature":
		return convertTemperature(n, from, to)
	default:
		return 0, fmt.Errorf("invalid type")
	}
}

var lengthunits = map[string]Unit{
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

func convertTemperature(n float64, from, to string) (float64, error) {
	var result float64

	switch from {
	case "celsius":
		result = n
	case "kelvin":
		result = n - 273.15
	case "fahrenheit":
		result = (n - 32) * 5 / 9
	default:
		return 0, fmt.Errorf("invalid temperature unit")
	}
	switch to {
	case "celsius":
		return result, nil
	case "kelvin":
		return result + 273.15, nil
	case "fahrenheit":
		return (result * 9 / 5) + 32, nil
	default:
		return 0, fmt.Errorf("invalid temperature unit")
	}

}
func convertWeight(n float64, from, to string) (float64, error) {
	var kg float64

	switch from {
	case "kilograms":
		kg = n
	case "grams":
		kg = n / 1000
	case "pounds":
		kg = n * 0.453592
	default:
		return 0, fmt.Errorf("invalid weight unit")
	}

	switch to {
	case "kilograms":
		return kg, nil
	case "grams":
		return kg * 1000, nil
	case "pounds":
		return kg * 2.20462, nil
	default:
		return 0, fmt.Errorf("invalid weight unit")
	}
}
