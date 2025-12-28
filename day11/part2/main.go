package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Memo struct {
	FindsFft bool
	FindsDac bool

	WaysToEndThroughBoth int
	WaysToEndThroughDac  int
	WaysToEndThroughFft  int
	WaysToEnd            int
}

func main() {
	start := time.Now()
	b, err := os.ReadFile("day_11_input.txt")
	// b, err := os.ReadFile("day_11_sample_input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]

	adj := map[string][]string{}

	for _, line := range lines {
		parts := strings.Split(line, ":")
		device := parts[0]
		outputs := parts[1]
		adj[device] = strings.Split(outputs, " ")[1:]
	}

	visited := map[string]bool{}
	cache := map[string]Memo{}

	result, _, _ := dfs(adj, visited, cache, "svr", "out")
	fmt.Println(time.Since(start))

	fmt.Println(result)
}

var largest = 0

func dfs(adj map[string][]string, visited map[string]bool, cache map[string]Memo, start string, goal string) (int, int, map[string]bool) {
	visits := map[string]bool{}
	visits[start] = true

	if start == goal {
		if visited["fft"] && visited["dac"] {
			return 1, 1, visits
		}
		return 0, 1, visits
	}

	visited[start] = true

	neighbors := adj[start]
	sumThroughFftAndDac := 0
	sumGoal := 0

	for _, n := range neighbors {
		if visited[n] {
			panic("loop back, didn't expect")
		}

		memo, ok := cache[n]

		if ok {
			if !visited["fft"] && !memo.FindsFft {
				continue
			}
			if !visited["dac"] && !memo.FindsDac {
				continue
			}
			if visited["dac"] && visited["fft"] {
				sumThroughFftAndDac += memo.WaysToEndThroughBoth
				sumGoal += memo.WaysToEnd
				continue
			}
		}

		resultThroughFftAndDac, resultGoal, subVisits := dfs(adj, visited, cache, n, goal)
		for k, v := range subVisits {
			visits[k] = v
		}
		sumThroughFftAndDac += resultThroughFftAndDac
		sumGoal += resultGoal
	}

	visited[start] = false

	vf := false
	vd := false
	for _, n := range neighbors {
		memo, ok := cache[n]
		if ok {
			vf = vf || memo.FindsFft
			vd = vd || memo.FindsDac
		}
	}

	cache[start] = Memo{
		FindsFft:             visits["fft"] || vf,
		FindsDac:             visits["dac"] || vd,
		WaysToEndThroughBoth: sumThroughFftAndDac,
		WaysToEnd:            sumGoal,
	}

	if sumThroughFftAndDac > largest {
		largest = sumThroughFftAndDac
		fmt.Println(sumThroughFftAndDac)
	}

	return sumThroughFftAndDac, sumGoal, visits
}
