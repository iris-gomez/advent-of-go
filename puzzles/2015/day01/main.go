package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := getInput()
	fmt.Println("Part 1 solution:", part1(input))
	fmt.Println("Part 2 solution:", part2(input))
}

func getInput() string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(input))
}

func part1(input string) int {
	floor := 0
	for _, char := range input {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}
	}

	return floor
}

func part2(input string) int {
	floor, position := 0, 1
	for i, char := range input {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor == -1 {
			position += i
			break
		}
	}

	return position
}
