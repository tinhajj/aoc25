package main

import (
	"aoc25/scan"
	"fmt"
	"os"
	"strings"
)

type Operation int

const (
	Add Operation = iota
	Multiply
)

type Problem struct {
	Nums []int
	Op   Operation
}

func main() {
	b, err := os.ReadFile("day_6_input.txt")
	// b, err := os.ReadFile("day_6_sample_input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]

	problems := make([]Problem, len(scan.Numbers(lines[0])))

	for i := 0; i < len(lines)-1; i++ {
		line := lines[i]
		nums := scan.Numbers(line)
		for j, n := range nums {
			problems[j].Nums = append(problems[j].Nums, n)
		}
	}

	ops := scan.Strings(lines[len(lines)-1])

	for i, op := range ops {
		switch op {
		case "+":
			problems[i].Op = Add
		case "*":
			problems[i].Op = Multiply
		}
	}

	sum := 0
	for _, p := range problems {
		r := 0
		switch p.Op {
		case Add:
			for _, num := range p.Nums {
				r += num
			}
		case Multiply:
			for _, num := range p.Nums {
				if r == 0 {
					r = num
				} else {
					r *= num
				}
			}
		}
		sum += r
	}
	fmt.Println(sum)
}
