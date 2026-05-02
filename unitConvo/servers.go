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
