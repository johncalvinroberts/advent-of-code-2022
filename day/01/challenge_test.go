package day_01

import (
	"testing"

	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

func TestPart1(t *testing.T) {
	fixture := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	res := Part1(fixture)
	utils.Assert(res, 24000, t)
}

func TestPart2(t *testing.T) {
	fixture := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	res := Part2(fixture)
	utils.Assert(res, 45000, t)
}
