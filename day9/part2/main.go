package main

import (
	"aoc25/scan"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"slices"
	"strings"
)

type Segment struct {
	Start Vec2
	End   Vec2
}

type Color int

const (
	ColorBlank Color = iota
	ColorRed         = iota
	ColorGreen

	ColorPurple
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
			Color: ColorRed,
			Pos:   Vec2{nums[0], nums[1]},
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

	xlookupInverse := map[int]int{}
	ylookupInverse := map[int]int{}

	start := 1
	for _, x := range xs {
		xlookup[x] = start
		xlookupInverse[start] = x
		start++
	}

	start = 1
	for _, y := range ys {
		ylookup[y] = start
		ylookupInverse[start] = y
		start++
	}

	largestX := 0
	largestY := 0

	for i, t := range redTiles {
		redTiles[i] = Tile{
			Color: ColorRed,
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

	for y, row := range grid {
		for x := range row {
			grid[y][x] = Tile{
				Pos: Vec2{
					X: x,
					Y: y,
				},
				Color: ColorBlank,
			}
		}
	}

	for _, t := range redTiles {
		t.Color = ColorRed
		grid[t.Pos.Y][t.Pos.X] = t
	}

	verticalSegments := []Segment{}

	for i, tile := range redTiles[:len(redTiles)-1] {
		nextTile := redTiles[i+1]
		MakeGreenBetween(grid, tile, nextTile)
		if tile.Pos.X == nextTile.Pos.X {
			verticalSegments = append(verticalSegments, Segment{
				Start: tile.Pos,
				End:   nextTile.Pos,
			})
		}
	}
	MakeGreenBetween(grid, redTiles[0], redTiles[len(redTiles)-1])
	if redTiles[0].Pos.X == redTiles[len(redTiles)-1].Pos.X {
		verticalSegments = append(verticalSegments, Segment{
			Start: redTiles[0].Pos,
			End:   redTiles[len(redTiles)-1].Pos,
		})
	}
	Draw(grid)

	var inside Tile
	for y, row := range grid {
		found := false
		for x, t := range row {
			if t.Pos.X == 0 && t.Pos.Y == 2 {
				// continue
				_ = rand.Intn(3)
			}
			hits := CastRay(grid, verticalSegments, t)
			if hits > 0 && hits%2 != 0 && t.Color != ColorRed && t.Color != ColorGreen {
				found = true
				inside = grid[y][x]
				break
			}
		}

		if found {
			break
		}
	}

	Bfs(grid, inside)
	Draw(grid)

	var largestV1 Vec2
	var largestV2 Vec2
	largestArea := 0

	for i, tile1 := range redTiles {
	outer:
		for _, tile2 := range redTiles[i+1:] {

			minX := min(tile1.Pos.X, tile2.Pos.X)
			minY := min(tile1.Pos.Y, tile2.Pos.Y)
			maxX := max(tile1.Pos.X, tile2.Pos.X)
			maxY := max(tile1.Pos.Y, tile2.Pos.Y)

			for x := minX; x <= maxX; x++ {
				for y := minY; y <= maxY; y++ {
					if grid[y][x].Color == ColorBlank {
						continue outer
					}
				}
			}

			v1 := Vec2{
				X: xlookupInverse[tile1.Pos.X],
				Y: ylookupInverse[tile1.Pos.Y],
			}
			v2 := Vec2{
				X: xlookupInverse[tile2.Pos.X],
				Y: ylookupInverse[tile2.Pos.Y],
			}

			area := Area(v1, v2)
			if area > largestArea {
				largestV1 = tile1.Pos
				largestV2 = tile2.Pos
				largestArea = area
			}
		}
	}

	// 92808 too low
	// 113979918 too low
	Draw(grid)

	fmt.Println(largestArea)
	fmt.Println(largestV1, largestV2)

	minX := min(largestV1.X, largestV2.X)
	minY := min(largestV1.Y, largestV2.Y)
	maxX := max(largestV1.X, largestV2.X)
	maxY := max(largestV1.Y, largestV2.Y)

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			grid[y][x].Color = ColorPurple
		}
	}

	Draw(grid)
}

func Bfs(grid [][]Tile, start Tile) {
	queue := []Tile{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		tile := grid[current.Pos.Y][current.Pos.X]
		if tile.Color == ColorBlank {
			grid[tile.Pos.Y][tile.Pos.X].Color = ColorGreen
		} else {
			continue
		}
		around := Around(grid, current)
		if len(around) > 0 {
			queue = append(queue, around...)
		}
	}
}

func Around(grid [][]Tile, t Tile) []Tile {
	v := t.Pos
	directions := []Vec2{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	adj := []Tile{}
	for _, d := range directions {
		other := AddVecs(v, d)
		if Oob(grid, other) {
			continue
		}
		color := grid[other.Y][other.X].Color
		if color == ColorGreen || color == ColorRed {
			continue
		}
		adj = append(adj, grid[other.Y][other.X])
	}
	return adj
}

func Oob(grid [][]Tile, v Vec2) bool {
	if v.X < 0 || v.Y < 0 {
		return true
	}

	if v.Y >= len(grid) {
		return true
	}

	row := grid[v.Y]
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

func CastRay(grid [][]Tile, verticalSegments []Segment, tile Tile) int {
	hitCount := 0
	if tile.Color != ColorBlank {
		return 0
	}

	for _, segment := range verticalSegments {
		if segment.Start.X < tile.Pos.X {
			continue
		}
		minY := min(segment.Start.Y, segment.End.Y)
		maxY := max(segment.Start.Y, segment.End.Y)

		if tile.Pos.Y >= minY && tile.Pos.Y < maxY {
			hitCount++
		}
	}
	return hitCount

	// width, _ := len(grid[0]), len(grid)

	// crossCount := 0

	// for i := start.X; i < width-1; i++ {
	// 	current := grid[start.Y][i]
	// 	next := grid[start.Y][i+1]

	// 	currentColored := current.Color == ColorRed || current.Color == ColorGreen
	// 	nextColored := next.Color == ColorRed || next.Color == ColorGreen

	// 	if nextColored {
	// 		if currentColored != nextColored {
	// 			crossCount++
	// 			i++
	// 		}
	// 	}
	// }

	// return crossCount
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
			case ColorPurple:
				img.Set(x, y, color.RGBA{R: 255, B: 255, A: 255})
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
