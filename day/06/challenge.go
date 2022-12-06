package day_06

// a marker is 4 unique characters in a chunk
// How many characters need to be processed before the first start-of-packet marker is detected?
func Part1(input string) int {
	var result int
	for i := 3; i < len(input); i++ {
		// counts for current 4 char chunk
		counts := make(map[string]int)
		counts[string(input[i])]++
		counts[string(input[i-1])]++
		counts[string(input[i-2])]++
		counts[string(input[i-3])]++
		if checkIsMatch(counts) {
			result = i + 1
			break
		}
	}
	return result
}

func Part2(input string) int {
	var result int
	for i := 13; i < len(input); i++ {
		// counts for current 14 char chunk
		counts := make(map[string]int)
		counts[string(input[i])]++
		for j := 13; j > 0; j-- {
			counts[string(input[i-j])]++
		}
		if checkIsMatch(counts) {
			result = i + 1
			break
		}
	}
	return result
}

func checkIsMatch(counts map[string]int) bool {
	var allOne bool = true
	for _, v := range counts {
		if v > 1 {
			allOne = false
			break
		}
	}
	return allOne
}
