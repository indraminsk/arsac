/*
Головоломка 1.  Периодическая последовательность.

Построить последовательность целых чисел. Написать программу, определяющую период этой последовательности. Ограничение:
вы не имеете права запоминать в таблице последовательные значения элементов последовательности (вы не имеете права
запоминать их и в любой другой форме).
 */

package main

import (
	"fmt"
)

const (
	DefaultA = 3
	DefaultB = 5
	DefaultC = 7

	DefaultSequenceLength = 10000
)

func main() {
	var (
		a, b, c, pPrev, pNext, sequenceLength, period, item int
		cumulativeDelta                                     map[int]int
	)

	cumulativeDelta = make(map[int]int)

	fmt.Print("set any three digits from 1 to 9 (default 3, 5, 7): ")
	_, _ = fmt.Scanf("%d %d %d", &a, &b, &c)

	// set default values for a, b, c, if it's needing
	if a == 0 {
		a = DefaultA
	}

	if b == 0 {
		b = DefaultB
	}

	if c == 0 {
		c = DefaultC
	}

	fmt.Printf("set sequence length (at least 1000, default – %d): ", DefaultSequenceLength)
	_, _ = fmt.Scanf("%d", &sequenceLength)

	// set default values for sequence length, if it's needing
	if sequenceLength == 0 {
		sequenceLength = DefaultSequenceLength
	}

	// calculate first element in sequence
	pPrev = int(a * 100 + b * 10 + c)

	// calculate other elements of sequence
	func () {
		for i := 0; i < sequenceLength; i++ {
			// calculate new element
			pNext = int(1000 * ale(float64(pPrev)/1000))

			// increase delta for each element in sequence
			for step := 0; step <= item; step++ {
				// increase delta
				cumulativeDelta[step] = cumulativeDelta[step] + (pNext - pPrev)

				// check when delta equals zero this signal that periodical sequence is found
				if cumulativeDelta[step] == 0 {
					// calculate period, plus 1 to do not lost end item
					period = item - step + 1
					fmt.Printf("period of sequance is %d (start item %d, end item %d)\n", period, step, item)

					return
				}
			}

			// increment step
			pPrev = pNext
			item++
		}
	}()

	fmt.Println()
}
