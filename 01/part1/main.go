package main

import (
	"advent-of-code/helper"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := helper.ReadFromFile("data.txt")
	stack1, stack2 := prepareStacks(data)

	diff := 0
	for i := 0; i < len(stack1); i++ {
		diff = diff + int(math.Abs(float64(stack1[i]-stack2[i])))
	}

	fmt.Print(diff)
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
	sort.Ints(stack1)
	sort.Ints(stack2)

	return stack1, stack2
}
