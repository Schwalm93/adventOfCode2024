package main

import (
	"advent-of-code/helper"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	data := helper.ReadFromFile("data.txt")
	stack := prepareStack(data, " ")

	saveCount := 0

	for _, report := range stack {
		if isSafe(report) {
			saveCount++
			continue
		}

		if canBeMadeSafe(report) {
			saveCount++
		}
	}

	fmt.Println(saveCount)
}

func isSafe(line []int) bool {
	return isSorted(line, true) || isSorted(line, false)
}

func canBeMadeSafe(line []int) bool {
	for i := range line {
		temp := append([]int{}, line[:i]...)
		temp = append(temp, line[i+1:]...)
		if isSafe(temp) {
			return true
		}
	}
	return false
}

func isSorted(line []int, ascending bool) bool {
	for i := 0; i < len(line)-1; i++ {
		diff := int(math.Abs(float64(line[i] - line[i+1])))
		if diff < 1 || diff > 3 {
			return false
		}
		if ascending && line[i] >= line[i+1] {
			return false
		}
		if !ascending && line[i] <= line[i+1] {
			return false
		}
	}
	return true
}

func prepareStack(data []string, seperator string) [][]int {
	stack := [][]int{}
	for _, line := range data {
		s := strings.Split(line, seperator)
		stackPart := []int{}
		for _, value := range s {
			atoi, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			stackPart = append(stackPart, atoi)
		}
		stack = append(stack, stackPart)
	}
	return stack
}
