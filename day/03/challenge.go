package day_03

import (
	"strings"

	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// find the item that appears in both compartments of each rucksack
// sum the priority of those items
func Part1(input string) int {
	var (
		priorities int
		rucksacks  = utils.StrToSlice(input, "\n")
	)
	for _, rs := range rucksacks {
		one, two := rs[0:len(rs)/2], rs[len(rs)/2:]
	findItemInBothCompartments:
		for _, item := range one {
			if strings.Contains(two, string(item)) {
				p := strings.Index(alphabet, string(item)) + 1
				priorities += p
				break findItemInBothCompartments
			}
		}
	}
	return priorities
}

func Part2(input string) int {
	var (
		priorities int
		rucksacks  = utils.StrToSlice(input, "\n")
	)

	for i := 0; i < len(rucksacks); i += 3 {
		var (
			j                = i + 1
			k                = i + 2
			firstElfRucksack = rucksacks[i]
		)
	findItemCommonAmongElves:
		for _, item := range firstElfRucksack {
			isInNextElfRucksack := strings.Contains(rucksacks[j], string(item))
			isInNextNextElfRucksack := strings.Contains(rucksacks[k], string(item))
			if isInNextElfRucksack && isInNextNextElfRucksack {
				p := strings.Index(alphabet, string(item)) + 1
				priorities += p
				break findItemCommonAmongElves
			}
		}
	}
	return priorities
}
