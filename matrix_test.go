package minesweeper_test

import (
	"testing"

	"github.com/regeda/minesweeper"
	"github.com/stretchr/testify/assert"
)

func TestMatrix_GenerateMatrix(t *testing.T) {
	for i := 0; i < 1000; i++ {
		m := minesweeper.GenerateMatrix(10, 10, .3)
		assert.Equal(t, minesweeper.CellsStat{
			Bombs: 30,
			Free:  70,
		}, m.Stat())
	}
}
