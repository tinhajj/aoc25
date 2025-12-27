package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
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

	result := dfs(adj, map[string]bool{}, "svr", "out")

	fmt.Println(result)
}

func dfs(adj map[string][]string, visited map[string]bool, start string, goal string) int {
	if start == goal {
		return 1
	}

	visited[start] = true

	neighbors := adj[start]
	sum := 0

	for _, n := range neighbors {
		if visited[n] {
			continue
		}
		result := dfs(adj, visited, n, goal)
		sum += result
	}

	visited[start] = false

	return sum
}
