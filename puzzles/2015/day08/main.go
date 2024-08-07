package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
	result := 0
	for _, line := range lines {
		unquoted, err := strconv.Unquote(line)
		if err != nil {
			log.Fatal(err)
		}

		result += len(line) - len(unquoted)
	}

	return result
}

func part2(lines []string) int {
	result := 0
	for _, line := range lines {
		result += len(strconv.QuoteToASCII(line)) - len(line)
	}

	return result
}
