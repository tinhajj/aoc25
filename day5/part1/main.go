package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FreshRange struct {
	Start int
	End   int
}

func main() {
	b, err := os.ReadFile("day_5_input.txt")
	// b, err := os.ReadFile("day_5_sample_input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]

	i := 0

	freshRanges := []FreshRange{}

	for i = 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		freshRanges = append(freshRanges, FreshRange{Start: start, End: end})
	}

	i++
	availables := []int{}

	for ; i < len(lines); i++ {
		line := lines[i]
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		availables = append(availables, num)
	}

	sum := 0

outer:
	for _, available := range availables {
		for _, fr := range freshRanges {
			if available >= fr.Start && available <= fr.End {
				sum++
				continue outer
			}
		}
	}

	fmt.Println(sum)
}
