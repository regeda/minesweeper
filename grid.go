package minesweeper

import (
	"math"
	"math/rand"
)

// Grid is playground.
type Grid [][]Cell

func newGrid(rows, cols int) Grid {
	m := make(Grid, rows)
	for i := range m {
		m[i] = make([]Cell, cols)
	}
	return m
}

// GenerateGrid creates a matrix with randomly distributed bombs.
// rand.Seed should be run manually before GenerateGrid call.
func GenerateGrid(rows, cols int, diffc float64) Grid {
	cells := rows * cols
	bombs := int(math.Ceil(float64(cells) * diffc))
	m := newGrid(rows, cols)
	var offset int
	for bombs > 0 {
		offset += rand.Intn(rand.Intn(cells-offset-bombs) + 1)
		i := offset / cols
		j := offset % cols
		m[i][j] = Bomb
		bombs--
		offset++
	}
	return m
}

// Stat returns info about cells.
func (m Grid) Stat() CellsStat {
	var s CellsStat
	for _, r := range m {
		for _, c := range r {
			if c.IsBomb() {
				s.Bombs++
			} else {
				s.Free++
			}
		}
	}
	return s
}

// CellsStat contains info about cells.
type CellsStat struct {
	Bombs, Free int
}
