package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operation int

const (
	Add Operation = iota
	Multiply
)

type Problem struct {
	Nums       []int
	Op         Operation
	OpLocation int
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

	longestLine := 0

	for _, l := range lines {
		if len(l) > longestLine {
			longestLine = len(l)
		}
	}

	problems := []*Problem{}

	for i, r := range lines[len(lines)-1] {
		c := string(r)
		if c != "+" && c != "*" {
			continue
		}

		var op Operation
		switch c {
		case "+":
			op = Add
		case "*":
			op = Multiply
		}

		problems = append(problems, &Problem{
			Op:         op,
			OpLocation: i,
		})
	}

	for i, problem := range problems {
		nextProblemLocation := longestLine
		if i+1 < len(problems) {
			nextProblemLocation = problems[i+1].OpLocation - 1
		}

		for j := problem.OpLocation; j < nextProblemLocation; j++ {
			buf := ""
			for _, l := range lines[:len(lines)-1] {
				if string(l[j]) == " " {
					continue
				}
				buf += string(l[j])
			}

			if buf == "" {
				continue
			}

			i, err := strconv.Atoi(buf)
			if err != nil {
				panic(err)
			}
			problem.Nums = append(problem.Nums, i)
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
