package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

func main() {
	// Read the entire contents of day3.txt
	data, err := ioutil.ReadFile("day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(data)

	// Define the regular expression pattern
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// Find all matches in the text
	matches := re.FindAllStringSubmatch(text, -1)

	totalSum := 0

	// Process each match
	for _, match := range matches {
		xStr := match[1]
		yStr := match[2]

		// Convert strings to integers
		x, err1 := strconv.Atoi(xStr)
		y, err2 := strconv.Atoi(yStr)

		if err1 != nil || err2 != nil {
			continue // Skip if conversion fails
		}

		product := x * y
		totalSum += product
	}

	// Output the total sum
	fmt.Println(totalSum)
}
