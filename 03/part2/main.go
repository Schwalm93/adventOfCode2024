package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var contentBuilder strings.Builder
	for scanner.Scan() {
		contentBuilder.WriteString(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	content := contentBuilder.String()

	doRegex := regexp.MustCompile(`\bdo\(\)`)
	dontRegex := regexp.MustCompile(`\bdon't\(\)`)

	mulRegex := regexp.MustCompile(`\bmul\s*\(\s*(\d+)\s*,\s*(\d+)\s*\)`)

	enabled := true
	sum := 0
	combinedRegex := regexp.MustCompile(`\bdon't\(\)|\bdo\(\)|\bmul\s*\(\s*\d+\s*,\s*\d+\s*\)`)
	matches := combinedRegex.FindAllStringIndex(content, -1)

	for _, m := range matches {
		token := content[m[0]:m[1]]

		if doRegex.MatchString(token) {
			enabled = true
		} else if dontRegex.MatchString(token) {
			enabled = false
		} else if mulRegex.MatchString(token) {
			if enabled {
				// Werte auslesen
				groups := mulRegex.FindStringSubmatch(token)
				if len(groups) == 3 {
					xStr := groups[1]
					yStr := groups[2]
					x, errX := strconv.Atoi(xStr)
					y, errY := strconv.Atoi(yStr)
					if errX == nil && errY == nil {
						sum += x * y
					}
				}
			}
		}
	}

	fmt.Println(sum)
}
