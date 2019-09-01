package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	//term "github.com/nsf/termbox-go"
)

func read(reader *bufio.Reader) string {
	cmd, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	return strings.Trim(cmd, "\n")
}

var hasField bool

func main() {
	clearScreen()
	fmt.Printf("Welcome to Minesweeper for %v\n", runtime.GOOS)
	fmt.Print("-> ")
	reader := bufio.NewReader(os.Stdin)
	field := &field{size: 9, minesCount: 10}
	for {
		cmd := read(reader)
		switch cmd {
		case "q":
			fmt.Println("Good bye!")
			return
		case "new":
			field.newField()
			break
		case "show":
			field.showField()
			break
		}

		if !hasField {
			fmt.Println("Create a new game using 'new'")
			fmt.Print("-> ")
			continue
		}

		if strings.HasPrefix(cmd, "flag") {
			pos := readPositionCoords(cmd, field)
			field.actionOnCell(actionFlag, pos)
		}

		if strings.HasPrefix(cmd, "open") {
			pos := readPositionCoords(cmd, field)
			field.actionOnCell(actionOpen, pos)
		}

		fmt.Printf("%q -> ", cmd)
	}
}

func readPositionCoords(cmd string, f *field) pos {
	flagCmd := strings.Split(cmd, " ")
	x, _ := strconv.Atoi(flagCmd[1])
	y, _ := strconv.Atoi(flagCmd[2])
	if x > f.size-1 {
		x = f.size - 1
	}
	if y > f.size-1 {
		y = f.size - 1
	}
	return pos{x, y}
}
