package day_01

import (
	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

// find the elf carrying the most calories + return how many calories that is
func Part1(input string) int {
	rawElfLoads := utils.StrToSlice(input, "\n\n")
	var max int
	for _, raw := range rawElfLoads {
		calories := utils.StrSliceToIntSlice(utils.StrToSlice(raw, "\n"))
		sum := utils.SumSlice(calories)
		if sum > max {
			max = sum
		}
	}
	return max
}

// find the sum of the top 3 elves carrying the most calories
func Part2(input string) int {
	rawElfLoads := utils.StrToSlice(input, "\n\n")
	// 3 item array of top elf calorie bearers, in ascending order
	// so index 0 = 3rd
	var topElves [3]int
	for _, raw := range rawElfLoads {
		calories := utils.StrSliceToIntSlice(utils.StrToSlice(raw, "\n"))
		sum := utils.SumSlice(calories)
	top:
		for i, x := range topElves {
			// if the elf is less than the third top elf, break
			if sum < x && i < 1 {
				break top
			}
			if sum > x && i == 2 {
				tmp1, tmp0 := topElves[2], topElves[1]
				topElves[i], topElves[1], topElves[0] = sum, tmp1, tmp0
				break top
			}
			if sum > x && sum < topElves[2] && sum > topElves[0] {
				tmp := topElves[1]
				topElves[1] = sum
				topElves[0] = tmp
				break top
			}

			topElves[0] = sum
		}
	}
	return utils.SumSlice(topElves[:])
}
