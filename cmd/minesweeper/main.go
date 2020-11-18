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
		fmt.Println("failed to read rows & cols:", err)
		return
	}
	rand.Seed(time.Now().UnixNano())
	m := minesweeper.GenerateGrid(rows, cols, 0.3)
	g := minesweeper.New(m)
	printGrid(m)
	var (
		cmd  string
		i, j int
	)
	for {
		fmt.Print("Go [x y [f-flag|v-unflag|u-unfold]]: ")
		if _, err := fmt.Fscan(os.Stdin, &i, &j, &cmd); err != nil {
			fmt.Println("failed to read an action:", err)
			continue
		}
		switch cmd {
		case "f":
			m[i][j].Flag(true)
			printGrid(m)
		case "v":
			m[i][j].Flag(false)
			printGrid(m)
		case "u":
			left, ok := g.Unfold(i, j)
			printGrid(m)
			if !ok {
				fmt.Println("GAME OVER")
				return
			}
			if left == 0 {
				fmt.Println("WIN")
				return
			}
			fmt.Println("left cells:", left)
		default:
			fmt.Println("Unknown command:", cmd)
		}
	}
}

func printGrid(m minesweeper.Grid) {
	for _, r := range m {
		for _, c := range r {
			if c.Unfolded() {
				if c.IsBomb() {
					fmt.Print("x")
				} else {
					fmt.Print(c.Bombs())
				}
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
