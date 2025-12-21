package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Machine struct {
	Lights  []bool
	Wirings [][]int
}

func main() {
	b, err := os.ReadFile("day_10_input.txt")
	// b, err := os.ReadFile("day_10_sample_input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]

	machines := []Machine{}

	for _, line := range lines {
		start := strings.Index(line, "[")
		end := strings.Index(line, "]")

		lights := make([]bool, len(line[start+1:end]))

		for i, r := range line[start+1 : end] {
			switch r {
			case '#':
				lights[i] = true
			case '.':
				lights[i] = false
			}
		}

		parts := strings.Split(line, " ")
		wirings := [][]int{}
		for _, part := range parts[1 : len(parts)-1] {
			part = part[1 : len(part)-1]
			ps := strings.Split(part, ",")

			wiring := []int{}

			for _, p := range ps {
				num, err := strconv.Atoi(p)
				if err != nil {
					panic(err)
				}
				wiring = append(wiring, num)
			}

			wirings = append(wirings, wiring)
		}

		machines = append(machines, Machine{
			Lights:  lights,
			Wirings: wirings,
		})

	}
	results := [][]int{}
	combinations([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, &results)
	fmt.Println(len(results))
}

func combinationsRec(prefix []int, choices []int, skiplist []bool, skipped int, combos *[][]int) {
	for i, choice := range choices {
		newPrefix := slices.Clone(prefix)
		newPrefix = append(newPrefix, choice)
		*combos = append(*combos, newPrefix)

		newChoices := slices.Clone(choices)
		newChoices = append(newChoices[:i], newChoices[i+1:]...)

		combinationsRec(newPrefix, newChoices, combos)
	}
}

func combinations(bases []int, combos *[][]int) {
	skiplist := make([]bool, len(bases))
	for i, b := range bases {
		*combos = append(*combos, []int{b})
		skiplist[i] = true
		combinationsRec([]int{b}, bases, skiplist, i, combos)
		skiplist[i] = false
	}
}
