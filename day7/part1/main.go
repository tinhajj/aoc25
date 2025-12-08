package main

import (
	"aoc25/scan"
	"fmt"
	"os"
	"strings"
)

func main() {
	b, err := os.ReadFile("day_7_input.txt")
	// b, err := os.ReadFile("day_7_sample_input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]

	m, _, _ := scan.RuneMatrix(lines)

	for i, r := range m[0] {
		if r == "S" {
			m[0][i] = "|"
		}
	}

	sum := 0

	for i, row := range m {
		if i == len(m)-1 {
			break
		}
		for j, r := range row {
			if r != "|" {
				continue
			}
			if m[i+1][j] == "^" {
				sum += 1
				left := j - 1
				right := j + 1

				if left >= 0 && left < len(m[i+1]) {
					m[i+1][j-1] = "|"
				}
				if right >= 0 && right < len(m[i+1]) {
					m[i+1][j-1] = "|"
				}

				m[i+1][j+1] = "|"
			} else {
				m[i+1][j] = "|"
			}
		}
	}

	fmt.Println(sum)
}
