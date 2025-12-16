package main

import (
	"aoc25/scan"
	"fmt"
	"math"
	"os"
	"strings"
)

type Vec2 struct {
	X int
	Y int
}

type TileGrid map[int]map[int]bool

func (tg TileGrid) Add(v Vec2) {
	m, ok := tg[v.X]
	if !ok {
		m = map[int]bool{}
		tg[v.X] = m
	}
	tg[v.X][v.Y] = true
}

type Tile struct {
	Pos Vec2
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
		t := Tile{
			Pos: Vec2{
				X: nums[0],
				Y: nums[1],
			},
		}
		tiles = append(tiles, t)
	}

	minX := math.MaxInt
	minY := math.MaxInt
	maxX := 0
	maxY := 0

	for _, tile := range tiles {
		if tile.Pos.X < minX {
			minX = tile.Pos.X
		}
		if tile.Pos.X > maxX {
			maxX = tile.Pos.X
		}
		if tile.Pos.Y < minY {
			minY = tile.Pos.Y
		}
		if tile.Pos.Y > maxY {
			maxY = tile.Pos.Y
		}
	}

	grid := make(TileGrid)

	for i, tile := range tiles[:len(tiles)-1] {
		nextTile := tiles[i+1]

		grid.Add(tile.Pos)
		grid.Add(nextTile.Pos)

		if tile.Pos.X == nextTile.Pos.X {
			end := max(tile.Pos.Y, nextTile.Pos.Y)
			start := min(tile.Pos.Y, nextTile.Pos.Y)

			for i := start + 1; i < end; i++ {
				grid.Add(Vec2{X: tile.Pos.X, Y: i})
			}
		} else if tile.Pos.Y == nextTile.Pos.Y {
			end := max(tile.Pos.Y, nextTile.Pos.Y)
			start := min(tile.Pos.Y, nextTile.Pos.Y)

			for i := start + 1; i < end; i++ {
				grid.Add(Vec2{X: tile.Pos.X, Y: i})
			}
		} else {
			panic("impossible tile ordering")
		}
	}

	fmt.Println(minX, minY, maxX, maxY)
}

func Area(t1, t2 Tile) int {
	x := Abs(t1.Pos.X-t2.Pos.X) + 1
	y := Abs(t1.Pos.Y-t2.Pos.Y) + 1

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
