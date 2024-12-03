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

	fmt.Println(stack)
	for x, report := range stack {

		count := 0

		for i := 1; i < len(report); i++ {
			if report[i-1] < report[i] {
				if int(math.Abs(float64(report[i-1]-report[i]))) > 3 {
					continue
				}
				count++
			}
		}
		fmt.Println(x, count, len(report))
		if count == (len(report) - 1) {
			fmt.Println("HIT Increase")
			saveCount++
			continue
		}

		count = 0
		for i := 1; i < len(report); i++ {
			if report[i-1] > report[i] {
				if int(math.Abs(float64(report[i-1]-report[i]))) > 3 {
					continue
				}
				fmt.Println(report[i-1], report[i])
				count++
			}
		}
		fmt.Println(count, len(report)-1)
		if count == (len(report) - 1) {
			saveCount++
			fmt.Println(report, "HIT Decrease")
			continue
		}
	}
	fmt.Println(saveCount)
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
