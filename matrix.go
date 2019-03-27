package minesweeper

import (
	"math"
	"math/rand"
)

// Matrix is playground.
type Matrix [][]Cell

func newMatrix(rows, cols int) Matrix {
	m := make(Matrix, rows)
	for i := range m {
		m[i] = make([]Cell, cols)
	}
	return m
}

// GenerateMatrix creates a matrix with randomly distributed bombs.
// rand.Seed should be run manually before GenerateMatrix call.
func GenerateMatrix(rows, cols int, diffc float64) Matrix {
	cells := rows * cols
	bombs := int(math.Ceil(float64(cells) * diffc))
	m := newMatrix(rows, cols)
	var offset int
	for bombs > 0 {
		offset += rand.Intn(cells - offset - bombs + 1)
		i := offset / cols
		j := offset % cols
		m[i][j] = Bomb
		bombs--
		offset++
	}
	return m
}

// Stat returns info about cells.
func (m Matrix) Stat() CellsStat {
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
