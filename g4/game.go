package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"strings"
)

const (
	MoveLeft  = "L"
	MoveRight = "R"
	MoveUp    = "U"
	MoveDown  = "D"

	AxisHeight = 0
	AxisWidth  = 1

	TurnFirst  = 0
	TurnSecond = 1
)

type InputType = []string

type MoveDataType = struct {
	offset []int
	axis   int
}

func moves(char string) (MoveDataType, bool) {
	var (
		ok     bool
		matrix map[string]MoveDataType
	)

	matrix = map[string]MoveDataType{
		MoveLeft: {
			offset: []int{-2, -1},
			axis:   AxisWidth,
		},
		MoveRight: {
			offset: []int{2, 1},
			axis:   AxisWidth,
		},
		MoveUp: {
			offset: []int{-2, -1},
			axis:   AxisHeight,
		},
		MoveDown: {
			offset: []int{2, 1},
			axis:   AxisHeight,
		},
	}

	_, ok = matrix[char]
	if !ok {
		return MoveDataType{}, false
	}

	return matrix[char], true
}

func getNewPosition(input []string, move int, current LabyrinthPointType) (pos int, axis int) {
	var (
		ok bool

		moveData MoveDataType
	)

	moveData, ok = moves(input[move])
	if !ok {
		return 0, -1
	}

	return current[moveData.axis] + moveData.offset[move], moveData.axis
}

func detectWrongDirection(input InputType, char string) bool {
	if len(input) > 0 {
		// skip the same second turn
		if input[TurnFirst] == char {
			return true
		}

		switch input[TurnFirst] {
		case MoveLeft:
			if char == MoveRight {
				return true
			}
		case MoveRight:
			if char == MoveLeft {
				return true
			}
		case MoveUp:
			if char == MoveDown {
				return true
			}
		case MoveDown:
			if char == MoveUp {
				return true
			}
		}
	}

	return false
}

func displayStatus(tries int) {
	fmt.Println("available turns:")
	fmt.Println("L – left", "R – right", "U – up", "D - down")
	fmt.Println()
	fmt.Println("first turn – two steps in any direction")
	fmt.Println("second turn – one step in any direction")
	fmt.Println()
	fmt.Println("you have", tries, "tries to reach left top corner")
	fmt.Print("Press ESC to quit", "\n\n")
	fmt.Print("your turns: ")
}

func loop(labyrinth LabyrinthType, start LabyrinthPointType, tries int) bool {
	var (
		err error

		current LabyrinthPointType
		input   InputType
	)

	current = start

	err = keyboard.Open()
	if err != nil {
		fmt.Println("something went wrong: open")
		return false
	}

	defer func() {
		_ = keyboard.Close()
	}()

	displayStatus(tries)

	for {
		var (
			pressed rune
			key     keyboard.Key

			char string
			ok   bool

			next      LabyrinthPointType
			pos, axis int
		)

		pressed, key, err = keyboard.GetKey()
		if err != nil {
			fmt.Println("something went wrong")
			return false
		}

		// stop game when we press Esc
		if key == keyboard.KeyEsc {
			break
		}

		// formatting pressed key
		char = strings.ToUpper(string(pressed))

		// skip all letters exclude turn letters
		_, ok = moves(char)
		if !ok {
			continue
		}

		// validate second turn
		if detectWrongDirection(input, char) {
			continue
		}

		// when choose only one turn then wait last
		input = append(input, strings.ToUpper(string(pressed)))
		if len(input) < 2 {
			// show user pressed key
			fmt.Print(char)

			continue
		}

		// show user pressed key
		fmt.Print(char)

		// for init use current position only for stable
		next = LabyrinthPointType{current[TurnFirst], current[TurnSecond]}

		// analyze first turn and calculate new position for needed axis
		pos, axis = getNewPosition(input, TurnFirst, current)
		if axis == -1 {
			continue
		}

		next[axis] = pos

		// analyze second turn and calculate new position for needed axis
		pos, axis = getNewPosition(input, TurnSecond, current)
		if axis == -1 {
			break
		}

		next[axis] = pos

		// check move across boundaries
		if isStepOut(labyrinth, next) {
			fmt.Println(" – move to out labyrinth forbidden")
			fmt.Print("your turns: ")

			// clear use input
			input = []string{}

			continue
		}

		// check new position is wall or not
		if isWall(labyrinth, next) {
			fmt.Println(" - move on wall forbidden")
			fmt.Print("your turns: ")

			// clear use input
			input = []string{}

			continue
		}

		removeHorseToPoint(&labyrinth, labyrinthCurrentPosition(current[TurnFirst], current[TurnSecond]))
		putHorseToPoint(&labyrinth, labyrinthCurrentPosition(next[TurnFirst], next[TurnSecond]))
		labyrinthOutput(labyrinth)

		// store new position as current
		current = next

		// clear use input
		input = []string{}

		// update move counter
		tries = tries - 1

		// when we use all available moves but not reach top left corner stop game
		if tries == 0 {
			fmt.Println("you use all your tries")

			break
		}

		displayStatus(tries)

		// maybe we reached goal
		if (current[0] == LabyrinthEndRow) && (current[1] == LabyrinthEndColumn) {
			fmt.Println("you win!")
			break
		}
	}

	return true
}
