package minesweeper

// Game implements game play.
type Game struct {
	m         Grid
	cellsStat CellsStat

	xwalk   xset
	xbuf    []xcell
	pressed int
}

// New creates a new Game.
func New(m Grid) *Game {
	return &Game{
		m:         m,
		cellsStat: m.Stat(),
		xwalk:     make(xset),
		xbuf:      make([]xcell, 0, 8),
	}
}

func (g *Game) cellsLeft() int {
	return g.cellsStat.Free - g.pressed
}

// Unfold touches a cell. It fails if a bomb pressed.
// Returns amount of remained cells.
// The game passed if no cells left.
func (g *Game) Unfold(i, j int) (int, bool) {
	if g.m[i][j].IsBomb() {
		g.m[i][j].unfold()
		return 0, false
	}
	g.unfold(i, j)
	return g.cellsLeft(), true
}

type xcell struct {
	i, j int
}

var xnil = xcell{-1, -1}

type xset map[xcell]struct{}

func (h xset) add(x xcell) {
	h[x] = struct{}{}
}

func (h xset) drop() xcell {
	for x := range h {
		delete(h, x)
		return x
	}
	return xnil
}

func (g *Game) unfold(i, j int) {
	x := xcell{i, j}
	for ; x != xnil; x = g.xwalk.drop() {
		c := &g.m[x.i][x.j]
		if c.any(Unfolded | Flagged) {
			continue
		}
		c.unfold()
		g.pressed++
		bombs := g.suggestBombs(x.i, x.j)
		if bombs == 0 {
			// need to unfold cells around
			for _, x := range g.xbuf {
				g.xwalk.add(x)
			}
		} else {
			c.suggest(bombs)
		}
		g.xbuf = g.xbuf[:0]
	}
}

func (g *Game) hasBomb(i, j int) byte {
	if g.m[i][j].IsBomb() {
		return 1
	}
	g.xbuf = append(g.xbuf, xcell{i, j})
	return 0
}

func (g *Game) suggestBombs(i, j int) (bombs byte) {
	if i-1 >= 0 {
		// up-left
		if j-1 >= 0 {
			bombs += g.hasBomb(i-1, j-1)
		}
		// up
		bombs += g.hasBomb(i-1, j)
		// up-right
		if j+1 < len(g.m[i-1]) {
			bombs += g.hasBomb(i-1, j+1)
		}
	}
	// left
	if j-1 >= 0 {
		bombs += g.hasBomb(i, j-1)
	}
	// right
	if j+1 < len(g.m[i]) {
		bombs += g.hasBomb(i, j+1)
	}
	if i+1 < len(g.m) {
		// down-left
		if j-1 >= 0 {
			bombs += g.hasBomb(i+1, j-1)
		}
		// down
		bombs += g.hasBomb(i+1, j)
		// down-right
		if j+1 < len(g.m[i+1]) {
			bombs += g.hasBomb(i+1, j+1)
		}
	}
	return bombs
}
