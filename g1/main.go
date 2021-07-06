/*
Игра 1.  Фальшивые кости.

Компьютер будет мошенником. Он играет одной-единственной костью и бросает ее столько раз, сколько вы требуете. Он дает
вам число выпаданий каждой из граней.Вы сообщаете ему, верите ли вы, что кости поддельные, и если да, то какая грань
выпадает чаще других. Компьютер отвечает вам, выиграли вы или проиграли, и случайным образом оценивает ваш выигрыш.

Нужно решать две задачи: компьютер должен выбрать — подделывать кости или не подделывать, и если он их подделывает, то
он должен решить, какая грань будет встречаться чаще остальных.
*/

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	DefaultRollings = 50
)

func main() {
	var (
		countRollings, diceEdge        int
		cheatEdge, cheatRoll           int
		isCheaterPC, possibleCheatEdge int
		distributionEdges              map[int]int
		cheatModeIsOn                  bool
	)

	distributionEdges = make(map[int]int)

	// randomly determine we will cheat or not
	cheatModeIsOn = (rand.New(rand.NewSource(time.Now().UnixNano())).Intn(99) + 1) > 50

	// when we are cheater
	if cheatModeIsOn {
		// so determine cheating edge
		cheatEdge = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(5) + 1

		// set randomly cheat level
		cheatRoll = (rand.New(rand.NewSource(time.Now().UnixNano())).Intn(2) + 1) * 2
	}

	fmt.Print("set count rollings: ")
	_, _ = fmt.Scanf("%d", &countRollings)

	if countRollings == 0 {
		countRollings = DefaultRollings
	}

	// collection dice's statistics
	for i := 0; i < countRollings; i++ {
		// do not roll dice when cheat mode is on and time to cheat
		if cheatModeIsOn && (math.Mod(float64(i+1), float64(cheatRoll)) == 0) {
			diceEdge = cheatEdge
		} else {
			diceEdge = mostOftenDropEdge()
		}

		distributionEdges[diceEdge] = distributionEdges[diceEdge] + 1
	}

	// output the result roll of dices
	fmt.Println("")
	for key, value := range distributionEdges {
		fmt.Printf("dice edges [%d] rolls %d (%.2f%%)\n", key, value, float64(value)*100/float64(countRollings))
	}

	// ask user about our cheating
	fmt.Print("\nis PC a cheater? ")
	_, _ = fmt.Scanf("%d", &isCheaterPC)

	// user think that PC is cheater
	if isCheaterPC == 1 {
		// ask user which edge is cheating
		fmt.Print("so, and who is cheat edge? ")
		_, _ = fmt.Scanf("%d", &possibleCheatEdge)

		// check user guess cheat edge or not
		if cheatModeIsOn && possibleCheatEdge == cheatEdge {
			fmt.Println("yes. i am a cheater. take my ", rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000), "$")
		} else {
			fmt.Println("no. i am honest PC. maybe... i take you money")
		}
	} else {
		fmt.Println("off course i am honest PC. maybe... i take you money")
	}
}
