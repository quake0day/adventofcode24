package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the input file
	file, err := os.Open("day1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Slices to hold numbers from the left and right lists
	var leftList []int
	var rightList []int

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		// Skip empty lines
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		// Split the line into fields (numbers)
		fields := strings.Fields(line)
		if len(fields) != 2 {
			fmt.Println("Invalid line (does not contain exactly two numbers):", line)
			return
		}

		// Parse the numbers from strings to integers
		leftNum, err := strconv.Atoi(fields[0])
		if err != nil {
			fmt.Println("Invalid number in left list:", fields[0])
			return
		}
		rightNum, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Println("Invalid number in right list:", fields[1])
			return
		}

		// Append the numbers to the respective lists
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	// Check for errors during file scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Compute the similarity score
	similarityScore := computeSimilarityScore(leftList, rightList)

	// Output the similarity score
	fmt.Println("Similarity score:", similarityScore)
}

// computeSimilarityScore calculates the similarity score as per the specified rules
func computeSimilarityScore(leftList, rightList []int) int {
	// Create a frequency map for the right list
	rightFreqMap := make(map[int]int)
	for _, num := range rightList {
		rightFreqMap[num]++
	}

	// Initialize similarity score
	similarityScore := 0

	// Iterate over the left list in order
	for _, leftNum := range leftList {
		// Get the frequency of leftNum in the right list
		freq := rightFreqMap[leftNum]
		// Increase the similarity score
		similarityScore += leftNum * freq
	}

	return similarityScore
}
