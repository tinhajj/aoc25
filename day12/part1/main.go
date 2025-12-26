package main

import (
	"os"
	"strings"
)

type Present struct {
	TileTotal int
}

func main() {
	b, err := os.ReadFile("day_12_input.txt")
	// b, err := os.ReadFile("day_12_sample_input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]

	var present Present
	for _, line := range lines {
		if strings.Contains(line, ":") {
		}
	}
}
