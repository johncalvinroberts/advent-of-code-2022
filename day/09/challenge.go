package day_09

import (
	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

type Coordinate struct {
	x, y int
}

// count how many positions the tail has occupied
// after every move, the tail needs to be adjacent to the head
// tail can only move one position at a time
// if head makes a move and is still adjacent to tail, tail doesn't move
func Part1(input string) int {
	return simulateRopeMotion(input, 2)
}

func Part2(input string) int {
	return simulateRopeMotion(input, 10)
}

func simulateRopeMotion(input string, length int) int {
	var (
		moves = utils.StrToSlice(input, "\n")
		// visited tail positions
		tailPath = &utils.Set[Coordinate]{}
		// state of all knots
		knots = make([]Coordinate, length)
	)
	// add tail starting path
	tailPath.Add(Coordinate{0, 0})
	for _, x := range moves {
		move := utils.StrToSlice(x, " ")
		dir, n := move[0], utils.StrToInt(move[1], 0)
		// first move head
		for n > 0 {
			n--
			switch dir {
			case "L":
				knots[0].x--
			case "R":
				knots[0].x++
			case "U":
				knots[0].y++
			case "D":
				knots[0].y--
			}
			// then, move tail
			for i := 0; i < length-1; i++ {
				head, tail := knots[i], &knots[i+1]
				var (
					xDiff = head.x - tail.x
					yDiff = head.y - tail.y
				)
				if utils.Absolute(xDiff) > 1 || utils.Absolute(yDiff) > 1 {
					tail.x += utils.Sign(xDiff)
					tail.y += utils.Sign(yDiff)
				}
			}
			tailPath.Add(knots[length-1])
		}
	}
	return len(*tailPath)
}
