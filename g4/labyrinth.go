package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	LabyrinthEndRow    = 1
	LabyrinthEndColumn = 1

	LabyrinthRows    = 12
	LabyrinthColumns = 20

	LabyrinthWall  = "0"
	LabyrinthPLace = "."
	LabyrinthHorse = "*"
)

type LabyrinthRowType = []string
type LabyrinthType = []LabyrinthRowType

type LabyrinthPointType = []int
type LabyrinthRouteType = []LabyrinthPointType

// algorithm building labyrinth
func labyrinthBuild(labyrinth *LabyrinthType, route *LabyrinthRouteType, start LabyrinthPointType, height, width int) bool {
	// make base matrix
	*labyrinth = labyrinthSeedWalls(height, width)

	// add to labyrinth real route to moving
	*route = labyrinthRoute(start, height, width)
	if *route == nil {
		return false
	}

	// add real route into base matrix
	labyrinthSeedRealRoute(labyrinth, *route)

	// add to labyrinth fake points to confused a player
	labyrinthSeedGaps(labyrinth)

	// put horse on the start
	putHorseToPoint(labyrinth, labyrinthCurrentPosition(height, width))

	return true
}

// convert real point in labyrinth point
func labyrinthCurrentPosition(height, width int) LabyrinthPointType {
	return LabyrinthPointType{height - 1, width - 1}
}

// init labyrinth by walls
func labyrinthSeedWalls(height, width int) (labyrinth LabyrinthType) {
	for row := 0; row < height; row++ {
		var (
			line []string
		)

		for column := 0; column < width; column++ {
			line = append(line, LabyrinthWall)
		}

		labyrinth = append(labyrinth, line)
	}

	return labyrinth
}

// init labyrinth by point available for horse
func labyrinthSeedRealRoute(labyrinth *LabyrinthType, points LabyrinthRouteType) {
	for _, point := range points {
		(*labyrinth)[point[0]-1][point[1]-1] = LabyrinthPLace
	}
}

// init labyrinth by fake point available for horse
func labyrinthSeedGaps(labyrinth *LabyrinthType) {
	// algorithm itself well generate fakes
}

// add horse mark in labyrinth
func putHorseToPoint(labyrinth *LabyrinthType, point LabyrinthPointType) {
	(*labyrinth)[point[0]][point[1]] = LabyrinthHorse
}

// remove horse mark in labyrinth
func removeHorseToPoint(labyrinth *LabyrinthType, point LabyrinthPointType) {
	(*labyrinth)[point[0]][point[1]] = LabyrinthPLace
}

// draws labyrinth for user
func labyrinthOutput(labyrinth LabyrinthType) {
	var (
		width, height int
	)

	height = len(labyrinth)
	width = len(labyrinth[0])

	fmt.Println()

	for row := 0; row < height; row++ {
		for column := 0; column < width; column++ {
			fmt.Print(labyrinth[row][column])
		}

		fmt.Println()
	}

	fmt.Printf("\n\n")
}

// detect near available points
func findNearAvailablePoints(node LabyrinthPointType, height, width int, selected *LabyrinthRouteType) (routes []LabyrinthPointType) {
	var (
		offset [][]int
	)

	offset = append(offset, []int{-1, -2}, []int{1, -2}, []int{-1, 2}, []int{1, 2})
	offset = append(offset, []int{-2, -1}, []int{-2, 1}, []int{2, -1}, []int{2, 1})

	for _, set := range offset {
		var (
			column, row int
			known       bool
		)

		row = node[0] + set[0]
		if (row > height) || (row < 1) {
			continue
		}

		column = node[1] + set[1]
		if (column > width) || (column < 1) {
			continue
		}

		// skip if point already in route
		if selected != nil {
			for _, point := range *selected {
				if (point[0] == row) && (point[1] == column) {
					known = true
					break
				}
			}
		}

		if known {
			continue
		}

		routes = append(routes, []int{row, column})
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(routes), func(i, j int) { routes[i], routes[j] = routes[j], routes[i] })

	return routes
}

// building step by step real route
func buildRoute(node LabyrinthPointType, height, width int, route, selected *LabyrinthRouteType) bool {
	var (
		nearAvailablePoints []LabyrinthPointType
	)

	*selected = append(*selected, node)

	// find near points
	nearAvailablePoints = findNearAvailablePoints(node, height, width, selected)

	// check among near points end point
	for _, point := range nearAvailablePoints {
		if (point[0] == LabyrinthEndRow) && (point[1] == LabyrinthEndColumn) {
			*route = append(*route, point)

			return true
		}
	}

	// go deeper
	for _, point := range nearAvailablePoints {
		if buildRoute(point, height, width, route, selected) {
			*route = append(*route, point)

			return true
		}
	}

	return false
}

// step by step build real routs and get only one
func labyrinthRoute(start LabyrinthPointType, height, width int) (route LabyrinthRouteType) {
	var (
		routes              []LabyrinthRouteType
		nearAvailablePoints []LabyrinthPointType
	)

	// find near points
	nearAvailablePoints = findNearAvailablePoints(start, height, width, nil)

	// check among near points end point
	for _, point := range nearAvailablePoints {
		if (point[0] == LabyrinthEndRow) && (point[1] == LabyrinthEndColumn) {
			route = append(route, point)

			return route
		}
	}

	// go deeper
	for _, point := range nearAvailablePoints {
		var (
			selected LabyrinthRouteType
		)

		route = make(LabyrinthRouteType, 0)
		selected = append(selected, start)

		if buildRoute(point, height, width, &route, &selected) {
			route = append(route, point, start)
			routes = append(routes, route)
		}
	}

	if len(routes) == 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(routes), func(i, j int) { routes[i], routes[j] = routes[j], routes[i] })

	return routes[0]
}

// check if point is wall
func isWall(labyrinth LabyrinthType, pos LabyrinthPointType) bool {
	if labyrinth[pos[0]-1][pos[1]-1] != LabyrinthWall {
		return false
	}

	return true
}

// check out of labyrinth range
func isStepOut(labyrinth LabyrinthType, pos LabyrinthPointType) bool {
	var (
		width, height int
	)

	height = len(labyrinth)
	width = len(labyrinth[0])

	if (pos[TurnFirst] < LabyrinthEndRow) || (pos[TurnFirst] > height) {
		return true
	}

	if (pos[TurnSecond] < LabyrinthEndColumn) || (pos[TurnSecond] > width) {
		return true
	}

	return false
}
