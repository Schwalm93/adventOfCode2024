package main

import (
	"advent-of-code/helper"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func main() {
	// Datei einlesen
	data := helper.ReadFromFile("data.txt")
	totalSum := int64(0)

	// Regulärer Ausdruck für gültige mul(X,Y)-Anweisungen
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	for _, line := range data {
		lineSum := int64(0)

		// Finde alle gültigen mul-Anweisungen und ihre Submatches (Zahlen)
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			// match[1] ist die erste Zahl, match[2] ist die zweite Zahl
			v1, err1 := strconv.Atoi(match[1])
			v2, err2 := strconv.Atoi(match[2])
			if err1 != nil || err2 != nil {
				log.Fatal(err1, err2)
			}

			// Berechne das Produkt
			product := int64(v1) * int64(v2)
			lineSum += product
		}

		fmt.Printf("Summe für Zeile: %d\n", lineSum)
		totalSum += lineSum
	}

	fmt.Printf("Gesamtsumme aller gültigen Multiplikationen: %d\n", totalSum)
}
