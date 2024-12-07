package main

import (
	"advent-of-code/helper"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data := helper.ReadFromFile("data.txt")
	for _, line := range data {
		sum := int64(0)

		re := regexp.MustCompile(`mul\(\d+,\d+\)`)

		matches := re.FindAllString(line, -1)

		fmt.Println(matches)

		for _, m := range matches {
			re2 := regexp.MustCompile(`\d+,\d+`)
			matches2 := re2.FindAllString(m, -1)

			for _, m := range matches2 {
				values := strings.Split(m, ",")
				v1, err := strconv.Atoi(values[0])
				v2, err2 := strconv.Atoi(values[1])
				if err != nil || err2 != nil {
					log.Fatal(err, err2)
				}
				fmt.Printf("%d * %d = %d\n", v1, v2, v1*v2)
				sum += int64(v1) * int64(v2)
			}

		}

		fmt.Println(sum)
	}
}
