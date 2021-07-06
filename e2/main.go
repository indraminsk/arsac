/*
Упражнение 2.  Поиск других последовательностей.

Измените программу Упражнения 1 так, чтобы она осуществляла чтение
— постоянной а,
— начального значения последовательности.
 */

package main

import (
	"fmt"
	"math"
)

const (
	DefaultSeedValue = 0.5
	DefaultCoefficient = math.Pi
	DistributionStep = 2
	DatasetLen       = 50
)

func main() {
	var (
		integer, fractional, coefficient  float64

		distribution map[float64]int
	)

	fmt.Println()
	fmt.Println("working expression: (x+a)^8")

	distribution = make(map[float64]int)

	// init distribution keys with step equals DistributionStep (in real step equals DistributionStep/100)
	for integer < 100 {
		integer = integer + DistributionStep
		distribution[integer] = 0
	}

	// reading seed element
	fmt.Print("init random sequence (x): ")
	_, _ = fmt.Scanf("%f", &fractional)

	// when seed element not set
	if fractional == 0 {
		fractional = DefaultSeedValue
	}

	// reading coefficient
	fmt.Print("init coefficient (a): ")
	_, _ = fmt.Scanf("%f", &coefficient)

	// when coefficient not set
	if coefficient == 0 {
		coefficient = DefaultCoefficient
	}

	// generate randomly fifty elements and calculates interval hitting
	for i := 0; i < DatasetLen; i++ {
		var (
			ok bool
		)

		// generate random element
		_, fractional = math.Modf(math.Pow(fractional + coefficient, 8))

		// getting two first digital to check hitting to interval
		integer, _ = math.Modf(fractional * 100)
		_, ok = distribution[integer]
		if !ok {
			if integer != 0 {
				// we have an odd number but we have interest only even numbers so increase odd to the closest even
				integer++
			} else {
				// force convert 0 to DistributionStep
				integer = DistributionStep
			}
		}

		// store hitting
		distribution[integer] = distribution[integer] + 1
	}

	// generate distribution report
	fmt.Println()
	s := 1
	for key, value := range distribution {
		if value == 0 {
			continue
		}

		fmt.Printf("[%2d] hit to inerval %.2f – %.2f: %d\n", s, (key - DistributionStep)/100, key/100, value)
		s++
	}

	fmt.Printf("\nwork completed")
}
