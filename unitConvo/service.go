package main

import (
	"fmt"
)

func convert(n float64, from, to, convType string) (float64, error) {
	switch convType {
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
func convertLength(n float64, from, to string) (float64, error) {
	var result float64
	switch from {
	case "metres":
		result = n
	case "kilometres":
		result = n * 1000
	case "centimetres":
		result = n * 100
	default:
		return 0, fmt.Errorf("invalid conversion type")
	}
	switch to {
	case "metres":
		return result, nil
	case "kilometres":
		return result * 1000, nil
	case "centimetres":
		return result * 100, nil
	default:
		return 0, fmt.Errorf("invalid conversion type")
	}
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
