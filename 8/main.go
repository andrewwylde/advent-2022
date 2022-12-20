package main

import (
	"advent/elfutils"
	"fmt"
	"log"
	"os"
	"strconv"
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

// write a func

func treeScore(treeHeight int, tlrb [][]int) (score int) {
	score = 1

	// for each direction, count the number of trees that are shorter than the current tree
	for _, directionSlice := range tlrb {
		if len(directionSlice) == 0 {
			// at edge
			return 0
		}
		shorterTrees := getShorterTrees(directionSlice, treeHeight)
		// fmt.Printf("trees shorter than %v: %v\n", treeHeight, shorterTrees)
		rowScore := len(shorterTrees)
		if rowScore == 0 {
			continue
		}

		score *= rowScore
	}

	return score
}

func getShorterTrees(directionSlice []int, treeHeight int) []int {
	for i, v2 := range directionSlice {
		if v2 >= treeHeight {
			return directionSlice[:i+1]
		}
	}
	return directionSlice
}

func prepend(arr []int, element int) []int {
	return append([]int{element}, arr...)
}

// generate a struct
type Tree struct {
	height int `json:"height"`
	score  int
	up     []int `json:"up"`
	down   []int `json:"down"`
	left   []int `json:"left"`
	right  []int `json:"right"`
}

func doWork(lines []string) (result int) {
	result = 1

	for row, treeLine := range lines {
		if row == 0 || row == len(lines)-1 {
			continue
		}
		trees := strings.Split(treeLine, "")
		for col, treeHeight := range trees {
			treeHeightInt, _ := strconv.Atoi(treeHeight)
			up := []int{}
			down := []int{}
			left := []int{}
			right := []int{}
			if col == 0 || col == len(trees)-1 {
				continue
			}

			for compareRow, compareLine := range lines {
				if compareRow == row {
					continue
				}
				compareVal, _ := strconv.Atoi(strings.Split(compareLine, "")[col])
				if compareRow < row {
					up = prepend(up, compareVal)
				}
				if compareRow > row {
					down = prepend(down, compareVal)
				}
			}

			for _, compareTree := range trees[0:col] {
				compareTreeInt, _ := strconv.Atoi(compareTree)
				left = prepend(left, compareTreeInt)
			}

			for _, compareTree := range trees[col+1:] {
				compareTreeInt, _ := strconv.Atoi(compareTree)
				right = append(right, compareTreeInt)
			}

			score := treeScore(treeHeightInt, [][]int{up, right, left, down})
			tree := Tree{treeHeightInt, score, up, down, left, right}
			if score > result {
				result = score
				fmt.Printf("tree: %v\n", tree)
				fmt.Printf("score: %v\n", score)
				println("--------------------------------------------------")
			}
		}
	}

	// printshit()

	// fmt.Printf("treeMap: %v\n", treeMap)
	// fmt.Printf("tree: %v\n", tree.height)
	// fmt.Printf("left: %v\n", tree.left)
	// fmt.Printf("right: %v\n", tree.right)
	// fmt.Printf("up: %v\n", tree.up)
	// fmt.Printf("down: %v\n", tree.down)

	return
}

func test(day string) {
	input := elfutils.GetTestInputByDay(day)
	lines := elfutils.SplitByLine(string(input))
	result := doWork(lines)
	x := 8
	if result != x {
		log.Fatalf("Test failed! Expected %v, got %v", x, result)
	} else {
		fmt.Printf("success: %v\n", result)
	}
}

func main() {
	args := os.Args
	isTest := false
	day := args[2]

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
