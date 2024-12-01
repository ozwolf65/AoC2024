package main

import (
	"fmt"
	"log/slog"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func generateLists(input string) ([]int, []int) {
	left := make([]int, 0)
	right := make([]int, 0)
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		if row == "" {
			continue
		}
		vals := strings.Split(row, "   ")
		leftInt, err := strconv.Atoi(vals[0])
		if err != nil {
			slog.Error(err.Error())
			return nil, nil
		}
		left = append(left, leftInt)

		rightInt, err := strconv.Atoi(vals[1])
		if err != nil {
			slog.Error(err.Error())
			return nil, nil
		}
		right = append(right, rightInt)
	}
	return left, right
}

func part1(input string) float64 {
	distance := 0.0
	left, right := generateLists(input)
	if left == nil || right == nil {
		return 0
	}
	slices.Sort(left)
	slices.Sort(right)
	for i := 0; i < len(left); i++ {
		distance += math.Abs(float64(left[i] - right[i]))
	}

	return distance
}

func part2(input string) int {
	similarity := 0
	left, right := generateLists(input)
	if left == nil || right == nil {
		return 0
	}
	rightCount := make(map[int]int)
	for _, num := range right {
		rightCount[num]++
	}
	for _, num := range left {
		similarity += num * rightCount[num]
	}
	return similarity
}

func main() {
	wd, err := os.Getwd()
	println(wd)
	f, err := os.ReadFile("./day1/input.txt")
	if err != nil {
		slog.Error(err.Error())
	}

	p1 := part1(string(f))
	p2 := part2(string(f))

	fmt.Println("Day 1 Advent of Code 2024")
	fmt.Printf("Part 1: %v \n", int(p1))
	fmt.Println("Part 2: ", p2)

}
