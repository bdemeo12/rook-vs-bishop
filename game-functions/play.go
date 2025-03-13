package gamefunctions

import (
	"fmt"
	"math/rand"
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
		board:     board,
		rookRow:   7,
		rookCol:   7,
		bishopRow: 5,
		bishopCol: 2,
	}
}

type GamePlay interface {
	PrintBoard()
	CoinToss() int
	DiceRoll() int
	PickDirection() int
	MoveRook(coinToss, diceRoll int) bool
	MoveBishop(direction, diceRoll int) bool
	GetBishopRow() int
	GetBishopCol() int
	GetRookRow() int
	GetRookCol() int
}

type gamePlay struct {
	board [][]string

	rookRow int
	rookCol int

	bishopRow int
	bishopCol int
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

	fmt.Printf("The total dice roll is %d\n", diceRollResult1+diceRollResult2)

	return diceRollResult1 + diceRollResult2
}

func (g *gamePlay) PickDirection() int {
	dir := rand.Intn(4)
	fmt.Printf("The Direction is: %d\n", dir)

	return dir
}

func (g *gamePlay) MoveRook(coinToss, diceRoll int) bool {

	// Clean up old R
	g.board[g.rookRow][g.rookCol] = "."

	// Find new row,col of R
	if coinToss == 0 { // Heads - move up
		g.rookRow -= diceRoll

		// Handle edge cases
		for g.rookRow < 0 {
			g.rookRow = (g.rookRow + 8) % 8
		}
	} else if coinToss == 1 { // Tails - move right
		g.rookCol += diceRoll

		// Handle edge cases
		if g.rookCol >= 8 {
			g.rookCol = g.rookCol % 8
		}
	}

	// Print new coordinates
	fmt.Printf("New position: %c%d\n", 'a'+g.rookCol, 8-g.rookRow)

	// Check to see if we have captured the Bishop
	if g.board[g.rookRow][g.rookCol] == "B" {
		g.board[g.rookRow][g.rookCol] = "R"
		return true
	}

	// update R
	g.board[g.rookRow][g.rookCol] = "R"

	return false
}

func (g *gamePlay) MoveBishop(direction, diceRoll int) bool {
	// Clean up old B
	g.board[g.bishopRow][g.bishopCol] = "."

	// Find new row/col of B
	for i := 0; i < diceRoll; i++ {
		switch direction {
		case 0: // Top right
			g.bishopRow--
			g.bishopCol++

			if g.bishopRow < 0 || g.bishopCol >= 8 { // If we reach the edge of the board, switch direction to the other side
				g.bishopRow++
				g.bishopCol--
				direction = 3
			}
		case 1: // Top left
			g.bishopRow--
			g.bishopCol--

			if g.bishopRow < 0 || g.bishopCol < 0 {
				g.bishopRow++
				g.bishopCol++
				direction = 2
			}
		case 2: // Bottom right
			g.bishopRow++
			g.bishopCol++

			if g.bishopRow >= 8 || g.bishopCol >= 8 {
				g.bishopRow--
				g.bishopCol--
				direction = 1
			}
		case 3: // Bottom left
			g.bishopRow++
			g.bishopCol--

			if g.bishopRow >= 8 || g.bishopCol < 0 {
				g.bishopRow--
				g.bishopCol++
				direction = 0
			}
		}
	}
	// Print new coordinates
	fmt.Printf("New position: %c%d\n", 'a'+g.bishopCol, 8-g.bishopRow)

	// Check to see if we have captured the rook
	if g.board[g.bishopRow][g.bishopCol] == "R" {
		g.board[g.bishopRow][g.bishopCol] = "B"
		return true
	}
	//Update B
	g.board[g.bishopRow][g.bishopCol] = "B"

	return false
}

// Helper functions for testing

func (g *gamePlay) GetBishopRow() int {
	return g.bishopRow
}

func (g *gamePlay) GetBishopCol() int {
	return g.bishopCol
}

func (g *gamePlay) GetRookRow() int {
	return g.rookRow
}

func (g *gamePlay) GetRookCol() int {
	return g.rookCol
}
