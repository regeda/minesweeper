# Minesweeper

Golang implementation of Minesweeper game API.

```go
cols := 5
rows := 5
difficulty := 0.3 // how much mines should be generated ceil(cols*rows*difficulty)

m := minesweeper.GenerateGrid(rows, cols, difficulty)
g := minesweeper.New(m)

for {
	// ... read i,j
	// you can set a flag
	m[i][j].Flag(true)
	// or unset
	m[i][j].Flag(false)
	// or unfold a cell
	left, ok := g.Unfold(i, j)
	if !ok {
		// game over, you touched a mine
	} else if left == 0 {
		// you won, all free cells unfolded
	}
}
```

Or play in command line:
```
go run cmd/minesweeper/main.go
```
