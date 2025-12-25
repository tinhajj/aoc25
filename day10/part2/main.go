// https://old.reddit.com/r/adventofcode/comments/1pk87hl/2025_day_10_part_2_bifurcate_your_way_to_victory/
package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Memo struct {
	Lookup map[string]int
}

func (m *Memo) Add(voltage []int, presses int) {
	pieces := []string{}
	for _, v := range voltage {
		pieces = append(pieces, strconv.Itoa(v))
	}
	k := strings.Join(pieces, ",")
	m.Lookup[k] = presses
}

func (m *Memo) Get(voltage []int) (int, bool) {
	pieces := []string{}
	for _, v := range voltage {
		pieces = append(pieces, strconv.Itoa(v))
	}
	k := strings.Join(pieces, ",")
	v, ok := m.Lookup[k]
	return v, ok
}

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

		bases := make([]int, len(m.Wirings))
		for i := range bases {
			bases[i] = i
		}

		sets := Subsets(bases)

		goal := VoltageToBool(m.VoltageGoal)

		validWirings := ValidWirings(m, goal, sets)

		smallest := math.MaxInt
		for _, wiring := range validWirings {
			x := Solve(m, sets, m.VoltageGoal, wiring)
			if x < smallest {
				smallest = x
			}
		}
		fmt.Println(smallest)
	}
}

func VoltageToBool(voltages []int) []bool {
	goal := []bool{}
	for _, v := range voltages {
		switch v%2 == 0 {
		case true:
			goal = append(goal, false)
		case false:
			goal = append(goal, true)
		}
	}
	return goal
}

func ValidWirings(m Machine, goal []bool, sets [][]int) [][][]int {
	valid := [][][]int{}
	for _, set := range sets {
		ok, _ := ValidSubset(m, goal, set)
		if ok {
			wires := [][]int{}
			for _, i := range set {
				wires = append(wires, m.Wirings[i])
			}
			valid = append(valid, wires)
		}
	}
	return valid
}

func Solve(m Machine, sets [][]int, voltages []int, wiring [][]int) int {
	subVoltages := slices.Clone(voltages)

	presses := 0
	for _, w := range wiring {
		VoltageApply(subVoltages, w, 1)
	}
	presses += len(wiring)

	invalid, solved := VoltageStatus(subVoltages)
	if invalid {
		return 1_000_000
	}
	if solved {
		return presses
	}

	VoltageDivide(subVoltages)
	fmt.Println(subVoltages)

	subWirings := ValidWirings(m, VoltageToBool(subVoltages), sets)
	if len(subWirings) < 1 {
		return 1_000_000
	}

	smallest := math.MaxInt
	for _, w := range subWirings {
		subPresses := Solve(m, sets, subVoltages, w)
		x := (subPresses * 2) + presses
		if x < smallest {
			smallest = x
		}
	}

	return smallest
}

func VoltageApply(voltages []int, wirings []int, times int) {
	for i := 0; i < times; i++ {
		for _, w := range wirings {
			voltages[w] -= 1
		}
	}
}

func VoltageDivide(voltages []int) {
	for i := range voltages {
		voltages[i] = voltages[i] / 2
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

	return false, true
}

func VoltageLowest(voltages []int) int {
	lowest := math.MaxInt
	for _, v := range voltages {
		if v < lowest {
			lowest = v
		}
	}
	return lowest
}

func Subsets(bases []int) [][]int {
	results := [][]int{}
	for _, b := range bases {
		for _, result := range results {
			copy := slices.Clone(result)
			copy = append(copy, b)
			results = append(results, copy)
		}
		results = append(results, []int{b})
	}
	return results
}

func ValidSubset(m Machine, goal []bool, choices []int) (bool, int) {
	allFalse := true

	for _, g := range goal {
		if g == true {
			allFalse = false
			break
		}
	}

	if allFalse {
		return true, 0
	}

	wires := [][]int{}
	for _, choice := range choices {
		wires = append(wires, m.Wirings[choice])
	}
	sum := map[int]int{}
	for _, w := range wires {
		for _, b := range w {
			sum[b] = sum[b] + 1
		}
	}
	failed := false
	for i, on := range goal {
		if on && sum[i]%2 == 0 {
			failed = true
			break
		}
		if !on && sum[i]%2 != 0 {
			failed = true
			break
		}
	}
	if failed {
		return false, 0
	}
	return true, len(choices)
}
