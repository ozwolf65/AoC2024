package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func p2Conversion(input []byte) string {
	r, err := regexp.Compile(`(?s)don't\(\).*?do\(\)`)
	if err != nil {
		return ""
	}
	del := r.FindAll(input, -1)
	data := string(input)
	for _, elem := range del {
		data = strings.Replace(data, string(elem), "", 1)
	}
	return data
}

func main() {
	f, err := os.ReadFile("./day3/input.txt")
	if err != nil {
		panic(err)
	}
	input := string(f)
	r, err := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 2; i++ {
		println("Part ", i+1)
		total := 0
		if i == 1 {
			input = p2Conversion(f)
		}
		matches := r.FindAll([]byte(input), -1)
		for _, match := range matches {
			str := string(match)
			lower := strings.Index(str, "(")
			higher := strings.Index(str, ")")
			nums := str[lower+1 : higher]
			if nums == "" {
				if strings.Index(str, ")") != -1 {
					nums = str[:strings.Index(str, ")")]
				}
			}
			left, right := strings.Split(nums, ",")[0], strings.Split(nums, ",")[1]
			leftInt, err := strconv.Atoi(left)
			if err != nil {
				panic(err)
			}
			rightInt, err := strconv.Atoi(right)
			if err != nil {
				panic(err)
			}
			total += leftInt * rightInt
		}
		println(total)
		println("--------------------")

	}

}
