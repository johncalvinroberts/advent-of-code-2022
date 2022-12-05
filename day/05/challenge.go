package day_05

import (
	"math"

	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

func Part1(input string) string {
	stacksAndProc := utils.StrToSlice(input, "\n\n")
	rawStacks, rawCmds :=
		utils.StrToSlice(stacksAndProc[0], "\n"),
		utils.StrToSlice(stacksAndProc[1], "\n")
	width := len(rawStacks[len(rawStacks)-1])
	rawStacks = rawStacks[:len(rawStacks)-1]
	// get the stacks
	// massaged stacks are the columns, ascending up the column
	var massagedStacks [][]string
	for i := 1; i < width; i += 4 {
		for _, v := range rawStacks {
			value := string(v[i])
			if value != " " {
				targetIndex := int(math.Floor(float64(i) / float64(4)))
				if len(massagedStacks) == targetIndex {
					massagedStacks = append(massagedStacks, []string{})
				}
				massagedStacks[targetIndex] = append(massagedStacks[targetIndex], value)
			}
		}
	}
	for _, rawCmd := range rawCmds {
		// cmd is an int slice, each index representing a different part of the move
		// move {cmd[0] (quantity)} from {cmd[1] (source stack)} to {cmd[2] (target stack)}
		cmd := utils.StrSliceToIntSlice(utils.ExtractIntsToStrSlice(rawCmd))
		var (
			qty         = cmd[0]
			sourceIndex = cmd[1] - 1
			targetIndex = cmd[2] - 1
		)

		for i := 0; i < qty; i++ {

			item, nextSourceStack := massagedStacks[sourceIndex][0], massagedStacks[sourceIndex][1:]
			massagedStacks[sourceIndex] = nextSourceStack
			nextTargetStack := append([]string{item}, massagedStacks[targetIndex]...)
			massagedStacks[targetIndex] = nextTargetStack
		}
	}
	var answer string
	for _, x := range massagedStacks {
		answer += x[0]
	}
	return answer
}

func Part2(input string) {}
