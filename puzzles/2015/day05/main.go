package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	lines := strings.Split(getInput(), "\n")
	fmt.Println("Part 1 solution:", part1(lines))
	fmt.Println("Part 2 solution:", part2(lines))
}

func getInput() string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(input))
}

func part1(lines []string) int {
	niceStrings := 0
	for _, line := range lines {
		if isNiceString(line) {
			niceStrings += 1
		}
	}
	return niceStrings
}

func part2(lines []string) int {
	nicerStrings := 0
	for _, line := range lines {
		if isNicerString(line) {
			nicerStrings += 1
		}
	}
	return nicerStrings
}

func isNiceString(line string) bool {
	for _, mustNotContain := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(line, mustNotContain) {
			return false
		}
	}

	vowelCount := 0
	for _, vowel := range "aeiou" {
		vowelCount += strings.Count(line, string(vowel))
	}
	if vowelCount < 3 {
		return false
	}

	repeats := 0
	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] {
			repeats++
		}
	}
	if repeats == 0 {
		return false
	}

	return true
}

func isNicerString(line string) bool {
	repeats := 0
	for i := 1; i < len(line)-1; i++ {
		if line[i-1] == line[i+1] {
			repeats++
		}
	}
	if repeats == 0 {
		return false
	}

	pairs := 0
	for i := 0; i < len(line)-2; i++ {
		if strings.Index(line[i+2:], line[i:i+2]) > -1 {
			pairs++
		}
	}
	if pairs == 0 {
		return false
	}

	return true
}
