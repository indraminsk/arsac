/*
Головоломка 2.  Последовательности «орлов» и «решек».

Осуществим n  выбрасываний «орла» и «решки» с большим n  (например 10000). Сколько раз встретится в ней данная
комбинация из m  следующих друг за другом выбрасываний (например, 10 раз «орел» или чередование из 10 выбрасываний
«орла» и «решки», начиная с «орла»).
 */

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	DefaultDatasetLength  = 10000
	DefaultSequenceLength = 10
	CatchHeads            = 1
	CatchTails            = 2
	CatchBoth             = 3
	ObverseHead           = -1
	ObverseTail           = 1
	ObverseDecimeter      = 0.5
)

func setupSequence(datasetLength, sequenceType, sequenceLength *int) {
	// init dataset and sequence parameters
	fmt.Printf("set dataset length (default – %d): ", DefaultDatasetLength)
	_, _ = fmt.Scanf("%d", datasetLength)

	if *datasetLength == 0 {
		*datasetLength = DefaultDatasetLength
	}

	fmt.Printf("set sequence type (heads – %d, tails – %d, both (default, head is leading) – %d): ", CatchHeads, CatchTails, CatchBoth)
	_, _ = fmt.Scanf("%d", sequenceType)

	if *sequenceType == 0 {
		*sequenceType = CatchBoth
	}

	fmt.Printf("set sequence length (default – %d): ", DefaultSequenceLength)
	_, _ = fmt.Scanf("%d", sequenceLength)

	if *sequenceLength == 0 {
		*sequenceLength = DefaultSequenceLength
	}
}

func main() {
	var (
		start                                                                     time.Time
		datasetLength, sequenceType, sequenceLength, reachedSequence, reachedPair int
		obverse, checkSum                                                         float64
	)

	start = time.Now()

	// user setup game
	setupSequence(&datasetLength, &sequenceType, &sequenceLength)

	for i := 0; i < datasetLength; i++ {
		// determine which obverse
		if ObverseDecimeter > rand.New(rand.NewSource(time.Now().UnixNano())).Float64() {
			obverse = ObverseHead
			//fmt.Print("h–")
		} else {
			obverse = ObverseTail
			//fmt.Print("t–")
		}

		// check sequence after last coin toss
		switch sequenceType {
		case CatchHeads:
			// skip when last obverse is tail
			if obverse == ObverseTail {
				// reset check sum because sequence is broken by toss tail
				checkSum = 0
				break
			}

			// update check sum
			checkSum = checkSum + obverse

			// check reaching goal sequence length
			if int(math.Abs(checkSum)) == sequenceLength {
				// increase sequence counter
				reachedSequence++
				// reset check sum to start new series
				checkSum = 0
			}

		case CatchTails:
			// skip when last obverse is tail
			if obverse == ObverseHead {
				// resetting check sum because sequence is broken by toss head
				checkSum = 0
				break
			}

			// update check sum
			checkSum = checkSum + obverse

			// check reaching goal sequence length
			if int(math.Abs(checkSum)) == sequenceLength {
				// increase sequence counter
				reachedSequence++
				// resetting check sum to start new series
				checkSum = 0
			}

		case CatchBoth:
			// skip because two next each other tails do not interesting us
			// checkSum zero says that before this we found pair and need head again
			if (checkSum == 0) && (obverse == ObverseTail) {
				// resetting check sum and counter of found pairs
				checkSum = 0
				reachedPair = 0

				break
			}

			// skip because two next each other heads also do not interesting us, we have interest to pair head–tail
			if (checkSum < 0) && (obverse == ObverseHead) {
				// resetting check sum and counter of found pairs
				checkSum = 0
				reachedPair = 0

				break
			}

			// update check sum
			checkSum = checkSum + obverse

			// pair is found because head plus tail equals zero
			if checkSum == 0 {
				// new pair is found then increase counter of pairs
				reachedPair++

				// check reaching goal pair sequence length
				if reachedPair == sequenceLength {
					// increase sequence counter
					reachedSequence++
					// resetting check sum to find new pair and to start new series
					checkSum = 0
					reachedPair = 0
				}
			}
		}
	}

	fmt.Printf("\nreach goal – %d [time of work – %s]\n\n", reachedSequence, time.Since(start))
}
