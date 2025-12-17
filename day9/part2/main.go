package main

import (
	"aoc25/scan"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"slices"
	"strings"
)

type Color int

const (
	ColorBlank Color = iota
	ColorRed         = iota
	ColorGreen
)

type Tile struct {
	Pos   Vec2
	Color Color
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

	redTiles := []Tile{}

	for _, line := range lines {
		nums := scan.Numbers(line)
		redTiles = append(redTiles, Tile{
			Pos: Vec2{nums[0], nums[1]},
		})
	}

	// unique x and y
	ux := map[int]struct{}{}
	uy := map[int]struct{}{}

	for _, t := range redTiles {
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

	largestX := 0
	largestY := 0

	for i, t := range redTiles {
		redTiles[i] = Tile{
			Pos: Vec2{
				X: xlookup[t.Pos.X],
				Y: ylookup[t.Pos.Y],
			},
		}

		if xlookup[t.Pos.X] > largestX {
			largestX = xlookup[t.Pos.X]
		}
		if ylookup[t.Pos.Y] > largestY {
			largestY = ylookup[t.Pos.Y]
		}
	}

	largestX += 10
	largestY += 10

	grid := make([][]Tile, largestY)
	for i := range grid {
		grid[i] = make([]Tile, largestX)
	}

	for _, t := range redTiles {
		t.Color = ColorRed
		grid[t.Pos.Y][t.Pos.X] = t
	}

	for i, tile := range redTiles[:len(redTiles)-1] {
		nextTile := redTiles[i+1]
		MakeGreenBetween(grid, tile, nextTile)
	}
	MakeGreenBetween(grid, redTiles[0], redTiles[len(redTiles)-1])

	for y, row := range grid {
		found := false
		for x := range row {
			point := Vec2{X: x, Y: y}
			hits := CastRay(grid, point)
			if hits > 0 && hits%2 != 0 {
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	Draw(grid)
}

func CastRay(grid [][]Tile, start Vec2) int {
	width, _ := len(grid[0]), len(grid)

	crossCount := 0

	for i := start.X; i < width-1; i++ {
		current := grid[start.Y][i]
		next := grid[start.Y][i+1]

		currentColored := current.Color == ColorRed || current.Color == ColorGreen
		nextColored := next.Color == ColorRed || next.Color == ColorGreen

		if currentColored != nextColored {
			crossCount++
		}
	}

	return crossCount
}

func Draw(grid [][]Tile) {
	width, height := len(grid[0]), len(grid)

	// 1. Create a new image rectangle with the specified bounds.
	// image.Rect defines the minimum and maximum coordinates.
	imgRect := image.Rect(0, 0, width, height)
	// Create a new RGBA image. This type implements the image.Image interface.
	img := image.NewRGBA(imgRect)

	// 2. Iterate over the pixels and set their colors.
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			tile := grid[y][x]

			switch tile.Color {
			case ColorRed:
				img.Set(x, y, color.RGBA{R: 255, G: 0, A: 255})
			case ColorGreen:
				img.Set(x, y, color.RGBA{R: 0, G: 255, A: 255})
			}
			// Set the pixel color using a color.RGBA struct
		}
	}

	// 3. Create the output file.
	f, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 4. Encode the image data to the file in PNG format.
	if err = png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}

func MakeGreenBetween(grid [][]Tile, t1, t2 Tile) {
	if t1.Pos.X == t2.Pos.X {
		minY := min(t1.Pos.Y, t2.Pos.Y)
		maxY := max(t1.Pos.Y, t2.Pos.Y)
		for i := minY; i <= maxY; i++ {
			if grid[i][t1.Pos.X].Color == ColorRed {
				continue
			}
			grid[i][t1.Pos.X] = Tile{
				Color: ColorGreen,
				Pos:   Vec2{X: t1.Pos.X, Y: i},
			}
		}
	}
	if t1.Pos.Y == t2.Pos.Y {
		minX := min(t1.Pos.X, t2.Pos.X)
		maxX := max(t1.Pos.X, t2.Pos.X)
		for i := minX; i <= maxX; i++ {
			if grid[t1.Pos.Y][i].Color == ColorRed {
				continue
			}
			grid[t1.Pos.Y][i] = Tile{
				Color: ColorGreen,
				Pos:   Vec2{X: i, Y: t1.Pos.Y},
			}
		}
	}
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
