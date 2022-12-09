package day_08

import (
	"testing"

	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

var fixture = `30373
25512
65332
33549
35390`

func TestPart1(t *testing.T) {
	result := Part1(fixture)
	utils.Assert(result, 21, t)
}

func TestPart2(t *testing.T) {}
