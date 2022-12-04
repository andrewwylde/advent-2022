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
CrZsJsPPZsGzwwsLwLmpwMDw
`
	fmt.Printf("testInput: %v\n", testInput)
	x := splitByLine(testInput)
	sum := 0
	for _, v := range x {
		r := parseLine((v))
		sum += int(runeToPriority(r))
		fmt.Printf("priority: %v\n", runeToPriority(r))
	}
	fmt.Printf("sum: %v\n", sum)

}

func parseLine(v string) rune {
	fmt.Printf("v: %v\n", v)
	strLen := len(v)
	mid := strLen / 2
	lm := make(map[int32]int32)
	rm := make(map[int32]int32)
	foundRune := rune(0)

	for i, _ := range strings.Split(v, "")[0:mid] {
		lp := rune(v[i])
		rp := rune(v[mid+i])
		if lm[lp] == 0 {
			lm[lp] = 1
		}
		if rm[rp] == 0 {
			rm[rp] = 1
		}
		if lm[rp] != 0 {
			foundRune = rp
		}
		if rm[lp] != 0 {
			foundRune = lp
		}
	}
	return foundRune
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

	d, _ := os.Getwd()
	p := path.Join(d, "3", "input.txt")
	f, _ := os.ReadFile(p)

	x := splitByLine(string(f))
	sum := 0
	for _, v := range x {
		r := parseLine((v))
		sum += int(runeToPriority(r))
		fmt.Printf("priority: %v\n", runeToPriority(r))
	}
	fmt.Printf("sum: %v\n", sum)

}
