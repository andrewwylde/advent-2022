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

func doWork(lines []string) (result int) {
	// fmt.Printf("doWork!\n")

	result += (len(lines)*2 + (len(lines[0])-2)*2)

	for row, treeLine := range lines {
		if row == 0 || row == len(lines)-1 {
			continue
		}
		trees := strings.Split(treeLine, "")
		for col, treeHeight := range trees {
			treeHeightInt, _ := strconv.Atoi(treeHeight)
			aboveTrees := []int{}
			belowTrees := []int{}
			leftTrees := []int{}
			rightTrees := []int{}
			if col == 0 || col == len(trees)-1 {
				continue
			}
			// fmt.Printf("treeHeight: %v\n", treeHeight)

			for compareRow, compareLine := range lines {
				if compareRow == row {
					continue
				}
				compareVal, _ := strconv.Atoi(strings.Split(compareLine, "")[col])
				if compareRow < row {
					aboveTrees = append(aboveTrees, compareVal)
				}
				if compareRow > row {
					belowTrees = append(belowTrees, compareVal)
				}
			}

			for _, compareTree := range trees[0:col] {
				compareTreeInt, _ := strconv.Atoi(compareTree)
				leftTrees = append(leftTrees, compareTreeInt)
			}
			for _, compareTree := range trees[col+1:] {
				compareTreeInt, _ := strconv.Atoi(compareTree)
				rightTrees = append(rightTrees, compareTreeInt)
			}

			// fmt.Printf("aboveTrees: %v\n", aboveTrees)
			// fmt.Printf("belowTrees: %v\n", belowTrees)
			// fmt.Printf("leftTrees: %v\n", leftTrees)
			// fmt.Printf("rightTrees: %v\n", rightTrees)
			visible := false
			for _, treeList := range [][]int{aboveTrees, belowTrees, leftTrees, rightTrees} {
				if visible == true {
					break
				}
				isTallestTree := true
				for _, v := range treeList {
					if !isTallestTree {
						break
					}
					if v >= treeHeightInt {
						isTallestTree = false
					}
				}
				if isTallestTree {
					// fmt.Printf("%v is tallest in list %+v\n", treeHeightInt, treeList)
					visible = true
				}
			}
			if visible {
				result++
			}
			// fmt.Printf("\n\n-----------------------\n\n")
		}
	}

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
