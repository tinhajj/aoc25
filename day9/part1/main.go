package main

import (
	"aoc25/scan"
	"fmt"
	"os"
	"strings"
)

type Tile struct {
	X, Y int
}

func main() {
	b, err := os.ReadFile("day_9_input.txt")
	// b, err := os.ReadFile("day_9_sample_input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]

	tiles := []Tile{}

	largest := 0
	for _, line := range lines {
		nums := scan.Numbers(line)
		t := Tile{
			X: nums[0],
			Y: nums[1],
		}
		tiles = append(tiles, t)
	}

	for i, t1 := range tiles {
		for _, t2 := range tiles[i:] {
			a := Area(t1, t2)
			if a > largest {
				largest = a
			}
		}
	}

	fmt.Println(largest)
}

func Area(t1, t2 Tile) int {
	x := Abs(t1.X-t2.X) + 1
	y := Abs(t1.Y-t2.Y) + 1

	result := x * y
	if result < 0 {
		return result * -1
	}
	return result
}

func Abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
