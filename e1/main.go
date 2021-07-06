/*
Упражнение 1.  Поведение последовательности.

Речь идет о том, чтобы увидеть, как ведут себя числовые последовательности, порожденные уравнением.
Для этого вычислим большое число членов последовательности, порожденной своим первым элементом. Поместим каждый из этих
членов в один из 50 интервалов длины 0.02, составляющих интервал от 0 до 1. Выведем число членов последовательности,
попавших в каждый из этих интервалов. Если числа из последовательности равномерно распределены в интервале (0, 1), мы
должны будем обнаружить, что их количество в разных интервалах имеет ощутимую тенденцию к постоянству.
Составьте программу для проверки зтого утверждения. Начальное значение может, например, вводиться в начале каждого
вычисления.
 */

package main

import (
	"fmt"
	"math"
)

const (
	SequenceStart    = 0.5
	DistributionStep = 2
	DatasetLen       = 10000
)

func main() {
	var (
		integer, fractional float64

		distribution map[float64]int
	)

	distribution = make(map[float64]int)

	// init distribution keys with step equals DistributionStep (in real step equals DistributionStep/100)
	for integer < 100 {
		integer = integer + DistributionStep
		distribution[integer] = 0
	}

	// generate randomly fifty elements and calculates interval hitting
	fractional = SequenceStart
	for i := 0; i < DatasetLen; i++ {
		var (
			ok bool
		)
		
		// generate random element
		_, fractional = math.Modf(math.Pow(fractional + math.Pi, 8))

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
	for key, value := range distribution {
		fmt.Printf("hit to inerval %.2f – %.2f: %d (%.2f%%)\n", (key - DistributionStep)/100, key/100, value, float64(value)*100/DatasetLen)
	}

	fmt.Printf("\nwork completed")
}
