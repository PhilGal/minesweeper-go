package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// var f.cells [size][size]string

type pos struct {
	x, y int
}

func (p pos) equals(other pos) bool {
	return p.x == other.x && p.y == other.y
}

type field struct {
	size       int
	minesCount int
	cells      [][]string
	mines      [][]bool
	cursor     pos
}

func (f *field) moveCursorUp() {
	if f.cursor.y > 0 {
		f.cursor.y--
	}
}
func (f *field) moveCursorDown() {
	if f.cursor.y < f.size {
		f.cursor.y++
	}
}
func (f *field) moveCursorLeft() {
	if f.cursor.x > 0 {
		f.cursor.x--
	}
}
func (f *field) moveCursorRight() {
	if f.cursor.x < f.size {
		f.cursor.x++
	}
}

func (f *field) printCell(p pos, open bool) {
	cell := f.cells[p.x][p.y]
	if open && f.mines[p.x][p.y] {
		cell = mineCell
	}

	if p.equals(f.cursor) {
		cell = "<" + cell + ">"
	} else {
		cell = "[" + cell + "]"
	}
	fmt.Print(cell)
}

// func numberCell(n int) string {
// 	if n <= 0 {
// 		return emptyCell
// 	}
// 	if n > 8 {
// 		n = 8
// 	}
// 	return fmt.Sprintf("[%v]", n)
// }

func (f *field) newField() {
	f.cells = make([][]string, f.size)
	f.mines = make([][]bool, f.size)
	minesSet := 0
	for y := 0; y < f.size; y++ {
		rowc := make([]string, f.size)
		rowm := make([]bool, f.size)
		for x := 0; x < f.size; x++ {
			rowc[x] = closedCell
			f.cells[y] = rowc
			//f.cells[x][y] = closedCell
			f.mines[y] = rowm
			if minesSet < f.minesCount {
				if rand.Int()%(f.size) == x {
					rowm[x] = true
					minesSet++
				}
			} else {
				rowm[x] = false
			}

		}
	}
	fmt.Printf("Created a new field of size %v\nType 'show' to display it!\n", f.size)
	hasField = true
}

func (f *field) drawCells(open bool) {
	for y := 0; y < f.size; y++ {
		fmt.Println()
		for x := 0; x < f.size; x++ {
			f.printCell(pos{x, y}, open)
		}
	}
	fmt.Println()
}

func (f *field) actionOnCell(a action, p pos) {
	var gameOverError error
	switch a {
	case actionFlag:
		f.doFlag(p)
		break
	case actionOpen:
		gameOverError = f.doOpen(p)
		break
	default:
		//todo say unknown action
		break
	}

	if gameOverError != nil {
		f.openField()
	} else {
		// clearScreen()
		f.showField()
	}
}

func (f *field) doFlag(p pos) {
	if f.cells[p.x][p.y] == flaggedCell {
		fmt.Printf("Removing flag at %v", p)
		f.cells[p.x][p.y] = closedCell
		return
	}

	if f.cells[p.x][p.y] != closedCell {
		fmt.Println("Cell is already open!")
		return
	}

	f.cells[p.x][p.y] = flaggedCell
	f.cursor = p
}

func (f *field) doOpen(p pos) error {
	defer f.setCursor(p)

	if f.cells[p.x][p.y] == flaggedCell {
		fmt.Printf("Cell must be unflagged first! [%v,%v]", p.x, p.y)
		return nil
	}
	if f.cells[p.x][p.y] != closedCell {
		fmt.Println("Cell is already open!")
		return nil
	}
	//todo generate mines or something?
	if f.mines[p.x][p.x] {
		f.cells[p.x][p.y] = mineCell
		return errors.New("game over")
	}

	f.cells[p.x][p.y] = emptyCell
	return nil
}

func (f *field) setCursor(p pos) {
	f.cursor = p
}

func (f *field) showField() {
	if !hasField {
		fmt.Println("Create a new game using 'new'")
		fmt.Print("-> ")
		return
	}
	f.drawCells(false)
}

//opens a field on game over
func (f *field) openField() {
	f.drawCells(true)
}
