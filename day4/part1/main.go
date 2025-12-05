package main

import (
	"fmt"
	"os"
	"strings"
)

type Vec2 struct {
	X, Y int
}

type Spot struct {
	X, Y int
	V    string
}

func main() {
	b, err := os.ReadFile("day_4_input.txt")
	// b, err := os.ReadFile("day_4_sample_input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]

	matrix := [][]*Spot{}
	adj := map[*Spot][]*Spot{}

	for i, line := range lines {
		row := []*Spot{}

		for j, c := range line {
			row = append(row, &Spot{X: j, Y: i, V: string(c)})
		}
		matrix = append(matrix, row)
	}

	sum := 0

	for _, row := range matrix {
		for _, spot := range row {
			if spot.V != "@" {
				continue
			}
			a := around(matrix, Vec2{X: spot.X, Y: spot.Y}, "@")
			adj[spot] = a
			if len(a) < 4 {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

func around(matrix [][]*Spot, v Vec2, char string) []*Spot {
	directions := []Vec2{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}

	adj := []*Spot{}
	for _, d := range directions {
		other := AddVecs(v, d)
		if Oob(matrix, other) {
			continue
		}
		if matrix[other.Y][other.X].V != "@" {
			continue
		}
		adj = append(adj, matrix[other.Y][other.X])
	}
	return adj
}

func Oob(matrix [][]*Spot, v Vec2) bool {
	if v.X < 0 || v.Y < 0 {
		return true
	}

	if v.Y >= len(matrix) {
		return true
	}

	row := matrix[v.Y]
	if v.X >= len(row) {
		return true
	}

	return false
}

func AddVecs(v1, v2 Vec2) Vec2 {
	return Vec2{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}
