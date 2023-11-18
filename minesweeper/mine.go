package minesweeper

import (
	"fmt"
	"math/rand"
	"time"
)

// x and y from user input
var xui, yui int
var validInput = false
var gameStatus int

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Start() {
	genBoxes()
	for !validInput {
		drawBoard()
		processInput()
	}
	validInput = false
	setBoxMine()
	openBox(yui, xui)

	for gameStatus == 0 {
		for !validInput {
			drawBoard()
			processInput()
		}
		validInput = false

		openBox(yui, xui)
		if wide-myOpenedCount == myMineCount {
			gameStatus = 1
		}
	}
	if gameStatus == 1 {
		fmt.Println("Congratulations!!")
	}
	if gameStatus == 2 {
		revealMine()
		drawBoard()
		fmt.Println("Game is over")
	}
}

func processInput() {
	var y, x string

	// uses x,y format for convenient method used by math in general
	fmt.Print("Input coordinate (X,Y) sparated by space (e.g: \"C D\"): ")
	fmt.Scanln(&x, &y)
	if x == "" || y == "" {
		validInput = false
		return
	}

	// in byte
	xb := []byte(x)
	yb := []byte(y)

	// in integer
	xi := int(xb[0] - 65)
	yi := int(yb[0] - 65)

	// invalid input
	if xi < 0 || xi > width-1 || yi < 0 || yi > height-1 {
		validInput = false
		return
	}
	validInput = true

	// if it's already opened
	if myBoxes[yi][xi].openStatus {
		return
	}

	// if it hits the mine
	if myBoxes[yi][xi].mineStatus {
		gameStatus = 2
		return
	}

	yui = yi
	xui = xi

	return
}
