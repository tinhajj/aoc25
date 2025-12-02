package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	Left Direction = iota
	Right
)

type Turn struct {
	Direction Direction
	Amount    int
}

func main() {
	b, err := os.ReadFile("day_1_input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]

	turns := []Turn{}

	for _, line := range lines {
		t := Turn{}

		switch line[0] {
		case 'L':
			t.Direction = Left
		case 'R':
			t.Direction = Right
		}

		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		t.Amount = amount

		turns = append(turns, t)
	}

	position := 50
	count := 0

	for _, t := range turns {
		var interval int
		amount := t.Amount % 100
		switch t.Direction {
		case Left:
			interval = -1 * amount
		case Right:
			interval = amount
		}

		position = position + interval
		if position < 0 {
			position = 100 + position
		} else if position > 99 {
			position = position - 100
		}

		if position == 0 {
			count++
		}
	}

	fmt.Println(count)
}
