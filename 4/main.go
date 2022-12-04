package main

import (
	"advent/elfutils"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func test() {
	input := elfutils.GetTestInputByDay("4")
	lines := elfutils.SplitByLine(string(input))
	result := getPairCount(lines)
	if result != 4 {
		log.Fatalf("Test failed! Expected 4, got %v", result)
	} else {
		fmt.Printf("result: %v\n", result)
	}
}

func getPairCount(input []string) (sum int) {

	for _, v := range input {
		elves := strings.Split(v, ",")
		a := getStartEnd(elves[0])
		b := getStartEnd(elves[1])
		/*
		   elf has cross over if their ranges overlap at all.
		   overlap - startA <= endB OR if endA
		*/
		if elves[0] == elves[1] {
			fmt.Printf("same")
			sum += 1
		} else {

			x := compareAss(a, b)
			logShit(a, b, x)
			if x {
				sum += 1
			}
		}

	}
	return
}

func logShit(a Assignment, b Assignment, x bool) {
	fmt.Printf("x: %v\n", x)
}

func compareAss(a Assignment, b Assignment) bool {
	fmt.Printf("a: %+v\n", a)
	fmt.Printf("b: %+v\n", b)
	if b.start <= a.end && b.end >= a.start {
		fmt.Printf("b start <= a.end")
		return true
	}
	if a.start <= b.end && a.end >= b.start {
		fmt.Printf("a start <= b.end")
		return true
	}
	return false

}

type Assignment struct {
	start float64
	end   float64
}

func getStartEnd(elf string) (ass Assignment) {
	for i, v := range strings.Split(elf, "-") {
		if i == 0 {
			start, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf(err.Error())
			}
			ass.start = float64(start)
		} else {
			end, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf(err.Error())
			}
			ass.end = float64(end)

		}
	}
	return ass
}

func main() {
	args := os.Args
	isTest := false
	for _, v := range args {
		if v == "test" {
			isTest = true
		}
	}
	if isTest {
		test()
	} else {
		input := elfutils.GetInputByDay("4")
		lines := elfutils.SplitByLine(string(input))
		result := getPairCount(lines)
		fmt.Printf("result: %v\n", result)
	}
}
