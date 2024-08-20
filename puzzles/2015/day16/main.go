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
	aunts := getAunts(lines)
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
	aunts := getAunts(lines)
	for i, aunt := range aunts {
		if isRealGiftingAunt(tape, aunt) {
			number = i + 1
		}
	}
	return number
}

func getAunts(lines []string) []map[string]int {
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

func isGiftingAunt(tape, sub map[string]int) bool {
	if len(sub) > len(tape) {
		return false
	}
	for k, vsub := range sub {
		if vm, found := tape[k]; !found || vm != vsub {
			return false
		}
	}
	return true
}

func isRealGiftingAunt(tape, aunt map[string]int) bool {
	if len(aunt) > len(tape) {
		return false
	}
	for k, vaunt := range aunt {
		vtape, found := tape[k]
		if !found {
			return false
		}

		switch k {
		case "cats", "trees":
			if vtape >= vaunt {
				return false
			}
		case "pomeranians", "goldfish":
			if vtape <= vaunt {
				return false
			}
		default:
			if vtape != vaunt {
				return false
			}
		}
	}
	return true
}
