package day_09

import (
	"fmt"
	"strings"

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
	var (
		moves    = utils.StrToSlice(input, "\n")
		head     = &Coordinate{0, 0}
		tail     = &Coordinate{0, 0}
		tailPath = &utils.Set[Coordinate]{}
	)

	moveTail := func() {
		var (
			xDiff = head.x - tail.x
			yDiff = head.y - tail.y
		)
		if utils.Absolute(xDiff) > 1 || utils.Absolute(yDiff) > 1 {
			tail.x += utils.Sign(xDiff)
			tail.y += utils.Sign(yDiff)
		}
		tailPath.Add(*tail)
	}
	tailPath.Add(*tail)
	for _, x := range moves {
		move := utils.StrToSlice(x, " ")
		dir, n := move[0], utils.StrToInt(move[1], 0)
		// first move head
		for n > 0 {
			n--
			switch dir {
			case "L":
				head.x--
			case "R":
				head.x++
			case "U":
				head.y++
			case "D":
				head.y--
			}
			moveTail()
		}
	}
	return len(*tailPath)
}

func Part2(input string) {}

func PrintTailPath(p *utils.Set[Coordinate]) {
	format := [][]string{
		{".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "."},
	}
	for k := range *p {
		row := format[4-k.y]
		row[k.x] = "#"
	}
	var joined string
	for _, line := range format {
		row := strings.Join(line, "")
		joined = strings.Join([]string{joined, row}, "\n")
	}
	fmt.Println(joined)
}
