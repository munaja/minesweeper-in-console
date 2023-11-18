package minesweeper

type mine struct{}

type box struct {
	label      string
	mineStatus bool
	openStatus bool
}
