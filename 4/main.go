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
	if result != 2 {
		log.Fatalf("Test failed! Expected 2, got %v", result)
	} else {
		fmt.Printf("result: %v\n", result)
	}
}

func getPairCount(input []string) (sum int) {

	for _, v := range input {
		elves := strings.Split(v, ",")
		elf1 := getStartEnd(elves[0])
		elf2 := getStartEnd(elves[1])

		if elves[0] == elves[1] {
			sum += 1
		} else {

			if compareAss(elf1, elf2) {
				sum += 1
			}
			if compareAss(elf2, elf1) {
				sum += 1
			}
		}

	}
	return
}

func compareAss(a Assignment, b Assignment) bool {
	return a.start <= b.start && a.end >= b.end
}

type Assignment struct {
	start int
	end   int
}

func getStartEnd(elf string) (ass Assignment) {
	for i, v := range strings.Split(elf, "-") {
		if i == 0 {
			start, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf(err.Error())
			}
			ass.start = start
		} else {
			end, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf(err.Error())
			}
			ass.end = end

		}
	}
	return
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
