package main

import "math"

// generate random value by formula: (x + const)^8
func ale(x float64) float64 {
	var (
		fractional float64
	)

	// generate random element
	_, fractional = math.Modf(math.Pow(x + math.Pi, 8))

	return fractional
}