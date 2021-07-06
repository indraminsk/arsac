/*
Игра 5.  Спящая красавица.

Вы представляете замок тем же самым прямоугольным пространством, где точки обозначают свободные места, нули —
препятствия, а звездочка сообщает местоположение Тони, В начале игры Тони находится в правом нижнем углу, а выход
находится в левом верхнем углу. Препятствий вначале крайне мало. После каждого хода Тони компьютер случайным образом
формирует новые препятствия и размещает их на игровом поле. Если Тони оказывается на месте одного из них, то он
раздавлен и для него все кончено…
*/

package main

import "fmt"

const (
	AvailableCountMoves = 148
)

func main() {
	var (
		height, width int
		start         CastlePointType
		castle        CastleType
	)

	sayHello("let's escape begin", "'0' is obstacle", "'.' is available places to move")

	fmt.Println("set (separate by space) castle size (height, width) or use default values")
	_, _ = fmt.Scanf("%d %d", &height, &width)

	// init by default values if needed
	if height == 0 {
		height = CastleRows
	}

	if width == 0 {
		width = CastleColumns
	}

	// init start position
	start = CastlePointType{height, width}

	// build castle
	if !castleBuild(&castle, start, height, width) {
		sayBye("bye!")
		return
	}

	// rendering castle
	castleOutput(castle)

	// let's game start
	loop(castle, start, AvailableCountMoves, height, width)

	sayBye("bye! see you later")
}
