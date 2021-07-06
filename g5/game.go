package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"strings"
)

const (
	MoveLeft  = "A"
	MoveRight = "D"
	MoveUp    = "W"
	MoveDown  = "S"
)

func loop(castle CastleType, start CastlePointType, tries, height, width int) bool {
	var (
		err error

		current CastlePointType
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
			ok bool

			pressed rune
			key     keyboard.Key
			char    string
			next    CastlePointType
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
		ok = moves(char)
		if !ok {
			continue
		}

		// show user pressed key
		fmt.Print(char)

		// calculate new position
		next = getNewPos(char, current)

		// check move across boundaries
		if isStepOut(castle, next) {
			fmt.Println(" – move to out castle forbidden")
			fmt.Print("your turns: ")

			continue
		}

		// check new position is wall or not
		if isWall(castle, next) {
			fmt.Println(" - move on wall forbidden")
			fmt.Print("your turns: ")

			continue
		}

		removeHeroFromPoint(&castle, casteCurrentPosition(current))
		putHeroToPoint(&castle, casteCurrentPosition(next))

		// apply to hte matrix
		if !castleSeedObstacles(&castle, height, width, generateObstacles(height, width), casteCurrentPosition(next)) {
			fmt.Println("... and obstacle fall on you")
			break
		}

		castleOutput(castle)

		// store new position as current
		current = next

		// update move counter
		tries = tries - 1

		// when we use all available moves but not reach top left corner stop game
		if tries == 0 {
			fmt.Println("you use all your tries")

			break
		}

		displayStatus(tries)

		// maybe we reached goal
		if (current[0] == CastleEndRow) && (current[1] == CastleEndColumn) {
			fmt.Println("you win!")
			break
		}
	}

	return true
}

func displayStatus(tries int) {
	fmt.Println("available turns:")
	fmt.Println("A – left", "D – right", "W – up", "S - down")
	fmt.Println()
	fmt.Println("you have", tries, "tries to reach left top corner")
	fmt.Print("Press ESC to quit", "\n\n")
	fmt.Print("your turns: ")
}

// check typing letter
func moves(char string) bool {
	switch char {
	case MoveUp, MoveDown, MoveLeft, MoveRight:
		return true
	default:
		return false
	}
}

// calculate new position
func getNewPos(char string, current CastlePointType) CastlePointType {
	switch char {
	case MoveUp:
		return CastlePointType{current[0] - 1, current[1]}
	case MoveDown:
		return CastlePointType{current[0] + 1, current[1]}
	case MoveLeft:
		return CastlePointType{current[0], current[1] - 1}
	case MoveRight:
		return CastlePointType{current[0], current[1] + 1}
	default:
		return current
	}
}
