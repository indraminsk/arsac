/*
Игра 3.  Покер — М — С.

Тасуем карточную колоду. Разыгрывается некоторая сумма. Берем верхнюю карту из колоды и требуем от игрока, чтобы он
угадал, является ли следующая карта младшей или старшей по отношению к только что взятой. Учитывается только число
очков, а не масть карты. Валет всегда больше девяти, король больше валета, туз больше всех. Если игрок угадал правильно,
то сумма в игре удваивается. Если он не угадывает, он теряет все, В конце некоторого фиксированного числа
конов игрок (например пяти), если он всегда оказывался прав, присваивает полученную сумму.

Составьте программу, которая позволит вам быть игроком, а компьютер пусть будет всем остальным (за исключением того,
что вы называете и сумму игры).

Можно реализовать накопление выигрыша и если в очередной раз проигрышь перекрывает доступные средства, то игра
заканчивается досрочно
 */

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	DefaultBid = 1000
)

type DeckType map[float64]int

// using random number as position card in deck
func generatePlaceInDeck() float64 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
}

// shuffle deck
func shuffleDeck() (deck DeckType, places []float64) {
	deck = make(DeckType)
	deck = make(DeckType)

	// bind card with random number
	for card := 2; card <= 14; card++ {
		for suit := 1; suit <= 4; suit++ {
			deck[generatePlaceInDeck()] = card
		}
	}

	// fetch random numbers to next sort
	for place := range deck {
		places = append(places, place)
	}

	// sort random number
	sort.Float64s(places)

	return deck, places
}

// get human card's name
func getCardName(cardNumber int) (cardName string) {
	switch cardNumber {
	case 11:
		cardName = "J"
	case 12:
		cardName = "D"
	case 13:
		cardName = "K"
	case 14:
		cardName = "A"
	default:
		cardName = fmt.Sprintf("%d", cardNumber)
	}

	return cardName
}

// output deck in growth order
func showDeck(deck DeckType, places []float64) {
	fmt.Println("deck in growth order:")
	for _, place := range places {
		fmt.Print(getCardName(deck[place]), " ")
	}
	fmt.Printf("\n\n")
}

func main() {
	var (
		bid int
		deck            DeckType
		places          []float64
	)

	deck = make(DeckType)

	fmt.Printf("how much money (default – %d) will you play in the game? ", DefaultBid)
	_, _ = fmt.Scanf("%d", &bid)
	fmt.Println()

	if bid == 0 {
		bid = DefaultBid
	}

	// shuffle deck
	deck, places = shuffleDeck()

	for i := 1; i < 6; i++ {
		var (
			greater int
		)

		fmt.Println("your card – ", getCardName(deck[places[i - 1]]))

		fmt.Print("is the next card greater than your (yes – 1, no – 0)? ")
		_, _ = fmt.Scanf("%d", &greater)

		fmt.Println("deck card – ", getCardName(deck[places[i]]))

		// check user choice
		if greater > 0 {
			if deck[places[i - 1]] > deck[places[i]] {
				// user wait that next card greater yours but no
				fmt.Printf("\nyou lost your money (:\n\n")
				showDeck(deck, places)

				return
			}
		} else {
			// user wait that next card less yours but no
			if deck[places[i - 1]] < deck[places[i]] {
				fmt.Printf("\nyou lost your money (:\n\n")
				showDeck(deck, places)

				return
			}
		}

		// increase bid
		bid = bid * 2
		fmt.Println()
	}

	fmt.Printf("\nyou win! get your money %d \n\n", bid)
	showDeck(deck, places)
}
