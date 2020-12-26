package main

import (
	"fmt"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2020/util"
)

const testData = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

const testData2 = `Player 1:
43
19

Player 2:
2
29
14`

func main() {
	input := strings.Split(util.ReadEntireFile("input.txt"), "\n\n")
	// input = strings.Split(testData, "\n\n")
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))

	startTime = time.Now()
	res = part2(input) // strings.Split(testData, "\n\n"))
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

func parseDecks(input []string) (out [][]int) {
	for _, player := range input {
		out = append(out, util.GetInts(strings.Split(player, "\n")[1:]))
	}
	return out
}

func calculateDeckWorth(deck []int) int {
	out := 0
	data := make([]int, len(deck))
	copy(data, deck)
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

	for i, v := range data {
		out += (v * (i + 1))
	}
	return out
}

func part1(input []string) string {
	decks := parseDecks(input)
	winner := -1
	for {
		player1, player2 := decks[0], decks[1]
		var topCard1, topCard2 int
		topCard1, player1 = player1[0], player1[1:]
		topCard2, player2 = player2[0], player2[1:]
		if topCard1 > topCard2 {
			player1 = append(player1, topCard1, topCard2)
		} else {
			player2 = append(player2, topCard2, topCard1)
		}
		decks[0] = player1
		decks[1] = player2

		if len(player1) == 0 {
			winner = 1
			break
		} else if len(player2) == 0 {
			winner = 0
			break
		}

	}

	return fmt.Sprint(calculateDeckWorth(decks[winner]))
}

func copyIntSlice(a []int) []int {
	out := make([]int, len(a))
	copy(out, a)
	return out
}

func cmpIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func debugPrint(a ...interface{}) {
	if debugGame {
		fmt.Println(a...)
	}
}

func debugPrintf(f string, a ...interface{}) {
	if debugGame {
		fmt.Printf(f, a...)
	}
}

var lastGame = 1

const debugGame = false

func playRecursive(decks [][]int, gameNo int) (winner int) {
	debugPrintf("=== Game %d ===\n", gameNo)
	round := 0
	defer func() {
		debugPrintf("The winner of game %d is player %d\n", gameNo, winner)
		debugPrint("Anyway... Back to game", gameNo-1)
	}()

	prevGames := [][][]int{}

	// Draw cards
	for {
		// var pg1, pg2 []int
		round++
		debugPrintf("\n--Round %d (game %d)--\n", round, gameNo)
		player1, player2 := decks[0], decks[1]

		debugPrintf("Player 1's deck: %v\n", player1)
		debugPrintf("Player 2's deck: %v\n", player2)
		for _, prevDeckPair := range prevGames {
			prevPlayer1, prevPlayer2 := prevDeckPair[0], prevDeckPair[1]
			if cmpIntSlice(prevPlayer1, player1) || cmpIntSlice(prevPlayer2, player2) {
				debugPrintf("Player %d instantly wins the game\n", 1)
				return 1 // Game instantly wins for player 1
			}
		}

		prevGames = append(prevGames, [][]int{copyIntSlice(decks[0]), copyIntSlice(decks[1])})

		var topCard1, topCard2 int
		topCard1, player1 = player1[0], player1[1:]
		topCard2, player2 = player2[0], player2[1:]
		debugPrintf("Player 1 plays: %d\nPlayer 2 plays: %d\n", topCard1, topCard2)

		if len(player1) >= topCard1 && len(player2) >= topCard2 {
			debugPrint("Playing a sub-game to determine the winner...")
			var rP1, rP2 []int
			for i := 0; i < topCard1; i++ {
				rP1 = append(rP1, player1[i])
			}
			for i := 0; i < topCard2; i++ {
				rP2 = append(rP2, player2[i])
			}
			lastGame++
			winner := playRecursive([][]int{rP1, rP2}, gameNo+1)

			if winner == 1 {
				debugPrintf("Player 1 wins round %d of game %d\n", round, gameNo)
				// Player 1 won
				player1 = append(player1, topCard1, topCard2)
			} else if winner == 2 {
				debugPrintf("Player 2 wins round %d of game %d\n", round, gameNo)
				player2 = append(player2, topCard2, topCard1)
			}

		} else if topCard1 > topCard2 {
			debugPrintf("Player 1 wins round %d of game %d\n", round, gameNo)
			// Player 1 won
			player1 = append(player1, topCard1, topCard2)
		} else if topCard2 > topCard1 {
			debugPrintf("Player 2 wins round %d of game %d\n", round, gameNo)
			player2 = append(player2, topCard2, topCard1)
		}

		decks[0] = copyIntSlice(player1)
		decks[1] = copyIntSlice(player2)

		// prevGames = append(prevGames, [][]int{pg1, pg2})

		if len(player1) == 0 {
			// player 2 wins
			winner = 2
			return 2
		} else if len(player2) == 0 {
			winner = 1
			return 1
		}
	}
}

func part2(input []string) string {
	decks := parseDecks(input)
	winner := playRecursive(decks, 1)

	return fmt.Sprint(calculateDeckWorth(decks[winner-1]))
}
