package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Input struct {
	Start string
	IsFft bool
	IsDac bool
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

	cache := map[Input]int{}

	ways := dfs(adj, cache, false, false, "svr", "out")
	fmt.Println(time.Since(start))

	fmt.Println(ways)
}

func dfs(adj map[string][]string, cache map[Input]int, isFft, isDac bool, start string, goal string) int {
	memo, ok := cache[Input{Start: start, IsFft: isFft, IsDac: isDac}]
	if ok {
		return memo
	}

	if start == goal {
		if isFft && isDac {
			return 1
		}
	}

	neighbors := adj[start]

	total := 0
	for _, n := range neighbors {
		total += dfs(adj, cache, isFft || n == "fft", isDac || n == "dac", n, goal)
	}

	cache[Input{Start: start, IsFft: isFft, IsDac: isDac}] = total

	return total
}
