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

	result := dfs(adj, map[string]bool{}, map[string]bool{}, "svr", "fft")

	fmt.Println(result)
	// bfs(adj, "you")
}

func copy(input map[string]bool) map[string]bool {
	result := map[string]bool{}
	for k, v := range input {
		result[k] = v
	}
	return result
}

func dfs(adj map[string][]string, visited map[string]bool, skip map[string]bool, start string, goal string) int {
	visited[start] = true
	if start == goal {
		fmt.Println("found")
		return 1
		if visited["fft"] && visited["dac"] {
			return 1
		}
		return 0
	}

	neighbors := adj[start]
	sum := 0
	for _, n := range neighbors {
		if visited[n] {
			continue
		}
		result := dfs(adj, copy(visited), skip, n, goal)
		sum += result
	}

	return sum
}

func bfs(adj map[string][]string, start string) {
	queue := []string{start}
	visited := map[string]bool{}
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Println(current)

		adjacents := adj[current]
		for _, a := range adjacents {
			if visited[a] {
				continue
			}
			visited[a] = true
			queue = append(queue, a)
		}
	}
}
