package day_10

import (
	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

/*
input is a series of cpu instructions for a single register cpu, which starts at 1
signal strength = the cycle number multiplied by the value of the X register
find the sum of the signal strength at the 20th, 60th, 100th, 140th, 180th, and 220th cycles

noop = 1 cycle
addx = two cycles
*/
func Part1(input string) int {
	var (
		instructions = utils.StrToSlice(input, "\n")
		register     = 1
		audit        = []int{register}
		debugCycles  = []int{20, 60, 100, 140, 180, 220}
		sum          int
	)
	for _, cmd := range instructions {
		audit = append(audit, register)
		if cmd == "noop" {
			continue
		}
		parsed := utils.StrToSlice(cmd, " ")
		register += utils.StrToInt(parsed[1], 0)
		audit = append(audit, register)
	}

	for _, x := range debugCycles {
		sum += (x * audit[x-1])
	}

	return sum
}

func Part2(input string) {}
