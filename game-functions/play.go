package gamefunctions

import (
	"fmt"
	"math/rand"
)

const (
	board_size = 8
)

func NewGamePlay() GamePlay {

	// Create board on start
	board := make([][]string, 8)
	for i := range board {
		board[i] = make([]string, 8)
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board[i][j] = "."
		}
	}

	board[7][7] = "R"
	board[5][2] = "B"

	return &gamePlay{
		board:   board,
		rookRow: 7,
		rookCol: 7,
	}
}

type GamePlay interface {
	PrintBoard()
	CoinToss() int
	DiceRoll() int
	MoveRook(coinToss, diceRoll int) bool
}

type gamePlay struct {
	board   [][]string
	rookRow int
	rookCol int
}

func (g *gamePlay) PrintBoard() {
	files := "abcdefgh" // columns
	ranks := "87654321" // rows

	fmt.Print("    ")
	for _, file := range files {
		fmt.Printf("%2s ", string(file))
	}
	fmt.Println()

	for i, rank := range ranks {
		fmt.Printf(" %2s ", string(rank))
		for j := 0; j < 8; j++ {
			fmt.Printf(" %s ", g.board[i][j])
		}
		fmt.Println("")
	}

}

func (g *gamePlay) CoinToss() int {
	coinTossResult := rand.Intn(2) // 0 is heads, 1 is tails

	if coinTossResult == 0 {
		fmt.Println("The coin toss result is: heads")
	} else {
		fmt.Println("The coin toss result is: tails")
	}

	return coinTossResult
}

func (g *gamePlay) DiceRoll() int {
	diceRollResult1 := rand.Intn(6) + 1
	diceRollResult2 := rand.Intn(6) + 1

	fmt.Printf("The first dice roll result is: %d\n", diceRollResult1)
	fmt.Printf("The second dice roll result is: %d\n", diceRollResult2)

	fmt.Printf("The total dice roll result is: %d \n", diceRollResult1+diceRollResult2)

	return diceRollResult1 + diceRollResult2
}

func (g *gamePlay) MoveRook(coinToss, diceRoll int) bool {

	// Clean up old R
	g.board[g.rookRow][g.rookCol] = "."

	// Find new row,col of R
	if coinToss == 0 { // Heads - move up
		g.rookRow -= diceRoll

		fmt.Println("in coin toss: ", g.rookRow)

		// hangle edge cases
		for g.rookRow < 0 {
			g.rookRow = (g.rookRow + 8) % 8
		}
	} else if coinToss == 1 { // tails - move right
		g.rookCol += diceRoll

		// hangle edge cases
		if g.rookCol >= 8 {
			g.rookCol = g.rookCol % 8
		}
	}

	// Check to see if we have captured the Bishop
	if g.board[g.rookRow][g.rookCol] == "B" {
		g.board[g.rookRow][g.rookCol] = "R"
		return true
	}

	// update R
	g.board[g.rookRow][g.rookCol] = "R"

	return false
}

// move bishop // extra
