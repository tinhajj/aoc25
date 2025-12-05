package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PowerBank []int

func main() {
	b, err := os.ReadFile("day_3_input.txt")
	// b, err := os.ReadFile("day_3_sample_input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]

	banks := []PowerBank{}

	for _, line := range lines {
		bank := PowerBank{}
		for _, c := range line {
			i, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			bank = append(bank, i)
		}
		banks = append(banks, bank)
	}

	sum := 0

	for _, bank := range banks {
		slidingWindowSize := len(bank) - 12 + 1
		digits := []int{}
		start := 0

		for len(digits) < 12 {
			var largest, largestIdx int

			for i := start; i < start+slidingWindowSize; i++ {
				v := bank[i]
				if v > largest {
					largest = v
					largestIdx = i
				}
			}
			slidingWindowSize = slidingWindowSize - (largestIdx - start)
			start = largestIdx + 1

			digits = append(digits, largest)
		}

		fullNumber := 0
		multiple := 1
		for i := len(digits) - 1; i >= 0; i-- {
			fullNumber += digits[i] * multiple
			multiple *= 10
		}
		sum += fullNumber
	}
	fmt.Println(sum)
}
