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
		var largest, largestIdx int
		var nextLargest int
		for i, v := range bank[:len(bank)-1] {
			if v > largest {
				largest = v
				largestIdx = i
			}
		}

		for i := largestIdx + 1; i < len(bank); i++ {
			v := bank[i]
			if v > nextLargest {
				nextLargest = v
			}
		}

		sum += (largest * 10) + nextLargest
	}
	fmt.Println(sum)
}
