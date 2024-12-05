package main

import (
	"AoC2024/common"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	text, err := common.ReadLinesToStringSlice("input.txt")
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

	underMaxDiff := isUnderMaxDiff(input, maxDiff)
	staysInSameOrder := staysInOrder(input)
	result := underMaxDiff && staysInSameOrder

	if !result {
		log.Print(input)
		log.Print(result)
	}
	return result
}

func staysInOrder(input []int) bool {
	asc := sort.SliceIsSorted(input, func(p, q int) bool { return input[p] < input[q] })
	desc := sort.SliceIsSorted(input, func(p, q int) bool { return input[p] > input[q] })
	var result = asc || desc

	return result
}

func isUnderMaxDiff(input []int, maxDiff int) bool {
	for indx, val := range input {
		if indx < len(input)-1 {
			left := val
			right := input[indx+1]
			diff := absInt(right - left)
			if diff == 0 || diff > maxDiff {
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

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
