package helper

import (
	"bufio"
	"log"
	"os"
)

func ReadFromFile(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	text := []string{}

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return text

}
