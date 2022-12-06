package day_06

import (
	"testing"

	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

func TestPart1(t *testing.T) {
	var (
		fixtures = map[string]int{
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    7,
			"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
			"nppdvjthqldpwncqszvftbrmjlhg":      6,
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
		}
	)
	for k, v := range fixtures {
		result := Part1(k)
		utils.Assert(result, v, t)
	}
}

func TestPart2(t *testing.T) {
	var (
		fixtures = map[string]int{
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    19,
			"bvwbjplbgvbhsrlpgdmjqwftvncz":      23,
			"nppdvjthqldpwncqszvftbrmjlhg":      23,
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  26,
		}
	)
	for k, v := range fixtures {
		result := Part2(k)
		utils.Assert(result, v, t)
	}
}
