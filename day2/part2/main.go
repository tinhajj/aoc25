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
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i != 0 {
			continue
		}

		if repeatedRun(s, i) {
			return true
		}
	}

	return false
}

func repeatedRun(s string, run int) bool {
	chunks := []string{}
	for len(s) > 0 {
		chunks = append(chunks, s[:run])
		s = s[run:]
	}
	first := chunks[0]

	for i := 1; i < len(chunks); i++ {
		if first != chunks[i] {
			return false
		}
	}

	return true
}
