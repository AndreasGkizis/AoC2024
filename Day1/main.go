package main

import (
	common "AoC2024/common"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	Part1()
	Part2()
}

func Part2() {
	text, _ := common.ReadLinesToStringSlice("input.txt")
	leftList, rightList := GetListsFromLines(text)
	result := CalculateSimilarity(leftList, rightList)

	log.Println(result)
}

func Part1() {
	text, _ := common.ReadLinesToStringSlice("input.txt")

	leftList, rightList := GetListsFromLines(text)

	result := CalculateDiffs(leftList, rightList)

	log.Println(result)
}

func CalculateSimilarity(leftList []int, rightList []int) int {
	result := 0
	// map [TheNumber]how many times it appears
	frequencyMap := CreateFrequencyMap(rightList)
	for _, value := range leftList {
		occuranceFreq, ok := frequencyMap[value]

		if ok {
			result += value * occuranceFreq
		}
	}
	return result
}

func CreateFrequencyMap(rightList []int) map[int]int {
	fmap := make(map[int]int)
	for _, val := range rightList {
		_, ok := fmap[val]
		if ok {
			fmap[val]++
		} else {
			fmap[val] = 1
		}
	}
	return fmap
}

func GetListsFromLines(text []string) ([]int, []int) {
	var leftList []int
	var rightList []int
	for _, val := range text {
		parts := strings.Fields(val)

		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)
	return leftList, rightList
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
