package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := getInput()
	solution := part1(input)
	fmt.Println("Part 1 solution:", solution)
	fmt.Println("Part 2 solution:", part2(solution))
}

func getInput() string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(input))
}

func part1(input string) string {
	for isSecurePassword(input) == false {
		input = incrementLetter(len(input)-1, input)
	}

	return input
}

func part2(input string) string {
	input = incrementLetter(len(input)-1, input)
	for isSecurePassword(input) == false {
		input = incrementLetter(len(input)-1, input)
	}

	return input
}

func incrementLetter(letter int, password string) string {
	var nextPassword string

	chars := []byte(password)
	switch chars[letter] {
	case 'z':
		chars[letter] = 'a'
		nextPassword = incrementLetter(letter-1, string(chars))
	case 'h', 'k', 'n':
		// skip i, l, and o
		chars[letter] += 2
		nextPassword = string(chars)
	default:
		chars[letter]++
		nextPassword = string(chars)
	}

	return nextPassword
}

func isSecurePassword(password string) bool {
	forbidden := "ilo"
	for i := 0; i < len(forbidden); i++ {
		if strings.IndexByte(password, forbidden[i]) != -1 {
			return false
		}
	}

	hasStraight := false
	for i := 0; i < len(password)-2; i++ {
		if password[i]+1 == password[i+1] && password[i+1] == password[i+2]-1 {
			hasStraight = true
		}
	}

	pairs := 0
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			pairs++
			i++
		}
	}

	return hasStraight && pairs >= 2
}
