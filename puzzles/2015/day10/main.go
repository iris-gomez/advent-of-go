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
	sequence := strings.TrimSpace(string(input))

	fmt.Printf("Part 1 solution: %v\n", part1(sequence))
	fmt.Printf("Part 2 solution: %v\n", part2(sequence))
}

func part1(sequence string) int {
	for i := 0; i < 40; i++ {
		sequence = lookAndSay(sequence)
	}

	return len(sequence)
}

func part2(sequence string) int {
	for i := 0; i < 50; i++ {
		sequence = lookAndSay(sequence)
	}

	return len(sequence)
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
