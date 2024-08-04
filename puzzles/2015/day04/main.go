package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
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
	secret := strings.TrimSpace(string(input))

	fmt.Printf("Part 1 solution: %d\n", part1(secret))
	fmt.Printf("Part 2 solution: %d\n", part2(secret))
}

func part1(secret string) int {
	number := 0
	for {
		h := md5.New()
		io.WriteString(h, secret+strconv.Itoa(number))
		sum := hex.EncodeToString(h.Sum(nil)[:])
		if strings.HasPrefix(sum, "00000") {
			break
		}

		number++
	}
	return number
}

func part2(secret string) int {
	number := 0
	for {
		h := md5.New()
		io.WriteString(h, secret+strconv.Itoa(number))
		sum := hex.EncodeToString(h.Sum(nil)[:])
		if strings.HasPrefix(sum, "000000") {
			break
		}

		number++
	}
	return number
}
