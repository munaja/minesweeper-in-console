package minesweeper

import (
	"fmt"
	"strings"
)

var width = 10
var height = 10
var wide = width * height

func drawBoard() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("=" + strings.Repeat("====", width+1))
	fmt.Print("Y\\X |")
	for x := 0; x < width; x++ {
		fmt.Print(" " + string(byte(x+65)) + " |")
	}
	fmt.Print("\n")
	fmt.Print("-" + strings.Repeat("----", width+1) + "\n")

	for y := 0; y < height; y++ {
		fmt.Print(" " + string(byte(y+65)) + "  |")
		for x := 0; x < width; x++ {
			if myBoxes[y][x].openStatus {
				fmt.Print(" " + myBoxes[y][x].label + " |")
			} else {
				fmt.Print("   |")
			}
		}
		fmt.Print("\n")
		fmt.Print("-" + strings.Repeat("----", width+1) + "\n")
	}

}
