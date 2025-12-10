package main

import (
	"aoc25/scan"
	"fmt"
	"math"
	"os"
	"strings"
)

type Box struct {
	X, Y, Z     int
	Connections []*Box
}

func (b *Box) ConnectedTo(other *Box) bool {
	for _, connection := range b.Connections {
		if connection == other {
			return true
		}
	}
	return false
}

func main() {
	b, err := os.ReadFile("day_8_input.txt")
	// b, err := os.ReadFile("day_8_sample_input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]

	boxes := []*Box{}

	for _, l := range lines {
		nums := scan.Numbers(l)
		b := &Box{
			X: nums[0],
			Y: nums[1],
			Z: nums[2],
		}

		boxes = append(boxes, b)
	}

	// Need to do at least 1000
	for i := 0; i < 1000; i++ {
		b1, b2 := Shortest(boxes)
		b1.Connections = append(b1.Connections, b2)
		b2.Connections = append(b2.Connections, b1)
	}

	var b1, b2 *Box
	for !FullCircuit(boxes) {
		b1, b2 = Shortest(boxes)
		b1.Connections = append(b1.Connections, b2)
		b2.Connections = append(b2.Connections, b1)
	}
	fmt.Println(b1.X * b2.X)
}

func FullCircuit(boxes []*Box) bool {
	visited := map[*Box]bool{}

	length := 0
	queue := []*Box{boxes[0]}
	visited[boxes[0]] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		length++

		for _, o := range cur.Connections {
			if visited[o] {
				continue
			}
			visited[o] = true
			queue = append(queue, o)
		}
	}

	return length == len(boxes)
}

func Shortest(boxes []*Box) (*Box, *Box) {
	shortest := math.MaxFloat64
	var first, second *Box

	for i, b1 := range boxes {
		for _, b2 := range boxes[i+1:] {
			if b1.ConnectedTo(b2) {
				continue
			}
			d := Distance(b1, b2)
			if d < shortest {
				shortest = d
				first, second = b1, b2
			}
		}
	}

	return first, second
}

func Distance(b1, b2 *Box) float64 {
	p1 := (b1.X - b2.X) * (b1.X - b2.X)
	p2 := (b1.Y - b2.Y) * (b1.Y - b2.Y)
	p3 := (b1.Z - b2.Z) * (b1.Z - b2.Z)

	f1 := float64(p1)
	f2 := float64(p2)
	f3 := float64(p3)

	return math.Sqrt(f1 + f2 + f3)
}
