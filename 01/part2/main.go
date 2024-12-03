package main

import (
	"advent-of-code/helper"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	data := helper.ReadFromFile("data.txt")
	stack1, stack2 := prepareStacks(data)

	sim := 0

	for _, row := range stack1 {
		multiplier := 0
		for _, row2 := range stack2 {
			if row == row2 {
				multiplier++
			}
		}

		sim += row * multiplier
	}

	fmt.Print(sim)
}

func prepareStacks(data []string) ([]int, []int) {
	stack1 := []int{}
	stack2 := []int{}

	for _, line := range data {
		s := strings.Split(line, "   ")
		v1, err := strconv.Atoi(s[0])
		v2, err2 := strconv.Atoi(s[1])
		if err != nil || err2 != nil {
			log.Fatal(err, err2)
		}
		stack1 = append(stack1, v1)
		stack2 = append(stack2, v2)
	}

	return stack1, stack2
}
