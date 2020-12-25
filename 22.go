package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

type hand []int

func readHands(path string) ([]int, []int) {
	var p1, p2 []int
	bytes, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(bytes), "\n")
	player := 1
	for i:=1;i<len(lines);i++{
		if lines[i] != "" {
			if lines[i] == "Player 2:" {
				player = 2
			} else if player == 1 {
				n, _ := strconv.Atoi(lines[i])
				p1 = append(p1, n)
			} else if player == 2 {
				n, _ := strconv.Atoi(lines[i])
				p2 = append(p2, n)
			}
		}
	}
	return p1, p2
}

func playRound(p1 []int, p2 []int) ([]int, []int, int) {
	d1, d2 := p1[0], p2[0]
	fmt.Println("play", d1, d2)
	winner := 0
	switch {
	case d1 > d2:
		// p1 = append(p1[:len(p1))
		p1 = append(p1[1:len(p1)], p1[0], p2[0])
		p2 = p2[1:]
		winner = -1
		// p2
	case d1 < d2:
		p2 = append(p2[1:len(p2)], p2[0], p1[0])
		p1 = p1[1:]
		winner = 1
	}
	return p1, p2, winner
}

func playGame(p1 []int, p2 []int) int {
	round := 1
	for len(p1) > 0 && len(p2) > 0 {
		fmt.Println("Round", round)
		fmt.Println(p1, p2)
		p1, p2, _ = playRound(p1, p2)
		round += 1
	}
	winningScore := 0
	if len(p1) > 0 {
		winningScore = scoreWinner(p1)
	} else if len(p2) > 0 {
		winningScore = scoreWinner(p2)
	}
	return winningScore
}

func scoreWinner(p []int) int {
	score := 0
	for i:= 0; i<len(p); i++{
		score = score + (p[i] * (len(p) - i))
	}
	return score

}

func main() {
	// path := "22_small.txt"
	path := "22.txt"
	p1, p2 := readHands(path)
	fmt.Println("Part 1:", playGame(p1, p2))
}
