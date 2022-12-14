package day_11

import (
	"testing"

	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

var fixture string = `Monkey 0:
	Starting items: 79, 98
	Operation: new = old * 19
	Test: divisible by 23
		If true: throw to monkey 2
		If false: throw to monkey 3

Monkey 1:
	Starting items: 54, 65, 75, 74
	Operation: new = old + 6
	Test: divisible by 19
		If true: throw to monkey 2
		If false: throw to monkey 0

Monkey 2:
	Starting items: 79, 60, 97
	Operation: new = old * old
	Test: divisible by 13
		If true: throw to monkey 1
		If false: throw to monkey 3

Monkey 3:
	Starting items: 74
	Operation: new = old + 3
	Test: divisible by 17
		If true: throw to monkey 0
		If false: throw to monkey 1`

func TestPart1(t *testing.T) {
	result := Part1(fixture)
	utils.Assert(result, 10605, t)
}

func TestPart2(t *testing.T) {
	result := Part2(fixture)
	utils.Assert(result, 2713310158, t)
}
