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
	var number int
	tape := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	aunts := parseLines(lines)
	for i, aunt := range aunts {
		if isGiftingAunt(tape, aunt) {
			number = i + 1
		}
	}
	return number
}

func part2(lines []string) int {
	var number int
	tape := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	aunts := parseLines(lines)
	for i, aunt := range aunts {
		if isTrueGiftingAunt(tape, aunt) {
			number = i + 1
		}
	}
	return number
}

func parseLines(lines []string) []map[string]int {
	aunts := []map[string]int{}
	for _, line := range lines {
		split := strings.Split(strings.ReplaceAll(strings.ReplaceAll(line, ",", ""), ":", ""), " ")

		aunt := make(map[string]int)
		for i := 2; i < len(split); i += 2 {
			value, err := strconv.Atoi(split[i+1])
			if err != nil {
				log.Fatal(err)
			}
			aunt[split[i]] = value
		}

		aunts = append(aunts, aunt)
	}

	return aunts
}

func isGiftingAunt(m, sub map[string]int) bool {
	if len(sub) > len(m) {
		return false
	}
	for k, vsub := range sub {
		if vm, found := m[k]; !found || vm != vsub {
			return false
		}
	}
	return true
}

func isTrueGiftingAunt(m, sub map[string]int) bool {
	if len(sub) > len(m) {
		return false
	}
	for k, vsub := range sub {
		vm, found := m[k]
		if !found {
			return false
		}

		switch k {
		case "cats", "trees":
			if vm >= vsub {
				return false
			}
		case "pomeranians", "goldfish":
			if vm <= vsub {
				return false
			}
		default:
			if vm != vsub {
				return false
			}
		}
	}
	return true
}
