package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
