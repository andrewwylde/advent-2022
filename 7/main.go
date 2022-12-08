package main

import (
	"advent/elfutils"
	"fmt"
	"log"
	"os"
)

func unique(chars []string) bool {
	seen := make(map[string]bool)
	for _, v := range chars {
		if seen[v] {
			return false
		} else {
			seen[v] = true
		}
	}
	return true
}

func doWork(lines []string) (result int) {

	return
}

func test() {
	input := elfutils.GetTestInputByDay("6")
	lines := elfutils.SplitByLine(string(input))
	result := doWork(lines)
	x := 21
	if result != x {
		log.Fatalf("Test failed! Expected %v, got %v", x, result)
	} else {
		fmt.Printf("success: %v\n", result)
	}
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
		input := elfutils.GetInputByDay("6")
		lines := elfutils.SplitByLine(string(input))
		result := doWork(lines)
		fmt.Printf("result: %v\n", result)
	}
}
