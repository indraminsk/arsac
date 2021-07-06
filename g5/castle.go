package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

const (
	SeedSpeed = 400

	CastleEndRow    = 1
	CastleEndColumn = 1

	CastleRows    = 36
	CastleColumns = 80

	CastleObstacle = "0"
	CastlePLace    = "."
	CastleHero     = "*"

	ObstacleDirectionD = 0
	ObstacleDirectionR = 1
)

type CastlePointType = []int

type CastleRowType = []string
type CastleType = []CastleRowType

type CastleObstacleType = struct {
	row    int
	column int
	axis   int
}
type CastleObstaclesType = []CastleObstacleType

// build castle
func castleBuild(castle *CastleType, start CastlePointType, height, width int) bool {
	// make base matrix
	castleInit(castle, height, width)

	// apply to hte matrix
	if !castleSeedObstacles(castle, height, width, generateObstacles(height, width), casteCurrentPosition(start)) {
		fmt.Println("... and obstacle fall on you")
		return false
	}

	// put hero on the start
	putHeroToPoint(castle, casteCurrentPosition(start))

	return true
}

// init castle
func castleInit(castle *CastleType, height, width int) {
	for i := 0; i < height; i++ {
		var (
			line CastleRowType
		)

		for j := 0; j < width; j++ {
			line = append(line, CastlePLace)
		}

		*castle = append(*castle, line)
	}
}

// generate randomly obstacle axis
func generateRandomAxis() int {
	var (
		axises []int
	)

	axises = []int{ObstacleDirectionD, ObstacleDirectionR}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(axises), func(i, j int) { axises[i], axises[j] = axises[j], axises[i] })

	return axises[0]
}

// generate randomly row or column
func generateRandomRC(dimension int) int {
	var (
		values []int
	)

	for i := 0; i < dimension; i++ {
		values = append(values, i+1)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(values), func(i, j int) { values[i], values[j] = values[j], values[i] })

	return values[0]
}

// generate obstacles
func generateObstacles(height, width int) (obstacles CastleObstaclesType) {
	var (
		places int
	)

	places = int((height * width) / SeedSpeed)

	for i := 0; i < places; i++ {
		obstacles = append(obstacles, CastleObstacleType{
			row:    generateRandomRC(height),
			column: generateRandomRC(width),
			axis:   generateRandomAxis()})
	}

	return obstacles
}

// add to castle obstacles
func castleSeedObstacles(castle *CastleType, height, width int, obstacles CastleObstaclesType, hero CastlePointType) bool {
	for _, obstacle := range obstacles {
		// check out of castle range
		if (obstacle.column == width) || (obstacle.row == height) {
			continue
		}

		// check out of castle range
		if (obstacle.row - 2) >= CastleEndRow {
			// don't allow near (one step to up ) points where place obstacle
			if (*castle)[obstacle.row-2][obstacle.column-1] == CastleObstacle {
				continue
			}
		}

		// check out of castle range
		if (obstacle.column - 2) >= CastleEndColumn {
			// don't allow near (one step to left) points where place obstacle
			if (*castle)[obstacle.row-1][obstacle.column-2] == CastleObstacle {
				continue
			}
		}

		// check exist obstacle for first obstacle point
		if (*castle)[obstacle.row-1][obstacle.column-1] == CastleObstacle {
			continue
		}

		// place if it possible obstacle
		if obstacle.axis == ObstacleDirectionR {
			// check exist obstacle near second obstacle point
			if isOffsetNotValid(*castle, obstacle, height, width) {
				continue
			}

			// check if hero in first point of obstacle
			if isThisPointWithHero(hero, CastlePointType{obstacle.row - 1, obstacle.column - 1}) {
				return false
			}

			// check if hero in second point of obstacle
			if isThisPointWithHero(hero, CastlePointType{obstacle.row - 1, obstacle.column}) {
				return false
			}

			// generate two points together (it's the strategy)
			(*castle)[obstacle.row-1][obstacle.column-1] = CastleObstacle
			(*castle)[obstacle.row-1][obstacle.column] = CastleObstacle
		} else {
			// check exist obstacle near second obstacle point
			if isOffsetNotValid(*castle, obstacle, height, width) {
				continue
			}

			// check if hero in first point of obstacle
			if isThisPointWithHero(hero, CastlePointType{obstacle.row - 1, obstacle.column - 1}) {
				return false
			}

			// check if hero in second point of obstacle
			if isThisPointWithHero(hero, CastlePointType{obstacle.row, obstacle.column - 1}) {
				return false
			}

			// generate two points together (it's the strategy)
			(*castle)[obstacle.row-1][obstacle.column-1] = CastleObstacle
			(*castle)[obstacle.row][obstacle.column-1] = CastleObstacle
		}
	}

	return true
}

// check exist obstacle
func isOffsetNotValid(castle CastleType, obstacle CastleObstacleType, height, width int) bool {
	if obstacle.axis == ObstacleDirectionR {
		if (obstacle.column + 1) != width {
			// just right one position after obstacle
			if castle[obstacle.row-1][obstacle.column+1] == CastleObstacle {
				return true
			}

			// just right and up one position after obstacle
			if (obstacle.row - 2) >= CastleEndRow {
				if castle[obstacle.row-2][obstacle.column+1] == CastleObstacle {
					return true
				}
			}

			// just right and down one position after obstacle
			if castle[obstacle.row][obstacle.column+1] == CastleObstacle {
				return true
			}
		}
	} else {
		// check out of castle range for second obstacle point
		if (obstacle.row + 1) != height {
			// just down one position after obstacle
			if castle[obstacle.row+1][obstacle.column-1] == CastleObstacle {
				return true
			}

			// just down and left one position after obstacle
			if (obstacle.column - 2) >= CastleEndRow {
				if castle[obstacle.row+1][obstacle.column-2] == CastleObstacle {
					return true
				}
			}

			// just down and right one position after obstacle
			if castle[obstacle.row+1][obstacle.column] == CastleObstacle {
				return true
			}
		}
	}

	return false
}

// compare current position of hero and assume position of obstacle
func isThisPointWithHero(hero, point CastlePointType) bool {
	if !reflect.DeepEqual(hero, point) {
		return false
	}

	return true
}

// convert row and column to castle position
func casteCurrentPosition(point CastlePointType) CastlePointType {
	return CastlePointType{point[0] - 1, point[1] - 1}
}

// remove hero mark in castle
func removeHeroFromPoint(castle *CastleType, point CastlePointType) {
	(*castle)[point[0]][point[1]] = CastlePLace
}

// put hero to certain point
func putHeroToPoint(castle *CastleType, point CastlePointType) {
	(*castle)[point[ObstacleDirectionD]][point[ObstacleDirectionR]] = CastleHero
}

// output castle
func castleOutput(castle CastleType) {
	fmt.Println()

	for _, row := range castle {
		for _, column := range row {
			fmt.Print(column)
		}

		fmt.Println()
	}
}

// check out of castle range
func isStepOut(castle CastleType, pos CastlePointType) bool {
	var (
		width, height int
	)

	height = len(castle)
	width = len(castle[0])

	if (pos[0] < CastleEndRow) || (pos[0] > height) {
		return true
	}

	if (pos[1] < CastleEndColumn) || (pos[1] > width) {
		return true
	}

	return false
}

// check if point is wall
func isWall(castle CastleType, pos CastlePointType) bool {
	if castle[pos[0]-1][pos[1]-1] != CastleObstacle {
		return false
	}

	return true
}
