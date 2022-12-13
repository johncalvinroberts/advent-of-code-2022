package day_11

import (
	"fmt"
	"strings"

	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

type MonkeyState struct {
	// list of items, value is worry level
	Items        *utils.Queue[int]
	Operation    string
	TestOperand  int
	PassMonkey   int
	FailMonkey   int
	InspectCount int
}

// monkey business = product of number of items inspected by two most active monkeys
// return the level of monkey business over twenty rounds
func Part1(input string) int {
	return watchMonkeysFiddleWithBackpackContents(20, 3, input)
}

// same as part 1, but no divide by three after each round, and do 10000 rounds
func Part2(input string) int {
	return watchMonkeysFiddleWithBackpackContents(10000, 1, input)
}

func watchMonkeysFiddleWithBackpackContents(rounds, worryDivisor int, input string) int {
	var (
		monkeys                    = parseMonkeyOperation(input)
		mostInspectiveMonkey       int
		secondMostInspectiveMonkey int
	)
	for i := 0; i < rounds; i++ {
		for _, currentMonkey := range monkeys {
			for !currentMonkey.Items.IsEmpty() {
				v := currentMonkey.Items.Dequeue()
				currentMonkey.InspectCount++
				nextWorryLevel, ok := executeOperation(v.Value, worryDivisor, currentMonkey)
				if !ok {
					monkeys[currentMonkey.FailMonkey].Items.Enqueue(nextWorryLevel)
				} else {
					monkeys[currentMonkey.PassMonkey].Items.Enqueue(nextWorryLevel)
				}
			}
		}
		// if i == 0 || i == 19 || i == 999 {
		// 	fmt.Println("")

		// 	fmt.Printf("i: %d\n", i)
		// 	for _, m := range monkeys {
		// 		fmt.Println(m.InspectCount)
		// 	}
		// 	fmt.Println("")
		// }
	}

	for _, m := range monkeys {
		if m.InspectCount > mostInspectiveMonkey {
			secondMostInspectiveMonkey, mostInspectiveMonkey = mostInspectiveMonkey, m.InspectCount
			continue
		}
		if m.InspectCount > secondMostInspectiveMonkey {
			secondMostInspectiveMonkey = m.InspectCount
		}
	}
	return mostInspectiveMonkey * secondMostInspectiveMonkey
}

func parseMonkeyOperation(raw string) []*MonkeyState {
	var (
		rawMonkeys = utils.StrToSlice(raw, "\n\n")
		monkeys    = make([]*MonkeyState, len(rawMonkeys))
	)
	for i, monk := range rawMonkeys {
		var (
			split       = strings.Split(monk, "\n")
			worryLevels = &utils.Queue[int]{}
			operation   string
			test        int
			passMonkey  int
			failMonkey  int
		)

		for _, x := range utils.StrToSlice(utils.StrToSlice(split[1], ": ")[1], ", ") {
			worryLevels.Enqueue(utils.MustStrToInt(x))
		}
		operation = utils.StrToSlice(strings.TrimSpace(split[2]), ": ")[1]
		fmt.Sscanf(strings.TrimSpace(split[3]), "Test: divisible by %d", &test)
		fmt.Sscanf(strings.TrimSpace(split[4]), "If true: throw to monkey %d", &passMonkey)
		fmt.Sscanf(strings.TrimSpace(split[5]), "If false: throw to monkey %d", &failMonkey)
		statefulMonk := &MonkeyState{
			Items:       worryLevels,
			Operation:   operation,
			PassMonkey:  passMonkey,
			FailMonkey:  failMonkey,
			TestOperand: test,
		}
		monkeys[i] = statefulMonk
	}
	return monkeys
}

func executeOperation(w, worryDivisor int, m *MonkeyState) (int, bool) {
	var (
		leftStr        string
		rightStr       string
		right          int
		left           int
		operator       string
		nextWorryLevel int
		ok             bool
	)
	parsedOp := utils.StrToSlice(m.Operation, " ")
	leftStr = parsedOp[2]
	operator = parsedOp[3]
	rightStr = parsedOp[4]

	if leftStr == "old" {
		left = w
	} else {
		left = utils.MustStrToInt(leftStr)
	}
	if rightStr == "old" {
		right = w
	} else {
		right = utils.MustStrToInt(rightStr)
	}
	switch operator {
	case "*":
		nextWorryLevel = left * right
	case "-":
		nextWorryLevel = left - right
	case "+":
		nextWorryLevel = left + right
	case "/":
		nextWorryLevel = left / right
	}
	// this happens every time a monkey gets bored with an item
	nextWorryLevel = nextWorryLevel / worryDivisor
	ok = nextWorryLevel%m.TestOperand == 0
	return nextWorryLevel, ok
}