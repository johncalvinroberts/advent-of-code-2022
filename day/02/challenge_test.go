package day_02

import (
	"testing"

	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

func TestPart1(t *testing.T) {
	var fixture = `A Y
B X
C Z`
	result := Part1(fixture)
	utils.Assert(result, 15, t)
}

func TestPart2(t *testing.T) {
	var fixture = `A Y
B X
C Z`
	result := Part2(fixture)
	utils.Assert(result, 12, t)
}
