package day_08

import (
	"fmt"

	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

type Grid [][]int

type Forest struct {
	grid          Grid
	height, width int
}

// used for converting rune to int
const zero rune = '0'

// gets visible trees from all 4 directions
func (g *Forest) countVisible() int {
	// set of any visible trees with key "rowcol"
	visible := utils.Set[string]{}

	check_height := func(r, c, max int) int {
		tree_height := g.grid[r][c]
		if tree_height > max {
			// tree is > max therefore visible
			visible.Add(fmt.Sprint(r, c))
			return tree_height
		}
		return max
	}

	// start now on inner trees HORIZONTAL
	for r := 1; r < g.height-1; r++ {
		// keep track of max so far
		// L->R (break if max == 9)
		max := g.grid[r][0]
		for c := 1; c < g.width-1 && max < 9; c++ {
			max = check_height(r, c, max)
		}

		// R->L (break if max == 9)
		max = g.grid[r][g.width-1]
		for c := g.width - 2; c > 0 && max < 9; c-- {
			max = check_height(r, c, max)
		}
	}

	// VERTICAL
	for c := 1; c < g.width-1; c++ {
		// keep track of max so far
		// T->B (break if max == 9)
		max := g.grid[0][c]
		for r := 1; r < g.height-1 && max < 9; r++ {
			max = check_height(r, c, max)
		}

		// B->T (break if max == 9)
		max = g.grid[g.height-1][c]
		for r := g.height - 2; r > 0 && max < 9; r-- {
			max = check_height(r, c, max)
		}
	}

	// include edges (top, bottom, left, right)
	edges := g.height*2 + (g.width-2)*2

	return len(visible) + edges
}

// find how many trees are visible from outside the grid
func Part1(input string) int {
	var (
		rows   = utils.StrToSlice(input, "\n")
		width  = len(rows[0])
		height = len(rows)
		grid   = make(Grid, height)
		trees  = &Forest{
			height: height,
			width:  width,
		}
	)

	for r, line := range rows {
		grid[r] = make([]int, width)
		for c, char := range line {
			grid[r][c] = int(char - zero)
		}
	}
	trees.grid = grid
	return trees.countVisible()
}

// find highest scenic score possible for any tree in the forest
// scenic score = multiplying the viewing distance of all four directions
func Part2(input string) int {
	var (
		rows   = utils.StrToSlice(input, "\n")
		width  = len(rows[0])
		height = len(rows)
		grid   = make(Grid, height)
		trees  = &Forest{
			height: height,
			width:  width,
		}
	)
	for r, line := range rows {
		grid[r] = make([]int, width)
		for c, char := range line {
			grid[r][c] = int(char - zero)
		}
	}
	trees.grid = grid

	calculateScenicScore := func(r, c int) int {
		var (
			left        = 0
			right       = 0
			top         = 0
			bottom      = 0
			rightFound  bool
			leftFound   bool
			topFound    bool
			bottomFound bool
			treeHeight  = trees.grid[r][c]
		)

		for i := 1; i < width; i++ {
			nextLeftIdx := c - i
			nextRightIdx := c + i
			if nextLeftIdx >= 0 && !leftFound {
				left++
				nextLeft := trees.grid[r][nextLeftIdx]
				if nextLeft >= treeHeight {
					leftFound = true
				}
			}
			if nextRightIdx < width && !rightFound {
				right++
				nextRight := trees.grid[r][nextRightIdx]
				if nextRight >= treeHeight {
					rightFound = true
				}
			}
		}
		for i := 1; i < height; i++ {
			nextTopIdx := r - i
			nextBottomIdx := r + i
			if nextTopIdx >= 0 && !topFound {
				top++
				nextTop := trees.grid[nextTopIdx][c]
				if nextTop >= treeHeight {
					topFound = true
				}
			}
			if nextBottomIdx < height && !bottomFound {
				bottom++
				nextBottom := trees.grid[nextBottomIdx][c]
				if nextBottom >= treeHeight {
					bottomFound = true
				}
			}
		}
		fmt.Printf("left: %d right: %d, top: %d, bottom %d\n", left, right, top, bottom)
		return left * right * top * bottom
	}
	var maxScenicScore int
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			scenicScore := calculateScenicScore(r, c)
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}
	return maxScenicScore
}
