package day_02

import (
	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

var (
	me        = map[string]int{"X": 1, "Y": 2, "Z": 3}
	opponent  = map[string]int{"A": 1, "B": 2, "C": 3}
	winmap    = map[string]string{"A": "Y", "B": "Z", "C": "X"}
	lossmap   = map[string]string{"A": "Z", "B": "X", "C": "Y"}
	drawmap   = map[string]string{"A": "X", "B": "Y", "C": "Z"}
	winAward  = 6
	drawAward = 3
)

func Part1(input string) int {
	rounds := utils.StrToSlice(input, "\n")
	var score int
	for _, x := range rounds {
		r := utils.StrToSlice(x, " ")
		score += calculateScore(r[0], r[1])
	}
	return score
}

func Part2(input string) int {
	rounds := utils.StrToSlice(input, "\n")
	var score int
	for _, x := range rounds {
		r := utils.StrToSlice(x, " ")
		theirs, mine := r[0], r[1]
		switch mine {
		case "X":
			// lose
			score += calculateScore(theirs, lossmap[theirs])
		case "Y":
			// draw
			score += calculateScore(theirs, drawmap[theirs])
		case "Z":
			// win
			score += calculateScore(theirs, winmap[theirs])
		}
	}
	return score
}

func calculateScore(theirs string, mine string) int {
	o := opponent[theirs]
	m := me[mine]
	tmp := m
	// draw
	if o == m {
		tmp += drawAward
	}
	// I rock'd opponent's scissors
	if m == me["X"] && o == opponent["C"] {
		tmp += winAward
	}
	// I beat opponent
	if (m == me["Y"] && o == opponent["A"]) || (m == me["Z"] && o == opponent["B"]) {
		tmp += winAward
	}
	return tmp
}
