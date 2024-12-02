package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(data [][]int) int {
	count := 0
	for _, row := range data {
		if comp(row) {
			count++
		}
	}
	return count
}

func part2(data [][]int) int {
	count := 0
	for _, row := range data {
		if comp(row) || compMinusOne(row) {
			count++
		}
	}
	return count
}

func comp(row []int) bool {
	if row[0] < row[1] {
		if asc(row) {
			return true
		}
	} else if row[0] > row[1] {
		if desc(row) {
			return true
		}
	}
	return false
}

func compMinusOne(row []int) bool {
	for i := 0; i < len(row); i++ {
		minusOne := make([]int, 0)
		minusOne = append(minusOne, row[:i]...)
		minusOne = append(minusOne, row[i+1:]...)
		if comp(minusOne) {
			return true
		}
	}
	return false
}

func asc(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		if row[i+1]-row[i] < 1 || row[i+1]-row[i] > 3 {
			return false
		}
	}
	return true
}

func desc(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		if row[i]-row[i+1] < 1 || row[i]-row[i+1] > 3 {
			return false
		}
	}
	return true
}

func main() {
	input, err := os.ReadFile("./day2/input.txt")
	if err != nil {
		panic(err)
	}
	rows := strings.Split(string(input), "\n")
	data := make([][]int, 0)
	for _, row := range rows {
		if row == "" {
			continue
		}
		rowVals := make([]int, 0)
		for _, val := range strings.Split(row, " ") {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			rowVals = append(rowVals, num)
		}
		data = append(data, rowVals)
	}
	p1 := part1(data)
	p2 := part2(data)
	fmt.Println("Day 1 Advent of Code 2024")
	fmt.Printf("Part 1: %v \n", int(p1))
	fmt.Println("Part 2: ", p2)
}
