package main

import (
	"math"
	"math/rand"
	"time"
)

const (
	DatasetLen        = 50000
	LimitationSteps   = 1000
	StatisticRangeMin = 0.01
	StatisticRangeMax = 0.97

	SideOne   = 1
	SideTwo   = 2
	SideThree = 3
	SideFour  = 4
	SideFive  = 5
	SideSix   = 6
)

// calculate new fractional and decisive number
func random(fractionalPrev, coefficient float64) (integer, fractionalNext float64) {
	// generate random element
	_, fractionalPrev = math.Modf(math.Pow(fractionalPrev+coefficient, 8))

	// getting two first digital to check hitting to interval
	integer, _ = math.Modf(fractionalPrev * 100)

	// return new fractionalPrev and first two digital for new fractionalPrev
	return integer, fractionalNext
}

// imitation roll dice
func rollDice(channelForRoll chan float64) {
	var (
		integer, fractional, coefficient float64
	)

	// init seed element
	fractional = rand.New(rand.NewSource(time.Now().UnixNano())).Float64()

	// reading coefficient
	coefficient = rand.New(rand.NewSource(time.Now().UnixNano())).Float64()*8 + 2

	// calculate new fractional for next calculation and return decisive number
	integer, fractional = random(fractional, coefficient)

	channelForRoll <- integer
	channelForRoll <- fractional
}

// getting result dice side
func mostOftenDropEdge() (diceEdge int) {
	var (
		integer, fractional float64
		distribution        map[int]int
		max                 int
	)

	distribution = make(map[int]int)

	channelForRoll := make(chan float64, 2)

	// generate randomly fifty elements and calculates interval hitting
	for i := 0; i < DatasetLen/LimitationSteps; i++ {
		// getting values from channel
		go rollDice(channelForRoll)
		integer, fractional = <-channelForRoll, <-channelForRoll

		// needing result – any numbers less than StatisticRangeMax or more then StatisticRangeMin
		if (fractional >= StatisticRangeMax) || (fractional < StatisticRangeMin) {
			// repeat calculation while not get needing result
			for {
				// getting values from channel
				go rollDice(channelForRoll)
				integer, fractional = <-channelForRoll, <-channelForRoll

				if fractional < StatisticRangeMax {
					break
				}
			}

		}

		// increase range in which we hit
		switch {
		case integer >= 1 && integer <= 16:
			distribution[SideOne] = distribution[SideOne] + 1
		case integer >= 17 && integer <= 32:
			distribution[SideTwo] = distribution[SideTwo] + 1
		case integer >= 33 && integer <= 48:
			distribution[SideThree] = distribution[SideThree] + 1
		case integer >= 49 && integer <= 64:
			distribution[SideFour] = distribution[SideFour] + 1
		case integer >= 65 && integer <= 80:
			distribution[SideFive] = distribution[SideFive] + 1
		case integer >= 81 && integer <= 96:
			distribution[SideSix] = distribution[SideSix] + 1
		}
	}

	// find most often drop side
	for key, value := range distribution {
		if max < value {
			max = value
			diceEdge = key
		}
	}

	return diceEdge
}
