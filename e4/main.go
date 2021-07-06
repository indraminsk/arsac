/*
Упражнение 4.  Игральные кости.

Вместо игры в «орла» или «решку» заставьте компьютер играть в кости. Напишите программу, симулирующую большое число
выбрасываний двух костей, и подсчитайте, сколько раз будет выпадать каждая комбинация от 2 до 12.
 */

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	DatasetLen        = 50000
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
		randSource  rand.Source
	)

	// init seed element
	randSource = rand.NewSource(time.Now().UnixNano())
	fractional = rand.New(randSource).Float64()

	// reading coefficient
	randSource = rand.NewSource(time.Now().UnixNano())
	coefficient = rand.New(randSource).Float64()*8 + 2

	// calculate new fractional for next calculation and return decisive number
	integer, fractional = random(fractional, coefficient)

	channelForRoll <- integer
	channelForRoll <- fractional
}

// getting result dice side
func mostOftenDropEdge(channelForEdge chan int) {
	var (
		integer, fractional float64
		distribution        map[int]int
		max, diceSide       int
	)

	distribution = make(map[int]int)

	channelForRoll := make(chan float64, 2)

	// generate randomly fifty elements and calculates interval hitting
	for i := 0; i < DatasetLen/1000; i++ {
		// getting values from channel
		go rollDice(channelForRoll)
		integer, fractional = <- channelForRoll, <- channelForRoll

		// needing result – any numbers less than StatisticRangeMax or more then StatisticRangeMin
		if (fractional >= StatisticRangeMax) || (fractional < StatisticRangeMin) {
			// repeat calculation while not get needing result
			for {
				// getting values from channel
				go rollDice(channelForRoll)
				integer, fractional = <- channelForRoll, <- channelForRoll

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
			diceSide = key
		}
	}

	channelForEdge <- diceSide
}

func main() {
	var (
		diceEdgeA, diceEdgeB int
		distributionEdges    map[string]int
	)

	distributionEdges = make(map[string]int)
	channelForEdge := make(chan int)

	fmt.Println("let's rolling dices")

	for i := 0; i < DatasetLen; i++ {
		var (
			key string
		)

		go mostOftenDropEdge(channelForEdge)
		diceEdgeA = <- channelForEdge

		go mostOftenDropEdge(channelForEdge)
		diceEdgeB = <- channelForEdge

		key = fmt.Sprintf("%d, %d", diceEdgeA, diceEdgeB)

		distributionEdges[key] = distributionEdges[key] + 1
	}

	// output the result roll of dices
	for key, value := range distributionEdges {
		fmt.Printf("dice edges [%s] rolls %d (%.2f%%)\n", key, value, float64(value)*100/DatasetLen)
	}

	fmt.Println()
}
