package main

import (
	"AoC2024/common"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	text, err := common.ReadLinesToStringSlice("testinput.txt")
	if err != nil {
		log.Panic(err)
	}
	intSlices := MakeIntSlices(text)
	safeCount := CountSafeLines(intSlices)
	log.Print(safeCount)
}

func CountSafeLines(intSlices [][]int) int {
	result := 0
	for _, val := range intSlices {
		isSafe := IsSliceSafe(val)
		if isSafe {
			result++
		}
	}
	return result

}

func IsSliceSafe(input []int) bool {
	maxDiff := 3
	isAscendingOrder := input[0] < input[1]

	underMaxDiff := isUnderMaxDiff(input, isAscendingOrder, maxDiff)
	staysInSameOrder := staysInOrder(input)

	return underMaxDiff && staysInSameOrder
}

func staysInOrder(input []int) bool {
	asc := sort.SliceIsSorted(input, func(p, q int) bool { return input[p] < input[q] })
	desc := sort.SliceIsSorted(input, func(p, q int) bool { return input[p] > input[q] })

	return asc || desc
}

func isUnderMaxDiff(input []int, isAscendingOrder bool, maxDiff int) bool {
	for indx, val := range input {
		if indx < len(input)-1 {
			left := val
			right := input[indx+1]

			if isAscendingOrder && right-left >= maxDiff {
				return false
			} else if !isAscendingOrder && left-right >= maxDiff {
				return false
			}
		}
	}
	return true
}

func MakeIntSlices(text []string) [][]int {
	var result [][]int
	for _, val := range text {
		nums := BreakStringToInts(val)
		result = append(result, nums)
	}
	return result
}

func BreakStringToInts(input string) []int {
	var result []int
	parts := strings.Fields(input)
	for _, val := range parts {
		number, _ := strconv.Atoi(val)
		result = append(result, number)
	}
	return result
}
