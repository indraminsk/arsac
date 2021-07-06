/*
Головоломка 3.  Вращающееся число.

Найти такое число, оканчивающееся на 5, что, умножая его на 5, мы получим новое число, полученное из предыдущего
вычеркиванием цифры 5 на конце и приписыванием ее в начале.

разпространить алгоритм на все числа в диапазоне от 2 до 9
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// revert result string slice
func humanizeGoal(in []string) string {
	for i := len(in)/2 - 1; i >= 0; i-- {
		var (
			mirrored int
		)

		mirrored = len(in) - 1 - i
		in[i], in[mirrored] = in[mirrored], in[i]
	}
	return strings.Join(in, "")
}

func main() {
	var (
		err error

		previous int
	)

	//sayHello("try find out loop number for each digit from range 2..9")
	sayHello()

	// main loop for numbers from 2 to 9
	for multiplayer := 2; multiplayer < 10; multiplayer++ {
		var (
			lastDigit, mindDigit int
			goal                 []string
		)

		// at start last digit always equals multiplayer
		lastDigit = multiplayer
		// init previous by last digit, whe previous will equals zero it will means stop
		//previous = 0

		// store first digit
		goal = append(goal, strconv.Itoa(lastDigit+mindDigit))

		// slave loop, step by step find digits
		//for previous != 0 {
		for !((previous == 0) && (lastDigit == 1)) {
			var (
				buffer  string
				product []rune
				dump int
			)

			// init buffer by product last digit and multiplayer
			buffer = strconv.Itoa(lastDigit * multiplayer)
			// convert string to rune for comfortable work with digits
			product = []rune(buffer)

			// change process previous
			previous = lastDigit

			// get next digit for store
			lastDigit, err = strconv.Atoi(string(product[len(buffer)-1]))
			if err != nil {
				fmt.Println("something went wrong (last digit):", err, "\nbuffer", buffer)
				return
			}

			// corrected last digit by mind digit
			lastDigit = lastDigit + mindDigit

			// processed case when product less than 10 to correct init mind digit
			if len(product) == 2 {
				mindDigit, err = strconv.Atoi(string(product[len(buffer)-2]))
				if err != nil {
					fmt.Println("something went wrong (mind digit):", err, "\nbuffer", buffer)
					return
				}
			} else {
				mindDigit = 0
			}

			// processed case whe we have last digit more then 9
			if lastDigit > 9 {
				// init buffer by product last digit and multiplayer
				buffer = strconv.Itoa(lastDigit)
				// convert string to rune for comfortable work with digits
				product = []rune(buffer)

				// get next digit for store again
				lastDigit, err = strconv.Atoi(string(product[len(buffer)-1]))
				if err != nil {
					fmt.Println("something went wrong (last digit):", err, "\nbuffer", buffer)
					return
				}

				dump, err = strconv.Atoi(string(product[len(buffer)-2]))
				mindDigit = mindDigit + dump
				if err != nil {
					fmt.Println("something went wrong (mind digit):", err, "\nbuffer", buffer)
					return
				}
			}

			// store found digit
			goal = append(goal, strconv.Itoa(lastDigit))
		}

		fmt.Printf("%d: %s\n", multiplayer, humanizeGoal(goal))
	}

	sayBye()

	return
}
