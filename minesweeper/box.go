package minesweeper

import (
	"math/rand"
	"strconv"
	"strings"
)

var myBoxes [][]box
var myMineCount = 10
var myMineIdx map[string]struct{}
var myOpenedCount = 0

func genBoxes() {
	for y := 0; y < height; y++ {
		xBox := []box{}
		for x := 0; x < width; x++ {
			xBox = append(xBox, box{})
		}
		myBoxes = append(myBoxes, xBox)
	}
}

func setBoxMine() {
	spacePicker := make([]string, 0)
	spaceCount := height*width - 1

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			scpaceCounter := strconv.Itoa(y) + "-" + strconv.Itoa(x)
			spacePicker = append(spacePicker, scpaceCounter)
		}
	}

	uix := (yui)*width + xui
	spacePicker[spaceCount], spacePicker[uix] = spacePicker[uix], spacePicker[spaceCount]
	spaceCount--

	myMineIdx = map[string]struct{}{}
	for i := 0; i < myMineCount; i++ {
		randI := rand.Intn(spaceCount)
		idxs := strings.Split(spacePicker[randI], "-")
		y, _ := strconv.Atoi(idxs[0])
		x, _ := strconv.Atoi(idxs[1])
		myBoxes[y][x].mineStatus = true
		myBoxes[y][x].label = "X"
		myMineIdx[spacePicker[randI]] = struct{}{}

		spacePicker[spaceCount], spacePicker[randI] = spacePicker[randI], spacePicker[spaceCount]
		spaceCount--
	}
}

func openBox(yin, xin int) {
	if myBoxes[yin][xin].openStatus {
		return
	}
	myBoxes[yin][xin].openStatus = true
	myBoxes[yin][xin].label = "#"
	myOpenedCount++

	yStart := yin - 1
	if yStart < 0 {
		yStart = 0
	}
	yEnd := yin + 1
	if yEnd > height-1 {
		yEnd = height - 1
	}

	xStart := xin - 1
	if xStart < 0 {
		xStart = 0
	}
	xEnd := xin + 1
	if xEnd > width-1 {
		xEnd = width - 1
	}

	minArroundCount := 0
	for y := yStart; y <= yEnd; y++ {
		for x := xStart; x <= xEnd; x++ {
			if (y == yin && x == xin) || myBoxes[y][x].openStatus {
				continue
			}
			if myBoxes[y][x].mineStatus {
				minArroundCount++
			}
		}
	}

	if minArroundCount == 0 {
		for y := yStart; y <= yEnd; y++ {
			for x := xStart; x <= xEnd; x++ {
				if (y == yin && x == xin) || myBoxes[y][x].openStatus {
					continue
				}
				myBoxes[yin][xin].label = "o" //string(byte(178))
				openBox(y, x)
			}
		}
	} else {
		myBoxes[yin][xin].label = strconv.Itoa(minArroundCount)
	}
	myBoxes[yin][xin].openStatus = true
}

func revealMine() {
	for idx := range myMineIdx {
		idxs := strings.Split(idx, "-")
		y, _ := strconv.Atoi(idxs[0])
		x, _ := strconv.Atoi(idxs[1])
		myBoxes[y][x].openStatus = true
	}
}
