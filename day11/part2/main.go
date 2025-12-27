package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Memo struct {
	HitsFft   bool
	HitsDac   bool
	WaysToEnd int
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

	result, _, _ := dfs(adj, map[string]bool{}, map[string]Memo{}, "svr", "out")
	fmt.Println(time.Since(start))

	fmt.Println(result)
}

func dfs(adj map[string][]string, visited map[string]bool, cache map[string]Memo, start string, goal string) (int, bool, bool) {
	if start == goal {
		if visited["fft"] && visited["dac"] {
			return 1, true, true
		}
		return 0, false, false
	}

	memo, ok := cache[start]
	if ok {
		if !visited["fft"] && !memo.HitsFft {
			return 0, memo.HitsFft, memo.HitsDac
		}
		if !visited["dac"] && !memo.HitsDac {
			return 0, memo.HitsFft, memo.HitsDac
		}
	}

	visited[start] = true

	neighbors := adj[start]
	sum := 0

	hitsFft := visited["fft"]
	hitsDac := visited["dac"]

	for _, n := range neighbors {
		if visited[n] {
			continue
		}
		result, hitFft, hitDac := dfs(adj, visited, cache, n, goal)
		hitsFft = hitsFft || hitFft
		hitsDac = hitsDac || hitDac
		sum += result
	}

	visited[start] = false

	return sum, hitsFft, hitsDac
}
