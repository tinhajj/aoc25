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

	freshRanges := []*FreshRange{}

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

		freshRanges = append(freshRanges, &FreshRange{Start: start, End: end})
	}

	sum := 0

	// augment ranges
	augmented := true
outer:
	for augmented {
		augmented = false

		for i := 0; i < len(freshRanges); i++ {
			first := freshRanges[i]
			for j := i + 1; j < len(freshRanges); j++ {
				other := freshRanges[j]

				ok, expandedRange := overlaps(first, other)
				if ok {
					freshRanges[i] = expandedRange
					freshRanges = append(freshRanges[:j], freshRanges[j+1:]...)
					augmented = true
					continue outer
				}
			}
		}
	}

	for _, fr := range freshRanges {
		sum += fr.End - fr.Start + 1
	}
	fmt.Println(sum)
}

func overlaps(fr1, fr2 *FreshRange) (bool, *FreshRange) {
	overlapped := false
	if between(fr1.Start, fr2) {
		overlapped = true
	}
	if between(fr1.End, fr2) {
		overlapped = true
	}
	if between(fr2.Start, fr1) {
		overlapped = true
	}
	if between(fr2.End, fr1) {
		overlapped = true
	}

	if overlapped {
		return true, &FreshRange{
			Start: min(fr1.Start, fr2.Start),
			End:   max(fr1.End, fr2.End),
		}
	}
	return false, nil
}

func between(n int, fr *FreshRange) bool {
	if n >= fr.Start && n <= fr.End {
		return true
	}
	return false
}
