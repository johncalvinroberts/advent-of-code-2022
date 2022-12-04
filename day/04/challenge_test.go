package day_04

import (
	"testing"

	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

var fixture string = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestPart1(t *testing.T) {
	result := Part1(fixture)
	utils.Assert(result, 2, t)
}

func TestPart2(t *testing.T) {
	result := Part2(fixture)
	utils.Assert(result, 4, t)
}
