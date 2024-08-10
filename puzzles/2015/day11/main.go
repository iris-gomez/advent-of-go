package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	password := strings.TrimSpace(string(input))

	solution := part1(password)
	fmt.Printf("Part 1 solution: %v\n", solution)
	fmt.Printf("Part 2 solution: %v\n", part2(solution))
}

func part1(password string) string {
	for isSecurePassword(password) == false {
		password = incrementLetter(len(password)-1, password)
	}

	return password
}

func part2(password string) string {
	password = incrementLetter(len(password)-1, password)
	for isSecurePassword(password) == false {
		password = incrementLetter(len(password)-1, password)
	}

	return password
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
