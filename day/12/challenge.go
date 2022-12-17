package day_12

import (
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
	lines := utils.StrToSlice(input, "\n")
	goal, start, _ := parseInput(lines)
	result := bfs(start, goal, []Coordinate{start}, lines)
	return result
}

func Part2(input string) int {
	lines := utils.StrToSlice(input, "\n")
	goal, start, lowPoints := parseInput(lines)
	result := bfs(start, goal, lowPoints, lines)
	return result
}

// breadth first search
func bfs(start, goal Coordinate, starts []Coordinate, grid []string) int {
	var (
		shortestPath = 10_000_000
		adjacents    = []Coordinate{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	)
	for _, startingPoint := range starts {
		var (
			currentShortestPath = 10_000_000
			distances           = make(map[Coordinate]int)
			queue               = utils.Queue[Coordinate]{}
		)
		queue.Enqueue(startingPoint)

		for !queue.IsEmpty() {
			var (
				cur = queue.Dequeue().Value
			)
			if cur.Equals(&goal) {
				// path is shorter than all previously traversed paths
				if distances[cur] < currentShortestPath {
					currentShortestPath = distances[cur]
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
		if currentShortestPath < shortestPath {
			shortestPath = currentShortestPath
		}
	}
	return shortestPath
}

func parseInput(lines []string) (Coordinate, Coordinate, []Coordinate) {
	var (
		start, goal Coordinate
		lowPoints   []Coordinate
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
			if char == 'a' {
				lowPoints = append(lowPoints, Coordinate{x, y})
			}
		}
	}
	return goal, start, lowPoints
}
