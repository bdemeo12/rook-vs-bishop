package gamefunctions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MoveBishop(t *testing.T) {

	type params struct {
		desc string

		direction int
		diceRoll  int

		expectedRow int
		expectedCol int
	}

	cases := []params{
		{
			desc:        "Move to the top right",
			direction:   0,
			diceRoll:    2,
			expectedRow: 3,
			expectedCol: 4,
		},
		{
			desc:        "handle edge case",
			direction:   3,
			diceRoll:    3,
			expectedRow: 7,
			expectedCol: 0,
		},
	}

	for _, c := range cases {

		game := NewGamePlay()

		game.MoveBishop(c.direction, c.diceRoll)

		assert.Equal(t, c.expectedRow, game.GetBishopRow())
		assert.Equal(t, c.expectedCol, game.GetBishopCol())
	}
}

func Test_MoveRook(t *testing.T) {
	type params struct {
		desc string

		coinToss int
		diceRoll int

		expectedRow int
		expectedCol int
	}

	cases := []params{
		{
			desc:        "Heads, move up 2",
			coinToss:    0,
			diceRoll:    2,
			expectedRow: 5,
			expectedCol: 7,
		},
		{
			desc:        "Tails, move right 2",
			coinToss:    1,
			diceRoll:    2,
			expectedRow: 7,
			expectedCol: 1,
		},
	}

	for _, c := range cases {
		game := NewGamePlay()

		game.MoveRook(c.coinToss, c.diceRoll)

		assert.Equal(t, c.expectedRow, game.GetRookRow())
		assert.Equal(t, c.expectedCol, game.GetRookCol())
	}
}
