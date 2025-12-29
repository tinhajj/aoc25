package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Memo struct {
	Visited map[string]bool

	WaysGoal int
	WaysFft  int
	WaysDac  int
	WaysBoth int
}

func main() {
	start := time.Now()
	// b, err := os.ReadFile("day_11_input.txt")
	b, err := os.ReadFile("day_11_sample_input.txt")
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

	goalWays, fftWays, dacWays, bothWays, _ := dfs(adj, visited, cache, "svr", "out")
	fmt.Println(time.Since(start))

	fmt.Println(goalWays, fftWays, dacWays, bothWays)
}

var largest = 0

func dfs(adj map[string][]string, visited map[string]bool, cache map[string]Memo, start string, goal string) (waysAny, waysFft, waysDac, waysBoth int, v map[string]bool) {
	visits := map[string]bool{}
	visits[start] = true

	if start == goal {
		return 1, 0, 0, 0, visits
	}

	visited[start] = true

	neighbors := adj[start]

	var outerWaysAny, outerWaysFft, outerWaysDac, outerWaysBoth int

	for _, n := range neighbors {
		if visited[n] {
			panic("loop back, didn't expect")
		}

		memo, ok := cache[n]

		var subVisits map[string]bool

		if ok {
			if visited["fft"] && visited["dac"] {
				outerWaysAny += memo.WaysGoal
				outerWaysFft += memo.WaysGoal
				outerWaysDac += memo.WaysGoal
				outerWaysBoth += memo.WaysGoal
			}

			if visited["fft"] {
				outerWaysAny += memo.WaysGoal
				outerWaysFft += memo.WaysGoal
				outerWaysDac += memo.WaysDac
				outerWaysBoth += memo.WaysDac
			}

			if visited["dac"] {
				outerWaysAny += memo.WaysGoal
				outerWaysDac += memo.WaysGoal
				outerWaysFft += memo.WaysFft
				outerWaysBoth += memo.WaysFft
			}

			if !visited["dac"] && !visited["fft"] {
				outerWaysAny += memo.WaysGoal
				outerWaysDac += memo.WaysDac
				outerWaysFft += memo.WaysFft
				outerWaysBoth += memo.WaysBoth
			}
		}

		if !ok {
			goalWays, fftWays, dacWays, bothWays, subVisits = dfs(adj, visited, cache, n, goal)

			outerGoalWays += goalWays
			outerFftWays += fftWays
			outerDacWays += dacWays
			outerBothWays += bothWays
		}

		for k, v := range subVisits {
			visits[k] = v
		}
	}

	visited[start] = false

	vf := false
	vd := false
	for _, n := range neighbors {
		memo, ok := cache[n]
		if ok {
			vf = vf || memo.Visited["fft"]
			vd = vd || memo.Visited["dac"]
		}
		for k, v := range memo.Visited {
			visits[k] = v
		}
	}

	cache[start] = Memo{
		WaysGoal: outerGoalWays,
		WaysFft:  outerFftWays,
		WaysDac:  outerDacWays,
		Visited:  visits,
	}

	return outerGoalWays, outerFftWays, outerDacWays, outerBothWays, visits
}
