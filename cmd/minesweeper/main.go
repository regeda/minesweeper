package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/regeda/minesweeper"
)

func main() {
	var rows, cols int
	fmt.Print("Enter [rows cols]: ")
	if _, err := fmt.Fscan(os.Stdin, &rows, &cols); err != nil {
		panic(err.Error())
	}
	rand.Seed(time.Now().UnixNano())
	m := minesweeper.GenerateMatrix(rows, cols, 0.3)
	g := minesweeper.New(m)
	printMatrix(m)
	var (
		ch   string
		i, j int
	)
	for {
		fmt.Print("Go [F|f|U x y]: ")
		if _, err := fmt.Fscan(os.Stdin, &ch, &i, &j); err != nil {
			panic(err.Error())
		}
		fmt.Println(ch, i, j)
		switch ch {
		case "F":
			m[i][j].Flag(true)
			printMatrix(m)
		case "f":
			m[i][j].Flag(false)
			printMatrix(m)
		case "U":
			left, ok := g.Unfold(i, j)
			printMatrix(m)
			if !ok {
				fmt.Println("GAME OVER")
				return
			}
			if left == 0 {
				fmt.Println("WIN")
				return
			}
			fmt.Println("left cells: ", left)

		}
	}
}

func printMatrix(m minesweeper.Matrix) {
	for _, r := range m {
		for _, c := range r {
			if c.Unfolded() {
				fmt.Print(c.Bombs())
			} else if c.Flagged() {
				fmt.Print("F")
			} else {
				fmt.Print("#")
			}
			fmt.Print(" ")
		}
		fmt.Printf("\n")
	}
}
