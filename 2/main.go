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
}

func baseScore(str string) (points int) {
	switch str {
	case "X":
		points = 1
	case "Y":
		points = 2
	case "Z":
		points = 3
	}
	return
}

func getScore(match []string) (points int) {
	elfPick, playerPick := match[0], match[1]
	x := baseScore(playerPick)
	points += x
	fmt.Printf("x: %v\n", x)
	x1 := wld(elfPick, playerPick)
	points += x1
	fmt.Printf("x1: %v\n", x1)
	return
}

func wld(elfPick string, playerPick string) (playerScore int) {
	switch playerPick {
	case "X":
		if elfPick == "A" {
			playerScore = 3
		}
		if elfPick == "B" {
			playerScore = 0
		}
		if elfPick == "C" {
			playerScore = 6
		}
	case "Y":
		if elfPick == "A" {
			playerScore = 6
		}
		if elfPick == "B" {
			playerScore = 3
		}
		if elfPick == "C" {
			playerScore = 0
		}
	case "Z":
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
