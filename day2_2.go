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
	file, err := os.Open("day2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		// Skip empty lines
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		// Split the line into numbers
		parts := strings.Fields(line)
		numbers := []int{}
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error parsing number:", part)
				return
			}
			numbers = append(numbers, num)
		}

		// Check if the report is safe (with possible removal)
		if isSafeWithRemoval(numbers) {
			safeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Number of safe reports:", safeCount)
}

// isSafe checks if a report meets the original safety criteria
func isSafe(levels []int) bool {
	if len(levels) < 2 {
		// A single level cannot be increasing or decreasing
		return false
	}

	diffs := []int{}
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if diff == 0 {
			// Adjacent levels are equal; neither increasing nor decreasing
			return false
		}
		if abs(diff) < 1 || abs(diff) > 3 {
			// Difference is outside the allowed range
			return false
		}
		diffs = append(diffs, diff)
	}

	// Check if all differences are positive (increasing) or all negative (decreasing)
	allPositive := true
	allNegative := true
	for _, diff := range diffs {
		if diff > 0 {
			allNegative = false
		} else if diff < 0 {
			allPositive = false
		}
	}

	// Report is safe if all differences are either positive or negative
	return allPositive || allNegative
}

// isSafeWithRemoval checks if a report is safe, possibly after removing one level
func isSafeWithRemoval(levels []int) bool {
	// First, check if the report is already safe
	if isSafe(levels) {
		return true
	}

	// Try removing each level one by one and check if the report becomes safe
	for i := 0; i < len(levels); i++ {
		modifiedLevels := make([]int, 0, len(levels)-1)
		modifiedLevels = append(modifiedLevels, levels[:i]...)
		modifiedLevels = append(modifiedLevels, levels[i+1:]...)

		if isSafe(modifiedLevels) {
			return true
		}
	}

	// Report is unsafe even after removing any one level
	return false
}

// abs returns the absolute value of an integer
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
