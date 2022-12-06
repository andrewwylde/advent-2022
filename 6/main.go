package main

import (
	"advent/elfutils"
	"fmt"
	"log"
	"os"
	"strings"
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
	for _, v := range lines {
		chars := strings.Split(v, "")
		last4 := chars[0:4]
		for i, v := range chars[4:] {
			if unique(last4) {
				result = i + 4
				break
			} else {
				last4 = append(last4[1:], v)
			}
		}
		fmt.Printf("result: %v\n", result)
	}
	return
}

func test() {
	input := elfutils.GetTestInputByDay("6")
	lines := elfutils.SplitByLine(string(input))
	result := doWork(lines)
	if result != 11 {
		log.Fatalf("Test failed! Expected 11, got %v", result)
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
