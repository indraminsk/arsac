/*
Игра 4.  Лабиринт для шахматного коня.

Пусть дана прямоугольная область, образованная n  строками с p  полями на каждой из них. Занятые места считаются
препятствиями (обозначенными здесь 0), пусть как-то помечены свободные места (здесь — точкой), пусть значок * обозначает
всадника. Конь перемещается, как конь в шахматах: два шага в одном направлении и еще один шаг перпендикулярно
предыдущему направлению. Конь может перемещаться только с одного свободного места на другое, В начальный момент он
находится в правом нижнем углу. Он должен попасть в верхний левый угол (который, таким образом, тоже должен быть
свободным). Число ходов игры ограничено.

Составьте программу для компьютера для создания этого лабиринта и попытки его пройти.

Компьютер сообщает число оставшихся ходов и требует ваших указаний о движении. Ответ дается в виде двух букв: первая из
этих букв дает направление, в котором нужно переместиться на два шага, вторая буква дает перпендикулярное предыдущему
направление, в котором нужно сделать один шаг: Н — для нижней, В — для верхней, П — для правой, Л — для левой сторон.

Компьютер анализирует ответ. Если превышено число ходов или ход встречает препятствие, то игрок проигрывает. Если нет —
звездочка, изображающая коня, перемещается в новое положение, число оставшихся ходов уменьшается на единицу, и игра
продолжается.
*/

package main

import (
	"fmt"
)

func main() {
	var (
		start         LabyrinthPointType
		labyrinth     LabyrinthType
		route         LabyrinthRouteType
		height, width int
	)

	sayHello("let's play begin", "'0' is wall", "'.' is available places to move")

	fmt.Println("set (separate by space) labyrinth size (height, width) or use default values")
	_, _ = fmt.Scanf("%d %d", &height, &width)

	// init by default values if needed
	if height == 0 {
		height = LabyrinthRows
	}

	if width == 0 {
		width = LabyrinthColumns
	}

	// first position into labyrinth
	start = LabyrinthPointType{height, width}

	// build a labyrinth
	if !labyrinthBuild(&labyrinth, &route, start, height, width) {
		sayBye(" bye!")
		return
	}

	// rendering labyrinth
	labyrinthOutput(labyrinth)

	// let's start game
	loop(labyrinth, start, len(route))

	sayBye("bye! see you later")
}
