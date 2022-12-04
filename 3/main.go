package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
	A        = "Rock"
	B        = "Paper"
	C        = "Scissors"
	X        = "Rock"
	Y        = "Paper"
	Z        = "Scissors"
)

func test() {
	testInput := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	x := splitByLine(testInput)
	sum := 0
	for start := 0; start < len(x); start += 3 {
		end := start + 3
		x1 := x[start:end]
		r := parseGroup(x1)
		fmt.Printf("x1: %v\n", x1)
		fmt.Printf("start: %v\n", start)
		fmt.Printf("end: %v\n", end)

		fmt.Printf("r: %v %s \n", r, string(r))
		sum += int(runeToPriority(r))
	}

	fmt.Printf("sum: %v\n", sum)

}

func parseGroup(s []string) (r rune) {
	rootmap := make(map[rune]int)
	for _, v := range s {
		r := parseSack(v)
		for k, v2 := range r {
			rootmap[k] += v2
		}
	}
	for k, v := range rootmap {
		if v == 3 {
			r = k
		}
	}
	return
}

func parseSack(v string) map[rune]int {
	m := make(map[rune]int)
	for _, v2 := range v {
		// fmt.Printf("rune in sack: %v %s\n", v2, string(v2))
		m[v2] = 1
	}

	return m
}

func runeToPriority(r rune) int32 {
	if r > 96 {
		return r - 96
	}
	if r > 0 {
		return r - 26 - 12
	}
	return 0

}

func splitByLine(t string) []string {
	tmp := []string{}
	for _, v := range strings.Split(t, "\n") {
		if len(v) > 1 {
			tmp = append(tmp, v)
		}
	}
	return tmp
}

func main() {
	test()

	d, _ := os.Getwd()
	p := path.Join(d, "3", "input.txt")
	f, _ := os.ReadFile(p)

	x := splitByLine(string(f))
	sum := 0
	for i := 0; i < len(x); i += 3 {
		r := parseGroup(x[i : i+3])
		sum += int(runeToPriority(r))
	}
	fmt.Printf("sum: %v\n", sum)

}
