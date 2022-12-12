package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/alecthomas/kong"
	day01 "github.com/johncalvinroberts/advent-of-code-2022/day/01"
	day02 "github.com/johncalvinroberts/advent-of-code-2022/day/02"
	day03 "github.com/johncalvinroberts/advent-of-code-2022/day/03"
	day04 "github.com/johncalvinroberts/advent-of-code-2022/day/04"
	day05 "github.com/johncalvinroberts/advent-of-code-2022/day/05"
	day06 "github.com/johncalvinroberts/advent-of-code-2022/day/06"
	day07 "github.com/johncalvinroberts/advent-of-code-2022/day/07"
	day08 "github.com/johncalvinroberts/advent-of-code-2022/day/08"
	day09 "github.com/johncalvinroberts/advent-of-code-2022/day/09"
	day10 "github.com/johncalvinroberts/advent-of-code-2022/day/10"
	"github.com/johncalvinroberts/advent-of-code-2022/utils"
)

var Advent struct {
	Scaffold ScaffoldCmd `cmd:"" help:"Scaffold an AOC challenge."`
	Run      RunCmd      `cmd:"" help:"Run an AOC challenge."`
}

type RunCmd struct {
	Day int `help:"Which day do you want to run?"`
}

func (cmd *RunCmd) Run() error {
	if cmd.Day != 0 {
		fmt.Printf("🌲 Running day %02d\n", cmd.Day)
	} else {
		fmt.Println("🎅 Running all challenges!!")
	}
	switch cmd.Day {
	case 0:
		// TODO: run all challenges
	case 1:
		fmt.Printf("part 1: %d\n", day01.Part1(utils.ReadDayFile(cmd.Day)))
		fmt.Printf("part 2: %d\n", day01.Part2(utils.ReadDayFile(cmd.Day)))
	case 2:
		fmt.Printf("part 1: %d\n", day02.Part1(utils.ReadDayFile(cmd.Day)))
		fmt.Printf("part 2: %d\n", day02.Part2(utils.ReadDayFile(cmd.Day)))
	case 3:
		fmt.Printf("part 1: %d\n", day03.Part1(utils.ReadDayFile(cmd.Day)))
		fmt.Printf("part 2: %d\n", day03.Part2(utils.ReadDayFile(cmd.Day)))
	case 4:
		fmt.Printf("part 1: %d\n", day04.Part1(utils.ReadDayFile(cmd.Day)))
		fmt.Printf("part 2: %d\n", day04.Part2(utils.ReadDayFile(cmd.Day)))
	case 5:
		fmt.Printf("part 1: %s\n", day05.Part1(utils.ReadDayFile(cmd.Day)))
		fmt.Printf("part 2: %s\n", day05.Part2(utils.ReadDayFile(cmd.Day)))
	case 6:
		fmt.Printf("part 1: %d\n", day06.Part1(utils.ReadDayFile(cmd.Day)))
		fmt.Printf("part 2: %d\n", day06.Part2(utils.ReadDayFile(cmd.Day)))
	case 7:
		fmt.Printf("part 1: %d\n", day07.Part1(utils.ReadDayFile(cmd.Day)))
		fmt.Printf("part 2: %d\n", day07.Part2(utils.ReadDayFile(cmd.Day)))
	case 8:
		fmt.Printf("part 1: %d\n", day08.Part1(utils.ReadDayFile(cmd.Day)))
		fmt.Printf("part 2: %d\n", day08.Part2(utils.ReadDayFile(cmd.Day)))
	case 9:
		fmt.Printf("part 1: %d\n", day09.Part1(utils.ReadDayFile(cmd.Day)))
		fmt.Printf("part 2: %d\n", day09.Part2(utils.ReadDayFile(cmd.Day)))
	case 10:
		fmt.Printf("part 1: %d\n", day10.Part1(utils.ReadDayFile(cmd.Day)))
		fmt.Printf("part 2: \n%s\n", day10.Part2(utils.ReadDayFile(cmd.Day)))
	}
	return nil
}

type ScaffoldCmd struct {
	Day int `help:"Which day do you want to pull down?"`
}

func (cmd *ScaffoldCmd) Run() error {
	fmt.Printf("🧝‍♀️ scaffolding day %d\n", cmd.Day)
	cookie := utils.ReadFile("cookie")
	url := fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", cmd.Day)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("cookie", fmt.Sprintf("session=%s", cookie))
	res, err := client.Do(req)
	utils.PanicOnErr(err)
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	utils.PanicOnErr(err)
	var (
		dirname      = fmt.Sprintf("day/%02d", cmd.Day)
		inputFile    = fmt.Sprintf("%s/input.txt", dirname)
		chlgFile     = fmt.Sprintf("%s/challenge.go", dirname)
		testFile     = fmt.Sprintf("%s/challenge_test.go", dirname)
		chlgScaffold = fmt.Sprintf(`package day_%02d
			
func Part1(input string) {}
	
func Part2(input string) {}
	`, cmd.Day)
		testScaffold = fmt.Sprintf(`package day_%02d
import "testing"
func TestPart1(t *testing.T) {
	t.Error("Not Implemented")
}
	
func TestPart2(t *testing.T) {}
	`, cmd.Day)
	)

	utils.MakeDir(dirname)
	utils.WriteFile(inputFile, data)
	utils.WriteFile(chlgFile, []byte(chlgScaffold))
	utils.WriteFile(testFile, []byte(testScaffold))
	fmt.Println("🌲🌲好了！Done! ❄️❄️❄️❄️")
	return nil
}

func main() {
	ctx := kong.Parse(&Advent,
		kong.Description("🎄圣🎅诞🦌树🎄 ADVENT of CODE 2022 ⿅🌲⿅"),
	)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
