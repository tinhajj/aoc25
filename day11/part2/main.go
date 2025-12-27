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
	_ = cache

	result, _, _ := dfs(adj, visited, "svr", "out")
	fmt.Println(time.Since(start))

	fmt.Println(result)
}

func dfs(adj map[string][]string, visited map[string]bool, start string, goal string) (int, bool, bool) {
	if start == "fft" {
		fmt.Println("found fft")
	}

	if start == goal {
		if visited["fft"] && visited["dac"] {
			fmt.Println("found one")
			return 1, true, true
		}
		return 0, false, false
	}

	visited[start] = true

	neighbors := adj[start]
	sum := 0

	for _, n := range neighbors {
		if visited[n] {
			fmt.Println("hmm")
			continue
		}
		result, _, _ := dfs(adj, visited, n, goal)
		sum += result
	}

	visited[start] = false

	return sum, false, false
}
