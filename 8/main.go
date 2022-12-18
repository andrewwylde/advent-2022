package main

import (
	"advent/elfutils"
	// "encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ttacon/chalk"
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

func TreeScore(treeHeight int, tlrb [][]int) (score int) {
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

//

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
	Height int `json:"height"`
	score  int
	Up     []int `json:"up"`
	Down   []int `json:"down"`
	Left   []int `json:"left"`
	Right  []int `json:"right"`
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
					down = append(down, compareVal)
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

			score := TreeScore(treeHeightInt, [][]int{up, right, left, down})
			x := getShorterTrees(up, treeHeightInt)
			x1 := getShorterTrees(down, treeHeightInt)
			x2 := getShorterTrees(left, treeHeightInt)
			x3 := getShorterTrees(right, treeHeightInt)
			// tree := Tree{treeHeightInt, score, x, x1, x2, x3}

			if score > result {
				result = score
				fmt.Printf("x: %v\n", x)
				fmt.Printf("x1: %v\n", x1)
				fmt.Printf("x2: %v\n", x2)
				fmt.Printf("x3: %v\n", x3)
				highlightGrid(lines, row, col)
				// treeJSONIndented, _ := json.Marshal(tree)
				// fmt.Printf("%s%v\n", chalk.Bold.TextStyle("Tree:"), string(treeJSONIndented))
				println("--------------------------------------------------")
			}

			// logging bullshit
			// fmt.Printf("%s%v\n", chalk.Bold.TextStyle("Tree:"), tree)
			// fmt.Printf("%s%v\n", chalk.Bold.TextStyle("Score:"), score)
		}
	}

	return
}

// highlight a string in a grid
func highlightGrid(grid []string, x int, y int) {
	for i, row := range grid {
		for j, col := range row {
			if i == x && j == y {
				fmt.Printf("%s", chalk.Bold.TextStyle(string(col)))
			} else {
				fmt.Printf("%s", string(col))
			}
		}
		fmt.Printf("\n")
	}
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
