package day_12

import (
	"fmt"
	"strings"

	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

var (
	altitudes = "abcdefghijklmnopqrstuvwxyz"
)

type Coordinate struct {
	x, y int
}

func (c *Coordinate) Equals(oc *Coordinate) bool {
	return c.x == oc.x && c.y == oc.y
}

func (c *Coordinate) Add(oc *Coordinate) Coordinate {
	return Coordinate{x: c.x + oc.x, y: c.y + oc.y}
}

func (c *Coordinate) GetHeight(grid []string) int {
	char := string(grid[c.y][c.x])
	height := strings.Index(altitudes, char)
	return height
}

// Start at S, get to E
// S is altitude [a], e is at altitude [z].
// you can only move 1 altitude unit at a time. e.g., can go c to d, or f to e.
// Get the minimum number of steps required to get to the goal (E)
func Part1(input string) int {
	var (
		lines       = utils.StrToSlice(input, "\n")
		start, goal Coordinate
	)

	// start at S
	for y, line := range lines {
		for x, char := range line {
			if char == 'S' {
				lines[y] = strings.Replace(lines[y], "S", "a", 1)
				start.x, start.y = x, y
			}
			if char == 'E' {
				lines[y] = strings.Replace(lines[y], "E", "z", 1)
				goal.x, goal.y = x, y
			}
		}
	}
	result := bfs(start, goal, lines)
	return result
}

func Part2(input string) {}

// breadth first search
func bfs(start, goal Coordinate, grid []string) int {
	var (
		shortestPathLength = 10_000_000
		distances          = make(map[Coordinate]int)
		adjacents          = []Coordinate{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
		queue              = utils.Queue[Coordinate]{}
	)

	queue.Enqueue(start)

	for !queue.IsEmpty() {
		var (
			cur = queue.Dequeue().Value
		)
		if cur.y == 20 {
			fmt.Println(cur.x, cur.y)

		}
		if cur.Equals(&goal) {
			// path is shorter than all previously traversed paths
			if distances[cur] < shortestPathLength {
				shortestPathLength = distances[cur]
			}
		}

		for _, adj := range adjacents {
			next := cur.Add(&adj)
			if next.x > len(grid[0])-1 || next.y > len(grid)-1 || next.x < 0 || next.y < 0 {
				// off the grid, continue
				continue
			}
			// if next is less than or equal to 1 diff from current coordinate
			// and it's not been added to the map, add it
			var (
				_, exists      = distances[next]
				nextCharHeight = next.GetHeight(grid)
				currentHeight  = cur.GetHeight(grid)
			)
			if nextCharHeight-currentHeight <= 1 && !exists {
				queue.Enqueue(next)
				distances[next] = distances[cur] + 1
			}
		}

	}
	return shortestPathLength
}
