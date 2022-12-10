package day_09

import (
	"testing"

	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

var fixture = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

// var fixture = `R 4`

func TestPart1(t *testing.T) {
	result := Part1(fixture)
	utils.Assert(result, 13, t)
}

func TestPart2(t *testing.T) {}
