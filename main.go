package main

import (
	"fmt"

	gamefunctions "github.com/bdemeo12/rook-vs-bishop/game-functions"
)

const max_rounds = 15

func main() {
	rookGamePlay()
}

func rookGamePlay() {
	fmt.Println("Starting Game!")

	gameplay := gamefunctions.NewGamePlay()
	gameplay.PrintBoard()

	for i := 1; i <= max_rounds; i++ {
		fmt.Printf("Round %d\n", i)

		coinTossResult := gameplay.CoinToss()
		diceRollResult := gameplay.DiceRoll()

		if gameplay.MoveRook(coinTossResult, diceRollResult) {
			gameplay.PrintBoard()

			fmt.Println("!!! WINNER !!!")
			fmt.Println("We have Captured the Bishop!")

			return
		}

		gameplay.PrintBoard()
	}

	fmt.Println("You Lose!")
}
