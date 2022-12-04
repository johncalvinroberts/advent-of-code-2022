package day_04

import (
	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

func Part1(input string) int {
	pairs := utils.StrToSlice(input, "\n")
	var count int
	for _, s := range pairs {
		pair := utils.StrToSlice(s, ",")
		first, second := utils.StrSliceToIntSlice(
			utils.StrToSlice(pair[0], "-")),
			utils.StrSliceToIntSlice(utils.StrToSlice(pair[1], "-"))
		if (first[0] >= second[0] && first[1] <= second[1]) ||
			(second[0] >= first[0] && second[1] <= first[1]) {
			count++
		}
	}
	return count
}

func Part2(input string) int {
	pairs := utils.StrToSlice(input, "\n")
	var count int
	for _, s := range pairs {
		pair := utils.StrToSlice(s, ",")
		first, second := utils.StrSliceToIntSlice(
			utils.StrToSlice(pair[0], "-")),
			utils.StrSliceToIntSlice(utils.StrToSlice(pair[1], "-"))
		// second is the area with the higher starting point
		if first[0] > second[0] {
			second, first = first, second
		}
		// end of first section overlaps with start of second section
		if first[1] >= second[0] {
			count++
		}
	}
	return count
}
