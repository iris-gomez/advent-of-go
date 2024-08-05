package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	fmt.Printf("Part 1 solution: %d\n", part1(lines))
	fmt.Printf("Part 2 solution: %d\n", part2(lines))
}

func part1(lines []string) int {
	niceStrings := 0
	for _, line := range lines {
		if isNiceString1(line) {
			niceStrings += 1
		}
	}

	return niceStrings
}

func part2(lines []string) int {
	niceStrings := 0
	for _, line := range lines {
		if isNiceString2(line) {
			niceStrings += 1
		}
	}

	return niceStrings
}

func isNiceString1(line string) bool {
	hasNaughtyCombination, err := regexp.Match(`(ab|cd|pq|xy)`, []byte(line))
	if err != nil {
		log.Fatal(err)
	}

	if hasNaughtyCombination {
		return false
	}

	var prevR rune
	vowelCount := 0
	hasLetterTwiceInARow := false
	for _, r := range line {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowelCount += 1
		}

		if r == prevR {
			hasLetterTwiceInARow = true
		}

		prevR = r
	}

	return hasLetterTwiceInARow && vowelCount >= 3
}

func isNiceString2(line string) bool {
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
