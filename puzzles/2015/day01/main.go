package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1 solution: %d\n", part1(input))
	fmt.Printf("Part 2 solution: %d\n", part2(input))
}

func part1(input []byte) int {
	floor := 0
	for _, char := range input {
		switch char {
		case '(':
			floor += 1
		case ')':
			floor -= 1
		}
	}

	return floor
}

func part2(input []byte) int {
	position := 1
	floor := 0
	for i, char := range input {
		switch char {
		case '(':
			floor += 1
		case ')':
			floor -= 1
		}

		if floor == -1 {
			position += i
			break
		}
	}

	return position
}
