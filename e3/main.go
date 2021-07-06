/*
Упражнение 3.  «Орел» или «решка».

Составьте следующую программу:
— она спрашивает вас, что вы загадали, «орла» или «решку», и читает ваш ответ;
— она порождает случайное число и затем сообщает вам, выиграли вы или проиграли.

Сделайте программу, реализующую большое число испытаний, и подсчитайте число выпаданий орла.
 */

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	DatasetLen = 500000

	Head = 0
	Tail = 1

	CounterHead = "head"
	CounterTail = "tail"
)

// calculate by goroutine new fractional and decisive number
func random(fractional, coefficient float64, channel chan float64) {
	var (
		integer float64
	)

	// generate random element
	_, fractional = math.Modf(math.Pow(fractional+coefficient, 8))

	// getting two first digital to check hitting to interval
	integer, _ = math.Modf(fractional * 100)

	// return new fractional and first two digital for new fractional
	channel <- integer
	channel <- fractional
}

func main() {
	var (
		integer, fractional, coefficient float64
		distribution                     map[string]float64
		userChoice, coinSide             int
		randSource                       rand.Source
	)

	fmt.Println("heads & tails is started")

	distribution = make(map[string]float64)

	// init distribution keys with step equals DistributionStep
	distribution[CounterHead] = 0
	distribution[CounterTail] = 0

	// init seed element
	randSource = rand.NewSource(time.Now().UnixNano())
	fractional = rand.New(randSource).Float64()

	// reading coefficient
	randSource = rand.NewSource(time.Now().UnixNano())
	coefficient = rand.New(randSource).Float64()*8 + 2

	// reading user choice
	fmt.Print("your choice (0 – head, 1 – tails): ")
	_, _ = fmt.Scanf("%d", &userChoice)

	// generate randomly fifty elements and calculates interval hitting
	channel := make(chan float64, 2)
	for i := 0; i < DatasetLen; i++ {
		// calculate new fractional for next calculation and return decisive number
		go random(fractional, coefficient, channel)

		// getting values from channel
		integer = <-channel
		fractional = <-channel

		// check random value is even or odd
		if (math.Mod(integer, 2)) == 0 {
			// even
			distribution[CounterHead] = distribution[CounterHead] + 1
		} else {
			// odd
			distribution[CounterTail] = distribution[CounterTail] + 1
		}
	}

	// output the result of coin tossing
	fmt.Println()
	if distribution[CounterHead] >= distribution[CounterTail] {
		fmt.Println("result of coin tossing:", CounterHead)
		coinSide = Head
	} else {
		fmt.Println("result of coin tossing:", CounterTail)
		coinSide = Tail
	}

	// output the result of compare
	if userChoice == coinSide {
		fmt.Println("you win")
	} else {
		fmt.Println("you loos")
	}

	fmt.Println()
	fmt.Println("check for honesty (%):")
	fmt.Printf("%s: %.0f (%.2f)\n", CounterHead, distribution[CounterHead], distribution[CounterHead] * 100 / DatasetLen)
	fmt.Printf("%s: %.0f (%.2f)\n", CounterTail, distribution[CounterTail], distribution[CounterTail] * 100 / DatasetLen)

	fmt.Println()
}
