package day_12

import (
	"testing"

	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

var fixture string = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func TestPart1(t *testing.T) {
	result := Part1(fixture)
	utils.Assert(result, 31, t)
}

func TestPart2(t *testing.T) {}
