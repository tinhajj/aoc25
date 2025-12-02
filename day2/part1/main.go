package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("day_2_input.txt")
	// b, err := os.ReadFile("day_2_sample_input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]
	line := lines[0]

	sum := 0

	ranges := strings.Split(line, ",")
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		var start, end int
		var err error
		start, err = strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		end, err = strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)
			if repeated(s) {
				sum += i
			}
		}
	}
	fmt.Println(sum)
}

func repeated(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	mid := len(s) / 2
	first := s[:mid]
	second := s[mid:]

	return first == second
}
