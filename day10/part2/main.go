package main

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

type Machine struct {
	LightGoal   []bool
	VoltageGoal []int
	Wirings     [][]int
}

func main() {
	// b, err := os.ReadFile("day_10_input.txt")
	b, err := os.ReadFile("day_10_sample_input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]

	machines := []Machine{}

	for _, line := range lines {
		start := strings.Index(line, "[")
		end := strings.Index(line, "]")

		lights := make([]bool, len(line[start+1:end]))

		for i, r := range line[start+1 : end] {
			switch r {
			case '#':
				lights[i] = true
			case '.':
				lights[i] = false
			}
		}

		parts := strings.Split(line, " ")
		wirings := [][]int{}
		for _, part := range parts[1 : len(parts)-1] {
			part = part[1 : len(part)-1]
			ps := strings.Split(part, ",")

			wiring := []int{}

			for _, p := range ps {
				num, err := strconv.Atoi(p)
				if err != nil {
					panic(err)
				}
				wiring = append(wiring, num)
			}

			wirings = append(wirings, wiring)
		}

		voltages := []int{}
		section := parts[len(parts)-1]
		section = section[1 : len(section)-1]
		for _, v := range strings.Split(section, ",") {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			voltages = append(voltages, num)
		}

		m := Machine{
			LightGoal:   lights,
			Wirings:     wirings,
			VoltageGoal: voltages,
		}
		machines = append(machines, m)
	}
}

func Search(voltages []int, wirings [][]int, presses int) int {
	invalid, solved := VoltageStatus(voltages)
	if invalid {
		return -1
	}
	if solved {
		return presses
	}
	for _, w := range wirings {
		newVoltage := slices.Clone(voltages)
		VoltageApply(newVoltage, w)
		Search(newVoltage, wirings, presses+1)
	}
}

func VoltageApply(voltages []int, wirings []int) {
	for _, w := range wirings {
		voltages[w] -= 1
	}
}

func VoltageStatus(voltages []int) (invalid bool, solved bool) {
	for _, v := range voltages {
		if v < 0 {
			return true, false
		}
	}

	for _, v := range voltages {
		if v != 0 {
			return false, false
		}
	}

	return true, true
}
