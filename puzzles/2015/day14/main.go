package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Reindeer struct {
	name      string
	speed     int
	endurance int
	rests     int
	distance  int
	points    int
}

func (r Reindeer) getStatus(time int) string {
	var status string
	for {
		time = time - r.endurance
		if time < 0 {
			status = "flying"
			break
		}

		time = time - r.rests
		if time < 0 {
			status = "resting"
			break
		}
	}
	return status
}

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
	reindeers := parseLines(lines)
	maxDistance := 0
	for _, reindeer := range reindeers {
		for i := 0; i < 2503; i++ {
			if reindeer.getStatus(i) == "flying" {
				reindeer.distance += reindeer.speed
			}
		}
		maxDistance = max(maxDistance, reindeer.distance)
	}
	return maxDistance
}

func part2(lines []string) int {
	reindeers := parseLines(lines)
	maxDistance, maxPoints := 0, 0
	for i := 0; i < 2503; i++ {
		for j := 0; j < len(reindeers); j++ {
			if reindeers[j].getStatus(i) == "flying" {
				reindeers[j].distance += reindeers[j].speed
				maxDistance = max(maxDistance, reindeers[j].distance)
			}
		}
		for j := 0; j < len(reindeers); j++ {
			if reindeers[j].distance == maxDistance {
				reindeers[j].points++
				maxPoints = max(maxPoints, reindeers[j].points)
			}
		}
	}
	return maxPoints
}

func parseLines(lines []string) []Reindeer {
	reindeers := []Reindeer{}
	for _, line := range lines {
		split := strings.Split(line, " ")

		name := split[0]

		speed, err := strconv.Atoi(split[3])
		if err != nil {
			log.Fatal(err)
		}

		endurance, err := strconv.Atoi(split[6])
		if err != nil {
			log.Fatal(err)
		}

		rests, err := strconv.Atoi(split[13])
		if err != nil {
			log.Fatal(err)
		}

		reindeers = append(reindeers, Reindeer{name, speed, endurance, rests, 0, 0})
	}
	return reindeers
}
