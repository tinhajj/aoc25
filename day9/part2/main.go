package main

import (
	"aoc25/scan"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Tile struct {
	Pos Vec2
}

type Vec2 struct {
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

	for _, line := range lines {
		nums := scan.Numbers(line)
		tiles = append(tiles, Tile{
			Pos: Vec2{nums[0], nums[1]},
		})
	}

	// unique x and y
	ux := map[int]struct{}{}
	uy := map[int]struct{}{}

	for _, t := range tiles {
		ux[t.Pos.X] = struct{}{}
		uy[t.Pos.Y] = struct{}{}
	}

	xs := []int{}
	ys := []int{}

	for k := range ux {
		xs = append(xs, k)
	}
	for k := range uy {
		ys = append(ys, k)
	}

	slices.Sort(xs)
	slices.Sort(ys)

	xlookup := map[int]int{}
	ylookup := map[int]int{}

	start := 1
	for _, x := range xs {
		xlookup[x] = start
		start++
	}

	start = 1
	for _, y := range ys {
		ylookup[y] = start
		start++
	}

	for i, t := range tiles {
		tiles[i] = Tile{
			Pos: Vec2{
				X: xlookup[t.Pos.X],
				Y: ylookup[t.Pos.Y],
			},
		}
	}

	fmt.Println(tiles)
	// at this point the tiles are compressed correctly?
}

func Area(v1, v2 Vec2) int {
	x := Abs(v1.X-v2.X) + 1
	y := Abs(v1.Y-v2.Y) + 1

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
