package main

import (
	"errors"
	"fmt"
	"time"
	//"bufio"
	//"os"
)

type Grid struct {
	height, width int
	grid          []byte
}

func (g Grid) String() string {
	return string(g.grid)
}

func NewGrid(x, y int) Grid {
	wth := 2*x + 2 // +2 for right column + '\n'
	hgt := 2*y + 1 // +1 for bottom row

	g := make([]byte, wth*hgt)

	for i := 0; i < hgt; i += 2 {
		row0 := i * wth
		row1 := (i + 1) * wth
		for j := 0; j < wth-2; j += 2 {
			g[row0+j], g[row0+j+1] = '+', '-'
			if row1+j+1 <= wth*hgt {
				g[row1+j], g[row1+j+1] = '|', ' '
			}
		}
		g[row0+wth-2], g[row0+wth-1] = '+', '\n'
		if row1+wth < wth*hgt {
			g[row1+wth-2], g[row1+wth-1] = '|', '\n'
		}
	}

	return Grid{
		height: y,
		width:  x,
		grid:   g,
	}
}

func (g Grid) Set(c byte, x, y int) error {
	idx, err := g.cellAt(x, y)
	if err != nil {
		return err
	}
	g.grid[idx] = c
	return nil
}

func (g Grid) cellAt(x, y int) (int, error) {
	woff := g.width*2 + 2 // width offset
	foff := (y*2+1)*woff + x*2 + 1

	if foff > len(g.grid) {
		return 0, errors.New("out of range")
	}

	return (y*2+1)*woff + x*2 + 1, nil
}

func (g Grid) Draw() {
	fmt.Print("\033[H\033[2J")         // Clear screen
	fmt.Print("\x0c", g, "\n")         // Print frame
	time.Sleep(250 * time.Millisecond) // Delay between frames

}

func main() {
	const (
		w = 11
		h = 11
	)

	studentX := 0
	studentY := 0

	//CTM1.1
	g := NewGrid(w, h)
	g.Draw()

	// CTM1.2
	g.Set('s', studentX, studentY)
	g.Draw()

	//CTM2
	fmt.Print("Move: ") //Print function is used to display output in same line
	var move string
	fmt.Scanln(&move)
	fmt.Print(move)

	switch move {
	case "n":
		studentY = studentY + 1
	case "s":
		studentY = studentY - 1
	case "e":
		studentX = studentX + 1
	case "w":
		studentX = studentX - 1
	}

	g.Set('s', studentX, studentY)
	g.Draw()

}

func max(is ...int) int {
	if len(is) == 0 {
		return 0
	}

	m := is[0]
	for _, v := range is[1:] {
		if v > m {
			m = v
		}
	}
	return m
}
