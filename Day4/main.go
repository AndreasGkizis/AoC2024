package main

import (
	"AoC2024/common"
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
)

var runeList = []rune{'X', 'M', 'A', 'S'}

func main() {
	text, err := common.ReadLinesToStringSlice("testinput.txt")
	if err != nil {
		log.Panic(err)
	}
	part1 := SolvePart1(text)
	// part2 := SolvePart2(text)
	log.Printf("\t'XMAS' found %d times", part1)
	// log.Print(part2)
}

func SolvePart1(data []string) int {
	// var failedPaths []Point // Collect failed paths

	answer := 0
	points := convertStringToPoint(data)
	// printPoints(points)
	// clearScreen()
	for row, v := range points {
		for col, point := range v {
			if point.Value == string(runeList[0]) /*This is X */ {
				point := points[row][col]
				for _, dir := range Directions {
					if FindXMAS(point, 1, dir, points, true) {
						answer++
					}
				}
			}
		}
	}
	return answer
}

func FindXMAS(point Point, position int, dir Dir, data [][]Point, printgraph bool) bool {

	colNext := point.Col + dir.Col
	rowNext := point.Row + dir.Row

	// good case
	if position > len(runeList)-1 {
		if printgraph {
			colourXMAS(point, dir, color.FgHiRed, 0, data)
		}
		printPoints(data)
		return true /* completed word finding */
	}

	// bad cases
	if colNext < 0 || colNext >= len(data[point.Row]) {
		return false
	}
	if rowNext < 0 || rowNext >= len(data[point.Col]) {
		return false
	}

	if data[rowNext][colNext].Value == string(runeList[position]) {
		if printgraph {
			if data[point.Row][point.Col].Colour != color.FgRed {
				data[point.Row][point.Col].Colour = color.FgGreen
				data[point.Row][point.Col].BackColour = color.BgBlack
			}
			printPoints(data)
		}

		return FindXMAS(data[rowNext][colNext], position+1, dir, data, printgraph)
	}
	return false
}

func colourXMAS(point Point, dir Dir, colour color.Attribute, count int, data [][]Point) {
	reverseDir := LookupByReverseName(dir.CommingFrom)
	data[point.Row][point.Col].Colour = colour
	data[point.Row][point.Col].BackColour = color.BgHiBlack

	colNext := point.Col + reverseDir.Col
	rowNext := point.Row + reverseDir.Row
	if count < len(runeList)-1 {
		colourXMAS(data[rowNext][colNext], dir, colour, count+1, data)
	}
}

type Point struct {
	Row        int
	Col        int
	Value      string
	Colour     color.Attribute
	BackColour color.Attribute
}

type Dir struct {
	CommingFrom string
	Row         int
	Col         int
	ReverseName string
}

var Directions []Dir = []Dir{
	{Row: 1, Col: 0, CommingFrom: "UP", ReverseName: "DOWN"},
	{Row: -1, Col: 0, CommingFrom: "DOWN", ReverseName: "UP"},
	{Row: 0, Col: 1, CommingFrom: "LEFT", ReverseName: "RIGHT"},
	{Row: 0, Col: -1, CommingFrom: "RIGHT", ReverseName: "LEFT"},

	{Row: 1, Col: 1, CommingFrom: "UP-LEFT", ReverseName: "DOWN-RIGHT"},
	{Row: -1, Col: -1, CommingFrom: "DOWN-RIGHT", ReverseName: "UP-LEFT"},

	{Row: -1, Col: 1, CommingFrom: "DOWN-LEFT", ReverseName: "UP-RIGHT"},
	{Row: 1, Col: -1, CommingFrom: "UP-RIGHT", ReverseName: "DOWN-LEFT"},
}

func LookupByReverseName(reverseName string) *Dir {
	for _, dir := range Directions {
		if dir.ReverseName == reverseName {
			return &dir
		}
	}
	return nil
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func printPoints(input [][]Point) {
	clearScreen()
	fmt.Println("---------------------------------------------")
	for _, row := range input {
		line := "  "
		for _, char := range row {
			foreground := color.New(char.Colour)
			combinedColor := foreground.Add(char.BackColour)
			line += " "
			line += combinedColor.Sprint(" " + char.Value + " ")
		}
		fmt.Println(line)
	}
	fmt.Println("---------------------------------------------")
	time.Sleep(175 * time.Millisecond) // Add a delay for visualization
}

func convertStringToPoint(data []string) [][]Point {
	result := make([][]Point, len(data))
	for rowNum, row := range data {
		result[rowNum] = make([]Point, len(row)) // Initialize inner slice for each row
		for colNum, char := range row {
			result[rowNum][colNum] = Point{Value: string(char), BackColour: color.BgBlack, Colour: color.FgWhite, Row: rowNum, Col: colNum}
		}
	}
	return result
}
