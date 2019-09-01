package main

type state struct {
}

type action int

const (
	actionOpen action = iota
	actionFlag
)

type cell string

const (
	emptyCell   = "•"
	closedCell  = " "
	mineCell    = "*"
	flaggedCell = "f"
	boomCell    = "X"
)
