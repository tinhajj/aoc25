package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Present struct {
	TileTotal int
}

type Region struct {
	Width               int
	Height              int
	PresentRequirements []int
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
	var presents []Present
	var regions []Region

	var i int
	var line string

	for i, line = range lines {
		if strings.Contains(line, "x") && strings.Contains(line, ":") {
			break
		}

		if strings.Contains(line, ":") {
			continue
		}

		for _, r := range line {
			s := string(r)
			if s == "#" {
				present.TileTotal += 1
			}
		}

		if line == "" {
			presents = append(presents, present)
			present = Present{}
		}
	}

	for _, line := range lines[i:] {
		region := Region{}

		parts := strings.Split(line, ":")

		p := strings.Split(parts[0], "x")
		sHeight, sWidth := p[0], p[1]
		height, _ := strconv.Atoi(sHeight)
		width, _ := strconv.Atoi(sWidth)

		region.Height = height
		region.Width = width
		region.PresentRequirements = []int{}

		p = strings.Split(parts[1][1:], " ")
		for _, r := range p {
			s := string(r)
			num, _ := strconv.Atoi(s)
			region.PresentRequirements = append(region.PresentRequirements, num)
		}

		regions = append(regions, region)
	}

	total := 0

	for _, region := range regions {
		totalSize := region.Height * region.Width
		totalPresentRequirements := region.TotalPresentRequirements()

		if totalPresentRequirements <= (region.Width/3)*(region.Height/3) {
			total += 1
			continue
		}

		tileRequirement := 0
		for i, req := range region.PresentRequirements {
			present := presents[i]
			tileRequirement += present.TileTotal * req
		}

		if tileRequirement > totalSize {
			continue
		}

		panic("tricky case")
	}
	fmt.Println(total)
	// fmt.Println(presents, regions)
}

func (r Region) TotalPresentRequirements() int {
	sum := 0
	for _, req := range r.PresentRequirements {
		sum += req
	}
	return sum
}
