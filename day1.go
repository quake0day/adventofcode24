package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	// Ensure both lists have the same length
	if len(leftList) != len(rightList) {
		fmt.Println("The left and right lists have different lengths.")
		return
	}

	// Sort the left and right lists individually
	sort.Ints(leftList)
	sort.Ints(rightList)

	// Initialize total distance
	totalDistance := 0

	// Compute the total distance by pairing up the numbers
	for i := 0; i < len(leftList); i++ {
		distance := abs(leftList[i] - rightList[i])
		totalDistance += distance
	}

	// Output the total distance
	fmt.Println("Total distance between the lists:", totalDistance)
}

// abs returns the absolute value of an integer
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
