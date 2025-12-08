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
	// b, err := os.ReadFile("day_7_test.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]

	m, _, _ := scan.RuneMatrix(lines)
	ways := make([][]int, len(m))

	for i := range ways {
		ways[i] = make([]int, len(m[0]))
	}

	for i, r := range m[0] {
		if r == "S" {
			m[0][i] = "|"
			ways[0][i] = 1
		}
	}

	for i, row := range m {
		if i == len(m)-1 {
			break
		}
		for j, r := range row {
			if r != "|" {
				continue
			}
			if m[i+1][j] == "^" {
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

	for i, row := range m {
		if i == 0 {
			continue
		}
		for j, r := range row {
			if r != "|" {
				continue
			}
			ways[i][j] = above3Sum(ways, m, i, j)
		}
	}

	sum := 0
	for _, w := range ways[len(ways)-1] {
		sum += w
	}
	fmt.Println(sum)
}

type Vec2 struct {
	X int
	Y int
}

func Add(v1, v2 Vec2) Vec2 {
	return Vec2{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

func Oob(matrix [][]int, v Vec2) bool {
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

func above3Sum(m [][]int, sm [][]string, i, j int) int {
	origin := Vec2{Y: i, X: j}

	sum := 0

	left := Add(origin, Vec2{X: -1, Y: -1})
	middle := Add(origin, Vec2{X: 0, Y: -1})
	right := Add(origin, Vec2{X: 1, Y: -1})

	if !Oob(m, left) && sm[i][j-1] == "^" {
		sum += m[left.Y][left.X]
	}
	if !Oob(m, middle) {
		sum += m[middle.Y][middle.X]
	}
	if !Oob(m, right) && sm[i][j+1] == "^" {
		sum += m[right.Y][right.X]
	}

	return sum
}
