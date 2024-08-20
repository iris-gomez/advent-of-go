package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
	for i := 0; i < 40; i++ {
		input = lookAndSay(input)
	}

	return len(input)
}

func part2(input string) int {
	for i := 0; i < 50; i++ {
		input = lookAndSay(input)
	}

	return len(input)
}

func lookAndSay(sequence string) string {
	var currentDigit rune
	var output strings.Builder
	var reader strings.Reader

	reader = *strings.NewReader(sequence)

	counter := 1
	currentDigit, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}

	for {
		digit, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		if digit == currentDigit {
			counter++
		} else {
			output.WriteString(strconv.Itoa(counter))
			output.WriteRune(currentDigit)
			currentDigit = digit
			counter = 1
		}
	}

	output.WriteString(strconv.Itoa(counter))
	output.WriteRune(currentDigit)

	return output.String()
}
