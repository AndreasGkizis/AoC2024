package main

import (
	"AoC2024/common"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	text, err := common.ReadLinesToStringSlice("input.txt")
	if err != nil {
		log.Panic(err)
	}
	part1 := SolvePart1(text)
	log.Print(part1)

	part2 := SolvePart2(text)
	log.Print(part2)
}

func SolvePart1(text []string) int {
	var result int
	var pattern = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	for _, line := range text {
		muls := pattern.FindAllString(line, -1)
		result += CalculateMuls(muls)
	}
	return result
}

func CalculateMuls(muls []string) int {
	result := 0
	for _, mul := range muls {
		result += CalculateSingleMul(mul)
	}
	return result
}

func CalculateSingleMul(v string) int {
	nopre, _ := strings.CutPrefix(v, "mul(")

	nopreandsuf, _ := strings.CutSuffix(nopre, ")")
	nums := strings.Split(nopreandsuf, ",")

	firstNum, _ := strconv.Atoi(nums[len(nums)-1])
	secondNum, _ := strconv.Atoi(nums[0])
	return firstNum * secondNum
}

func SolvePart2(text []string) int {
	var result int
	var pattern = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	for _, line := range text {
		muls := pattern.FindAllString(line, -1)
		result += CalculateMuls(muls)
	}
	return result
}
