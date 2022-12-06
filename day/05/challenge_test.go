package day_05

import (
	"testing"

	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

var fixture string = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestPart1(t *testing.T) {
	result := Part1(fixture)
	utils.Assert(result, "CMZ", t)
}

func TestPart2(t *testing.T) {
	result := Part2(fixture)
	utils.Assert(result, "MCD", t)
}
