package main

import (
	"fmt"
	"log"
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

func main() {
	fmt.Println("Ho ho ho")
	d, _ := os.Getwd()
	p := path.Join(d, "2", "input.txt")
	pairs := generatePairs(p)
	var score = 0
	for _, v := range pairs {
		score += getScore(v)
	}
	fmt.Printf("score: %v\n", score)
	test()
}

func test() {
	fmt.Println("-----------------")
	res1 := getScore([]string{"A", "Y"})
	fmt.Printf("res1: %v\n", res1)
	res2 := getScore([]string{"B", "X"})
	fmt.Printf("res2: %v\n", res2)
	res3 := getScore([]string{"C", "Z"})
	fmt.Printf("res3: %v\n", res3)
}

func getScore(match []string) (points int) {
	elfPick, desiredResult := match[0], match[1]
	player := getPlayerChoice(elfPick, desiredResult)
	x := baseScore(player)
	points += x
	x1 := wld(elfPick, player)
	points += x1
	fmt.Printf("base: %v\n", x)
	fmt.Printf("wld: %v\n", x1)
	return
}

func getPlayerChoice(elfPick string, desired string) (choice string) {
	switch desired {
	case "X":
		switch elfPick {
		case "A":
			choice = "C"
		case "B":
			choice = "A"
		case "C":
			choice = "B"
		}
	case "Y":
		choice = elfPick
	case "Z":
		switch elfPick {
		case "A":
			choice = "B"
		case "B":
			choice = "C"
		case "C":
			choice = "A"
		}
	}
	return
}

func baseScore(str string) (points int) {
	switch str {
	case "A":
		points = 1
	case "B":
		points = 2
	case "C":
		points = 3
	}
	return
}

func wld(elfPick string, playerPick string) (playerScore int) {
	switch playerPick {
	case "A":
		if elfPick == "A" {
			playerScore = 3
		}
		if elfPick == "B" {
			playerScore = 0
		}
		if elfPick == "C" {
			playerScore = 6
		}
	case "B":
		if elfPick == "A" {
			playerScore = 6
		}
		if elfPick == "B" {
			playerScore = 3
		}
		if elfPick == "C" {
			playerScore = 0
		}
	case "C":
		if elfPick == "A" {
			playerScore = 0
		}
		if elfPick == "B" {
			playerScore = 6
		}
		if elfPick == "C" {
			playerScore = 3
		}
	}
	return
}

func generatePairs(filestr string) (arr [][]string) {
	f, err := os.ReadFile(filestr)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	s := strings.Split(string(f), "\n")
	for _, v := range s {
		if v != "" {
			vs := strings.Split(v, " ")
			arr = append(arr, vs)
		}
	}
	return
}

func readLine(line string) (opponent string, player string) {
	values := strings.Split(line, "")
	opponent = values[0]
	player = values[1]
	return
}
