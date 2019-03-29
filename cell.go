package minesweeper

// Cell constants
const (
	Bomb     = Cell(9)
	Unfolded = Cell(1 << 4)
	Flagged  = Cell(1 << 5)
)

// Cell contains cell's state.
type Cell byte

func (c Cell) IsBomb() bool {
	return c.Bombs() == byte(Bomb)
}

// Bombs returns bombs suggestions.
func (c Cell) Bombs() byte {
	return byte(c &^ (Flagged | Unfolded))
}

// Unfolded checks that a cell is open.
func (c Cell) Unfolded() bool {
	return c.is(Unfolded)
}

func (c Cell) is(t Cell) bool {
	return c&t == t
}

func (c *Cell) unfold() {
	c.Flag(false)
	*c |= Unfolded
}

func (c *Cell) suggest(bombs byte) {
	*c |= Cell(bombs)
}

// Flagged checks that a cell has a flag.
func (c Cell) Flagged() bool {
	return c.is(Flagged)
}

// Flag sets or unsets a flag on the cell.
func (c *Cell) Flag(enable bool) {
	if c.Unfolded() {
		return
	}
	if enable {
		*c |= Flagged
	} else {
		*c &^= Flagged
	}
}
