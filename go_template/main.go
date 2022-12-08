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

func test(day string) {
	input := elfutils.GetTestInputByDay(day)
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
	fmt.Printf("args: %v\n", args)
	os.Exit(0)
	isTest := false
	day := args[1]
	for _, v := range args {
		if v == "test" {
			isTest = true
		}
	}
	if isTest {
		test(day)
	} else {
		input := elfutils.GetInputByDay(day)
		lines := elfutils.SplitByLine(string(input))
		result := doWork(lines)
		fmt.Printf("result: %v\n", result)
	}
}
