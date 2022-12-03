package day_03

import (
	"testing"

	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

var fixture string = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestPart1(t *testing.T) {
	result := Part1(fixture)
	utils.Assert(result, 157, t)
}

func TestPart2(t *testing.T) {
	result := Part2(fixture)
	utils.Assert(result, 70, t)
}
