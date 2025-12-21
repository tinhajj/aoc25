package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Machine struct {
	LightGoal []bool
	Wirings   [][]int
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

	sum := 0

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

		m := Machine{
			LightGoal: lights,
			Wirings:   wirings,
		}
		machines = append(machines, m)

		bases := make([]int, len(m.Wirings))
		for i := range bases {
			bases[i] = i
		}

		sets := subsets(bases)
		shortest := math.MaxInt
		for _, s := range sets {
			ok, size := solve(m, s)
			if ok && size < shortest {
				shortest = size
			}
		}
		sum += shortest
	}
	fmt.Println(sum)
}

func solve(m Machine, choices []int) (bool, int) {
	wires := [][]int{}
	for _, choice := range choices {
		wires = append(wires, m.Wirings[choice])
	}
	sum := map[int]int{}
	for _, w := range wires {
		for _, b := range w {
			sum[b] = sum[b] + 1
		}
	}
	failed := false
	for i, on := range m.LightGoal {
		if on && sum[i]%2 == 0 {
			failed = true
			break
		}
		if !on && sum[i]%2 != 0 {
			failed = true
			break
		}
	}
	if failed {
		return false, 0
	}
	return true, len(choices)
}

func subsets(bases []int) [][]int {
	results := [][]int{}
	for _, b := range bases {
		for _, result := range results {
			copy := slices.Clone(result)
			copy = append(copy, b)
			results = append(results, copy)
		}
		results = append(results, []int{b})
	}
	return results
}
