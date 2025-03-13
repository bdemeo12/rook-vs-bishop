package main

import (
	"fmt"

	gamefunctions "github.com/bdemeo12/rook-vs-bishop/game-functions"
)

const max_rounds = 15

func main() {
	var game string

	fmt.Print("Would you like to play the rook game, or the bishop game: ")
	fmt.Scanln(&game)

	switch game {
	case "rook":
		rookGamePlay()
	case "bishop":
		bishopGamePlay()
	default:
		fmt.Println("Invalid game choice. Please choose 'rook' or 'bishop'.")
		return
	}
}

func rookGamePlay() {
	fmt.Println("Starting Game!")

	gameplay := gamefunctions.NewGamePlay()
	gameplay.PrintBoard()

	for i := 1; i <= max_rounds; i++ {
		fmt.Printf("Round %d\n\n", i)

		coinTossResult := gameplay.CoinToss()
		diceRollResult := gameplay.DiceRoll()

		if gameplay.MoveRook(coinTossResult, diceRollResult) {
			gameplay.PrintBoard()

			fmt.Println("!!! Rook Wins !!!")
			fmt.Println("We have Captured the Bishop!")

			return
		}

		gameplay.PrintBoard()
	}

	fmt.Println("You Lose!")
}

func bishopGamePlay() {
	fmt.Println("Starting Game!")

	gameplay := gamefunctions.NewGamePlay()
	gameplay.PrintBoard()

	for i := 1; i <= max_rounds; i++ {
		fmt.Printf("Round %d\n\n", i)

		println("Rooks Turn\n")
		coinTossResult := gameplay.CoinToss()
		diceRollResult := gameplay.DiceRoll()

		if gameplay.MoveRook(coinTossResult, diceRollResult) {
			gameplay.PrintBoard()

			fmt.Println("!!! Rook Wins !!!")
			fmt.Println("We have Captured the Bishop!")

			return
		}

		gameplay.PrintBoard()

		println("Bishops Turn\n")

		directionResult := gameplay.PickDirection()
		diceRollResult = gameplay.DiceRoll()

		if gameplay.MoveBishop(directionResult, diceRollResult) {
			gameplay.PrintBoard()

			fmt.Println("!!! Bishop Wins !!!")
			fmt.Println("We have Captured the Rook!")

			return
		}

		gameplay.PrintBoard()
	}

	fmt.Println("You Both Lose!")
}
