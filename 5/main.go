package main

import (
	"advent/elfutils"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseStacks(lines []string) (moveLines []string, stacks [][]string) {

	stackCount := (len(lines[0]) + 1) / 4
	stackLine := []string{}
	// fmt.Printf("stackCount: %v\n", stackCount)
	for i := 0; i < stackCount; i++ {
		stacks = append(stacks, []string{})
	}

	movesIndex := 0

	for i, v := range lines {
		isMove, err := regexp.MatchString(`1`, v)
		if err != nil {
			log.Fatalf("error encountered tyring to matve: %+v\n", err)
		}
		if isMove {
			movesIndex = i + 1
			break
		} else {
			stackLine = append(stackLine, v)
		}
	}

	for _, v := range stackLine {
		// fmt.Printf("now parsing line: %v\n", i)

		for i2, j := 0, 0; i2 <= len(v); i2 += 4 {
			// fmt.Printf("i: %v\n", i)
			item := v[i2 : i2+3]
			x := regexp.MustCompile(`\[([A-Z])\]`)
			it := x.FindStringSubmatch(item)
			if len(it) > 0 {
				// fmt.Printf("item: %v\n", it[1])
				// fmt.Printf("stacks: %v\n", stacks)
				// fmt.Printf("stacks[i]: %v\n", stacks[j])
				newItems := append(stacks[j], it[1])
				(stacks)[j] = newItems
			}
			j++
		}
	}
	for i := movesIndex; i < len(lines); i++ {
		v := lines[i]
		moveLines = append(moveLines, v)
	}
	// fmt.Printf("--------\n\n")
	return
}

func doWork(lines []string) string {

	moves, stacks := parseStacks(lines)
	// fmt.Printf("stacks: %v\n", stacks)
	result := doMoves(moves, stacks)

	return strings.Join(result, "")
}

func doMoves(moves []string, stacks [][]string) (result []string) {
	// fmt.Printf("\n-----------------\ndoMoves\n")
	r := regexp.MustCompile(`move (\d{1,2}) from (\d{1,2}) to (\d{1,2})`)
	for _, v := range moves {
		// fmt.Printf("move: %v\n", v)
		// fmt.Printf("stacks: %v\n", stacks)
		// fmt.Println("--------")
		m := r.FindStringSubmatch(v)[1:]
		count, _ := strconv.Atoi(m[0])
		x, _ := strconv.Atoi(m[1])
		from := x - 1
		x1, _ := strconv.Atoi(m[2])
		to := x1 - 1
		for i := 0; i < count; i++ {
			// fmt.Printf("count: %v\n", count)
			// fmt.Printf("from: %v\n", from)
			// fmt.Printf("to: %v\n", to)
			fromStack := stacks[from]
			toStack := stacks[to]
			// fmt.Printf("fromStack: %v\n", fromStack)
			// fmt.Printf("toStack: %v\n", toStack)
			item := fromStack[0]
			stacks[from] = fromStack[1:]
			stacks[to] = append([]string{item}, toStack...)
		}
	}
	for _, v := range stacks {
		result = append(result, v[0])
	}
	return
}

func test() {
	input := elfutils.GetTestInputByDay("5")
	lines := elfutils.SplitByLine(string(input))
	result := doWork(lines)
	if result != "CMZ" {
		log.Fatalf("Test failed! Expected CMZ, got %v", result)
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
		input := elfutils.GetInputByDay("5")
		lines := elfutils.SplitByLine(string(input))
		result := doWork(lines)
		fmt.Printf("result: %v\n", result)
	}
}
