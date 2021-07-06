/*
Игра 2.  Стратегия для одной игры в кости.

Каждый игрок в свою очередь хода бросает кость столько раз, сколько хочет. Если он не выбрасывает единицу, то он
записывает за этот ход сумму выпавших за бросания этого хода очков. Если же он выбрасывает единицу, то он не записывает
ничего (и его ход кончается с выбрасыванием единицы). Выигравшим считается тот, кто первым наберет (или превысит) 100
очков.

Программа реализует бросание кости. На своем ходе она честна и не мошенничает. На вашем ходе она бросает кость и
сообщает, что выпало, а вы требуете следующего бросания, если вы хотите играть дальше.
 */

package main

import "fmt"

const (
	No                = 0

	RollsNobody       = 0
	RollsMan          = 1
	RollsPC           = 2
	RollsPCLastChance = 3

	MaxPCRolls        = 3
	SpurtPCRolls      = 5
	
	MaxScorePoints    = 100
)

type ScoreTableType struct {
	Man int
	PC  int
}

func main() {
	var (
		whoRollingNext, diceEdge, earned, doNextTry, countPCTurns int
		scoreTable                                                ScoreTableType
	)

	fmt.Println("let's start the game!")

	// man always start first
	whoRollingNext = RollsMan

	// game process
	for whoRollingNext > 0 {
		switch whoRollingNext {
		case RollsMan:
			fmt.Println()

			// ask man about next turn
			fmt.Print("do you want to continue or next move of PC (0 – no, 1 – yes)? ")
			_, _ = fmt.Scanf("%d", &doNextTry)

			if doNextTry == No {
				// man choose stop
				scoreTable.Man = scoreTable.Man + earned

				// PC will be do next turn
				whoRollingNext = RollsPC
				earned = 0

				fmt.Println()
			} else {
				// make a new try
				diceEdge = mostOftenDropEdge()

				if diceEdge == 1 {
					// skip all earned because roll 1 and PC will be do next turn
					whoRollingNext = RollsPC
					earned = 0

					fmt.Printf("dice edge is %d, your score is %d \n\n", diceEdge, scoreTable.Man)
					fmt.Println("next turn of PC")
					fmt.Println()
				} else {
					earned = earned + diceEdge

					if (scoreTable.Man + earned) >= MaxScorePoints {
						// man has went on target
						scoreTable.Man = scoreTable.Man + earned

						fmt.Printf("dice edge is %d, you earn last turn %d, your score is %d \n\n", diceEdge, earned, scoreTable.Man)

						// give last chance for PC
						whoRollingNext = RollsPCLastChance
						earned = 0
					} else {
						// continue if man want it
						fmt.Printf("dice edge is %d, you earn last turn %d, your score is %d \n", diceEdge, earned, scoreTable.Man+earned)
					}
				}
			}

		case RollsPC:
			// get current PC turn
			countPCTurns++

			if countPCTurns > MaxPCRolls {
				// for PC not more possible try to increase his scoring, he store earned points
				scoreTable.PC = scoreTable.PC + earned

				fmt.Printf("earn last turn %d, PC score is %d vs %d \n", earned, scoreTable.PC, scoreTable.Man)
				fmt.Println("next turn of man")

				// passing the move
				whoRollingNext = RollsMan

				countPCTurns = 0
				earned = 0
			} else {
				diceEdge = mostOftenDropEdge()

				if diceEdge == 1 {
					// skip all earned because roll 1 and man will be do next turn
					whoRollingNext = RollsMan

					countPCTurns = 0
					earned = 0

					fmt.Printf("dice edge is %d, PC score is %d vs %d \n", diceEdge, scoreTable.PC, scoreTable.Man)
					fmt.Println("next turn of man")
				} else {
					earned = earned + diceEdge

					if (scoreTable.PC + earned) >= MaxScorePoints {
						// man has went on target
						scoreTable.PC = scoreTable.PC + earned

						fmt.Printf("earn last turn %d, PC score is %d vs %d \n", earned, scoreTable.PC, scoreTable.Man)
						fmt.Println("stop game")

						// stop game because PC roll second then last in player pair
						whoRollingNext = RollsNobody
					}
				}
			}

		case RollsPCLastChance:
			// get current PC turn
			countPCTurns++

			if countPCTurns > SpurtPCRolls {
				// for PC not more possible try to increase his scoring, he store earned points
				scoreTable.PC = scoreTable.PC + earned

				fmt.Printf("earn last turn %d, PC score is %d vs %d \n", earned, scoreTable.PC, scoreTable.Man)
				fmt.Println("stop game")

				// stop game
				whoRollingNext = RollsNobody
			} else {
				diceEdge = mostOftenDropEdge()

				if diceEdge == 1 {
					// stop game
					whoRollingNext = RollsNobody

					fmt.Printf("dice edge is %d, PC score is %d vs %d \n", diceEdge, scoreTable.PC, scoreTable.Man)
					fmt.Println("stop game")
				} else {
					earned = earned + diceEdge

					if (scoreTable.PC + earned) > scoreTable.Man {
						// PC give more points than man
						scoreTable.PC = scoreTable.PC + earned

						fmt.Printf("earn last turn %d, PC score is %d vs %d \n", earned, scoreTable.PC, scoreTable.Man)
						fmt.Println("stop game")

						// stop game
						whoRollingNext = RollsNobody
					}
				}
			}

		default:
			fmt.Println()
			fmt.Println("stop game: unknown player")

			// stopping game if it is not understand whose turn
			whoRollingNext = RollsNobody
		}
	}

	// output final score
	fmt.Println()
	fmt.Println("your score:", scoreTable.Man, "PC score:", scoreTable.PC)
}
