package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readHands(path string) ([]int, []int) {
	var p1, p2 []int
	bytes, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(bytes), "\n")
	player := 1
	for i := 1; i < len(lines); i++ {
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
	winner := 0
	switch {
	case d1 > d2:
		p1, p2 = append(p1[1:len(p1)], p1[0], p2[0]), p2[1:]
		winner = -1
	case d1 < d2:
		p2, p1 = append(p2[1:len(p2)], p2[0], p1[0]), p1[1:]
		winner = 1
	}
	return p1, p2, winner
}

func handKey(p []int) string {
	sAry := make([]string, len(p))
	for i := 0; i<len(p);i++ {
		sAry[i] = strconv.FormatInt(int64(p[i]), 10)
	}
	return strings.Join(sAry, ",")
}

// // Concatenates all digits in an integer array irrespective of the sizes of
// // individual integers.
// func handKey(p []int) int {
// 	key, tens := 0, 1
// 	p2 := make([]int, len(p))
// 	copy(p2, p)
// 	digits := []int{}
// 	for i:=0;i<len(p2);i++{
// 		if p2[i] > 9 {
// 			dig := []int{}
// 			for p2[i] > 9 {
// 				dig = append(dig, p2[i]%10)
// 				p2[i] = p2[i]/10
// 			}
// 			dig = append(dig, p2[i])
// 			for j:=len(dig)-1;j>-1;j--{
// 				digits = append(digits, dig[j])
// 			}
// 		} else {
// 			digits = append(digits, p[i])
// 		}
// 	}
// 	for i := len(digits)-1; i > -1; i-- {
// 		key += digits[i] * tens
// 		tens *= 10
// 	}
// 	return key
// }

func recursiveRounds(p1 []int, p2 []int) ([]int, []int, int) {
	winner := 0
	p3, p4 := make([]int, len(p1)), make([]int, len(p2))
	copy(p3, p1)
	copy(p4, p2)
	p3Hist, p4Hist := make(map[string]bool), make(map[string]bool)
	for ((len(p3) > 0) && (len(p4) > 0)) {
		p3Key, p4Key := handKey(p3), handKey(p4)
		if (p3Hist[p3Key]) || (p4Hist[p4Key]) {
			return p3, p4, -1
		}
		p3Hist[p3Key], p4Hist[p4Key] = true, true
		if len(p3) > p3[0] && len(p4) > p4[0] {
			_, _, winner = recursiveRounds(p3[1:p3[0]+1], p4[1:p4[0]+1])
			switch winner {
			case -1:
				p3, p4 = append(p3[1:], p3[0], p4[0]), p4[1:]
			case 1:
				p4, p3 = append(p4[1:], p4[0], p3[0]), p3[1:]
			}
		} else {
			p3, p4, winner = playRound(p3, p4)
		}
	}
	return p3, p4, winner
}

func playGame(p1 []int, p2 []int, gametype string) int {
	if gametype == "normal" {
		for len(p1) > 0 && len(p2) > 0 {
			p1, p2, _ = playRound(p1, p2)
		}
	} else if gametype == "recursive" {
		for len(p1) > 0 && len(p2) > 0 {
			p1, p2, _ = recursiveRounds(p1, p2)
		}
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
	for i := 0; i < len(p); i++ {
		score = score + (p[i] * (len(p) - i))
	}
	return score

}

func test() {
	s1 := []int{1,2,3,4,4,5, 70, 70, 82}
	fmt.Println(s1[1:3+1])
	fmt.Println(handKey(s1))
}

func main() {
	// path := "22_small.txt"
	// path := "22_small2.txt"
	path := "22.txt"
	p1, p2 := readHands(path)
	fmt.Println("Part 1:", playGame(p1, p2, "normal"))
	fmt.Println("Part 2:", playGame(p1, p2, "recursive"))
}
