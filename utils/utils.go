package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func StrToInt(s string, fallback int) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return fallback
	}
	return v
}

func StrSliceToIntSlice(strs []string) []int {
	ints := make([]int, len(strs))
	for in, v := range strs {
		i := StrToInt(v, 0)
		ints[in] = i
	}
	return ints
}

func StrToSlice(s string, delim string) []string {
	return strings.Split(s, delim)
}

func StrToTuple(s string, delim string) (string, string) {
	x := strings.Split(s, delim)
	return x[0], x[1]
}

func SumSlice(s []int) int {
	result := 0
	for _, v := range s {
		result += v
	}
	return result
}

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadDayFile(day int) string {
	filename := fmt.Sprintf("day/%02d/input.txt", day)
	return ReadFile(filename)
}

func ReadFile(filename string) string {
	file, err := os.Open(filename)
	PanicOnErr(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	contents, err := io.ReadAll(reader)
	PanicOnErr(err)

	return strings.TrimSuffix(string(contents), "\n")
}

func MakeDir(dirname string) {
	err := os.Mkdir(dirname, os.ModePerm)
	PanicOnErr(err)
}

func WriteFile(filename string, data []byte) {
	err := os.WriteFile(filename, data, os.ModePerm)
	PanicOnErr(err)
}

func ExtractIntsToStrSlice(v string) []string {
	re := regexp.MustCompile("[0-9]+")
	strs := re.FindAllString(v, -1)
	return strs
}

func Assert(got int, want int, t *testing.T) {
	if got != want {
		t.Errorf("Expected %d, received %d", want, got)
	} else {
		t.Logf("Got %d, want %d. Good job ✨.", got, want)
	}
}

func Absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}
