package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Numpair struct {
	LeftNum  int
	RightNum int
}

func (n Numpair) GetDif() int {
	if n.LeftNum > n.RightNum {
		return n.LeftNum - n.RightNum
	} else {
		return n.RightNum - n.LeftNum
	}
}

func main() {
	Part1()
	Part2()
}

func Part2() {
	text, _ := ReadFile("testinput.txt")
	leftList, rightList := GetListsFromLines(text)
	result := CalculateSimilarity(leftList, rightList)

	log.Println(result)
}

func CalculateSimilarity(leftList []int, rightList []int) int {
	panic("unimplemented")
}

func Part1() {
	text, _ := ReadFile("testinput.txt")

	leftList, rightList := GetListsFromLines(text)

	result := CalculateDiffs(leftList, rightList)

	log.Println(result)
}

func GetListsFromLines(text []string) ([]int, []int) {
	var leftList []int
	var rightList []int
	for _, val := range text {
		parts := strings.Split(val, "   ")

		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)
	return leftList, rightList
}

func ReadFile(filename string) ([]string, error) {
	readFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text, nil
}

func CalculateDiffs(leftList, rightList []int) int {
	result := 0
	for indx, val := range leftList {
		if val > rightList[indx] {
			result += val - rightList[indx]
		} else {
			result += rightList[indx] - val
		}
	}
	return result
}
