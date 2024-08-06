package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var wiring = make(map[string][]string)
var values = make(map[string]uint16)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	solution := part1(lines)
	fmt.Printf("Part 1 solution: %d\n", solution)
	fmt.Printf("Part 2 solution: %d\n", part2(solution))
}

func part1(lines []string) uint16 {
	for _, line := range lines {
		wire, operation := parseLine(line)
		wiring[wire] = operation
	}

	return solve("a")
}

func part2(overrides uint16) uint16 {
	// No reason to recreate the wiring map since it hasn't changed

	// Clearing the value cache, otherwise we'd get the same result again
	clear(values)

	// Overriding the value of wire "b" with the result of part 1
	wiring["b"] = []string{fmt.Sprint(overrides)}

	return solve("a")
}

func parseLine(line string) (string, []string) {
	split := strings.Split(line, " -> ")
	wire := split[1]
	operation := strings.Split(split[0], " ")
	return wire, operation
}

func solve(wire string) uint16 {
	u, err := strconv.ParseUint(wire, 10, 16)
	if err == nil {
		return uint16(u)
	}

	cached, ok := values[wire]
	if ok {
		return cached
	}

	var result uint16
	parts := wiring[wire]
	switch len(parts) {
	case 1:
		u, err := strconv.ParseUint(parts[0], 10, 16)
		if err == nil {
			result = uint16(u)
			values[wire] = result
		} else {
			result = solve(parts[0])
			values[wire] = result
		}
	case 2:
		result = ^solve(parts[1])
		values[wire] = result
	case 3:
		switch parts[1] {
		case "AND":
			result = solve(parts[0]) & solve(parts[2])
			values[wire] = result
		case "OR":
			result = solve(parts[0]) | solve(parts[2])
			values[wire] = result
		case "LSHIFT":
			result = solve(parts[0]) << solve(parts[2])
			values[wire] = result
		case "RSHIFT":
			result = solve(parts[0]) >> solve(parts[2])
			values[wire] = result
		}
	}

	return result
}
