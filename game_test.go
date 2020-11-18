package minesweeper_test

import (
	"testing"

	"github.com/regeda/minesweeper"
	"github.com/stretchr/testify/assert"
)

func assertEqualMatrix(t *testing.T, expected, actual minesweeper.Grid) {
	for i, r := range expected {
		for j, c := range r {
			if c != actual[i][j] {
				t.Helper()
				t.Errorf("actual cell %+v is not equal to expected cell %+v at [%d, %d]", actual[i][j], c, i, j)
			}
		}
	}
}

func TestGame_Unfold(t *testing.T) {
	t.Run("game over if bomb pressed", func(t *testing.T) {
		m := minesweeper.Grid{
			{0, minesweeper.Bomb, 0},
		}

		g := minesweeper.New(m)

		_, ok := g.Unfold(0, 1)

		assert.False(t, ok)
		assertEqualMatrix(t, minesweeper.Grid{
			{0, minesweeper.Bomb + minesweeper.Unfolded, 0},
		}, m)
	})

	t.Run("unfold empty cells", func(t *testing.T) {
		cases := []struct {
			name        string
			mtest, mexp minesweeper.Grid
			i, j        int
			left        int
		}{
			{
				name: "suggest number up-left",
				mtest: minesweeper.Grid{
					{0, minesweeper.Bomb},
					{minesweeper.Bomb, minesweeper.Bomb},
				},
				mexp: minesweeper.Grid{
					{3 + minesweeper.Unfolded, minesweeper.Bomb},
					{minesweeper.Bomb, minesweeper.Bomb},
				},
				i: 0, j: 0,
				left: 0,
			},
			{
				name: "suggest number up-right",
				mtest: minesweeper.Grid{
					{minesweeper.Bomb, 0},
					{minesweeper.Bomb, minesweeper.Bomb},
				},
				mexp: minesweeper.Grid{
					{minesweeper.Bomb, 3 + minesweeper.Unfolded},
					{minesweeper.Bomb, minesweeper.Bomb},
				},
				i: 0, j: 1,
				left: 0,
			},
			{
				name: "suggest number down-left",
				mtest: minesweeper.Grid{
					{minesweeper.Bomb, minesweeper.Bomb},
					{0, minesweeper.Bomb},
				},
				mexp: minesweeper.Grid{
					{minesweeper.Bomb, minesweeper.Bomb},
					{3 + minesweeper.Unfolded, minesweeper.Bomb},
				},
				i: 1, j: 0,
				left: 0,
			},
			{
				name: "suggest number down-right",
				mtest: minesweeper.Grid{
					{minesweeper.Bomb, minesweeper.Bomb},
					{minesweeper.Bomb, 0},
				},
				mexp: minesweeper.Grid{
					{minesweeper.Bomb, minesweeper.Bomb},
					{minesweeper.Bomb, 3 + minesweeper.Unfolded},
				},
				i: 1, j: 1,
				left: 0,
			},
			{
				name: "unfold empty cells",
				mtest: minesweeper.Grid{
					{0, 0, minesweeper.Bomb},
					{0, 0, minesweeper.Bomb},
					{minesweeper.Bomb, minesweeper.Bomb, minesweeper.Bomb},
				},
				mexp: minesweeper.Grid{
					{minesweeper.Unfolded, 2 + minesweeper.Unfolded, minesweeper.Bomb},
					{2 + minesweeper.Unfolded, 5 + minesweeper.Unfolded, minesweeper.Bomb},
					{minesweeper.Bomb, minesweeper.Bomb, minesweeper.Bomb},
				},
				i: 0, j: 0,
				left: 0,
			},
			{
				name: "unfold closest empty cells",
				mtest: minesweeper.Grid{
					{0, 0, minesweeper.Bomb},
					{0, 0, minesweeper.Bomb},
					{0, 0, minesweeper.Bomb},
				},
				mexp: minesweeper.Grid{
					{minesweeper.Unfolded, 2 + minesweeper.Unfolded, minesweeper.Bomb},
					{minesweeper.Unfolded, 3 + minesweeper.Unfolded, minesweeper.Bomb},
					{minesweeper.Unfolded, 2 + minesweeper.Unfolded, minesweeper.Bomb},
				},
				i: 1, j: 0,
				left: 0,
			},
			{
				name: "suggest number and continue game",
				mtest: minesweeper.Grid{
					{0, 0, minesweeper.Bomb},
				},
				mexp: minesweeper.Grid{
					{0, 1 + minesweeper.Unfolded, minesweeper.Bomb},
				},
				i: 0, j: 1,
				left: 1,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				g := minesweeper.New(c.mtest)

				left, ok := g.Unfold(c.i, c.j)
				assert.True(t, ok, "continue game")
				assert.Equal(t, c.left, left, "left cells")
				assertEqualMatrix(t, c.mexp, c.mtest)
			})
		}
	})

	t.Run("flagged cells", func(t *testing.T) {
		mtest := minesweeper.Grid{
			{0, 0, minesweeper.Bomb},
		}

		g := minesweeper.New(mtest)

		// set flag
		mtest[0][1].Flag(true)
		assert.True(t, mtest[0][1].Flagged())
		assertEqualMatrix(t, minesweeper.Grid{
			{0, minesweeper.Flagged, minesweeper.Bomb},
		}, mtest)

		// unfold empty cell, flagged cell is still unfolded
		left, ok := g.Unfold(0, 0)
		assert.True(t, ok)
		assert.Equal(t, 1, left)
		assertEqualMatrix(t, minesweeper.Grid{
			{minesweeper.Unfolded, minesweeper.Flagged, minesweeper.Bomb},
		}, mtest)

		// cannot set a flag on unfolded cell
		mtest[0][0].Flag(true)
		assertEqualMatrix(t, minesweeper.Grid{
			{minesweeper.Unfolded, minesweeper.Flagged, minesweeper.Bomb},
		}, mtest)

		// nothing changed if unfold flagged cell
		left, ok = g.Unfold(0, 1)
		assert.True(t, ok)
		assert.Equal(t, 1, left)
		assertEqualMatrix(t, minesweeper.Grid{
			{minesweeper.Unfolded, minesweeper.Flagged, minesweeper.Bomb},
		}, mtest)

		// unset flag
		mtest[0][1].Flag(false)
		assert.False(t, mtest[0][1].Flagged())
		assertEqualMatrix(t, minesweeper.Grid{
			{minesweeper.Unfolded, 0, minesweeper.Bomb},
		}, mtest)

		// unfold empty cell
		left, ok = g.Unfold(0, 1)
		assert.True(t, ok)
		assert.Equal(t, 0, left)
		assert.Equal(t, byte(1), mtest[0][1].Bombs())
		assertEqualMatrix(t, minesweeper.Grid{
			{minesweeper.Unfolded, 1 + minesweeper.Unfolded, minesweeper.Bomb},
		}, mtest)

	})

	t.Run("double press", func(t *testing.T) {
		mexp := minesweeper.Grid{
			{1 + minesweeper.Unfolded, minesweeper.Bomb},
		}

		mtest := minesweeper.Grid{
			{0, minesweeper.Bomb},
		}

		g := minesweeper.New(mtest)

		_, ok := g.Unfold(0, 0)
		assert.True(t, ok)
		assertEqualMatrix(t, mexp, mtest)

		_, ok = g.Unfold(0, 0)
		assert.True(t, ok)
		assertEqualMatrix(t, mexp, mtest)
	})
}
