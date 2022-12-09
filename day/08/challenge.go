package day_08

import "github.com/johncalvinroberts/advent-of-code-2022/utils"

type ForestTree struct {
	height  int
	visible bool
}

// find how many trees are visible from outside the grid
func Part1(input string) int {
	rows := utils.StrToSlice(input, "\n")
	// key = 2 int array, representing x and y
	allTreesVisibility := make(map[[2]int]*ForestTree)
	for y, row := range rows {
		for x, tree := range row {
			var visible bool
			key := [2]int{x, y}
			height := utils.StrToInt(string(tree), 0)
			if x == 0 || y == 0 || x == len(rows) || y == len(row) {
				visible = true
			}

			entry := &ForestTree{height: height, visible: visible}
			allTreesVisibility[key] = entry
		}
	}
	return 0
}

func Part2(input string) {}
