package day_08

import (
	"fmt"

	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

type ForestTree struct {
	height  int
	visible bool
}

// used for converting rune to int
const zero rune = '0'

// find how many trees are visible from outside the grid
func Part1(input string) int {
	rows := utils.StrToSlice(input, "\n")
	var (
		width  = len(rows[0])
		height = len(rows)
	)

	// build set of trees
	// key = string, representing row and column
	allTreesVisibility := make(map[string]*ForestTree)
	for y, row := range rows {
		for x, tree := range row {
			var (
				nodeHeight = int(tree - zero)
				visible    = false
			)
			if x == 0 || y == 0 || x == width-1 || y == height-1 {
				visible = true
			}
			entry := &ForestTree{height: nodeHeight, visible: visible}
			allTreesVisibility[getKey(x, y)] = entry
		}
	}
	// horizontal
	for r := 1; r < height-1; r++ {
		leftMax := int(rows[r][0])
		rightMax := int(rows[r][width-1])
		for c := 1; c < width-1; c++ {
			fmt.Printf("left: %v, right: %v\n", getKey(r, c), getKey(r, width-1-c))
			leftCurrent := int(rows[r][c])
			rightCurrent := int(rows[r][width-1-c])
			if leftCurrent > leftMax {
				leftMax = leftCurrent
				allTreesVisibility[getKey(r, c)].visible = true
			}
			if rightCurrent > rightMax {
				rightMax = rightCurrent
				allTreesVisibility[getKey(r, width-1-c)].visible = true
			}
		}
	}
	// vertical
	for c := 1; c < width-1; c++ {
		topMax := int(rows[0][c])
		bottomMax := int(rows[height-1][c])
		for r := 1; r < height-1; r++ {
			topCurrent := int(rows[r][c])
			bottomCurrent := int(rows[height-1-r][c])
			if topCurrent > topMax {
				topMax = topCurrent
				allTreesVisibility[getKey(r, c)].visible = true
			}
			if bottomCurrent > bottomMax {
				bottomCurrent = bottomMax
				allTreesVisibility[getKey(height-1-r, c)].visible = true
			}
		}
	}
	var count int
	for _, v := range allTreesVisibility {
		if v.visible {
			// fmt.Printf("k: %v, v: %v\n", k, v.height)
			count++
		}
	}
	return count
}

// guessed: 7581, too high

func Part2(input string) {}

func getKey(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}
