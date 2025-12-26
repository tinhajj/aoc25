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
	b, err := os.ReadFile("day_10_input.txt")
	// b, err := os.ReadFile("day_10_sample_input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]

	machines := []Machine{}

	sum := 0
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

		x := Solve(m, sets, m.VoltageGoal)
		sum += x
		fmt.Println(x)
	}
	fmt.Println("Sum:", sum)
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
		ok := ValidSubset(m, goal, set)
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

func Solve(m Machine, sets [][]int, voltages []int) int {
	result := 1_000_000

	if VoltageEmpty(voltages) {
		return 0
	}

	subWirings := ValidWirings(m, VoltageToBool(voltages), sets)
	if VoltageAllEven(voltages) {
		subWirings = append(subWirings, [][]int{})
	}

	for _, wiring := range subWirings {
		subVoltages := slices.Clone(voltages)
		VoltageApply(subVoltages, wiring)

		if VoltageInvalid(subVoltages) {
			continue
		}

		VoltageDivide(subVoltages)
		subPresses := Solve(m, sets, subVoltages)
		if subPresses >= 1_000_000 {
			continue
		}
		x := len(wiring) + (subPresses * 2)
		result = min(result, x)
	}

	return result
}

func VoltageAllEven(voltages []int) bool {
	for _, v := range voltages {
		if v%2 != 0 {
			return false
		}
	}
	return true
}

func VoltageApply(voltages []int, wirings [][]int) {
	for _, wiring := range wirings {
		for _, w := range wiring {
			voltages[w] -= 1
		}
	}
}

func VoltageDivide(voltages []int) {
	for i := range voltages {
		voltages[i] = voltages[i] / 2
	}
}

func VoltageEmpty(voltages []int) bool {
	for _, v := range voltages {
		if v != 0 {
			return false
		}
	}
	return true
}

func VoltageInvalid(voltages []int) bool {
	for _, v := range voltages {
		if v < 0 {
			return true
		}
	}
	return false
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

func ValidSubset(m Machine, goal []bool, choices []int) bool {
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
		return false
	}
	return true
}
