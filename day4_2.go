package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	grid, err := readGrid("day4.txt")
	if err != nil {
		log.Fatalf("Error reading grid: %v", err)
	}

	count := countXMASOccurrences(grid)
	fmt.Printf("The X-MAS appears %d times.\n", count)
}

func readGrid(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		grid = append(grid, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return grid, nil
}

func countXMASOccurrences(grid [][]rune) int {
	count := 0
	rows := len(grid)
	if rows < 3 {
		return 0
	}
	cols := len(grid[0])
	if cols < 3 {
		return 0
	}

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			// Get the diagonals crossing at (i, j)
			diag1 := []rune{grid[i-1][j-1], grid[i][j], grid[i+1][j+1]}
			diag2 := []rune{grid[i-1][j+1], grid[i][j], grid[i+1][j-1]}

			// Check if both diagonals form "MAS" or "SAM"
			if isMASOrSAM(diag1) && isMASOrSAM(diag2) {
				count++
			}
		}
	}
	return count
}

func isMASOrSAM(diagonal []rune) bool {
	if len(diagonal) != 3 {
		return false
	}
	sequences := [][]rune{
		{'M', 'A', 'S'},
		{'S', 'A', 'M'},
	}
	for _, seq := range sequences {
		if matchesSequence(diagonal, seq) || matchesSequence(reverseRuneSlice(diagonal), seq) {
			return true
		}
	}
	return false
}

func matchesSequence(diagonal, sequence []rune) bool {
	for i := 0; i < 3; i++ {
		if diagonal[i] != sequence[i] {
			return false
		}
	}
	return true
}

func reverseRuneSlice(s []rune) []rune {
	reversed := make([]rune, len(s))
	for i := 0; i < len(s); i++ {
		reversed[i] = s[len(s)-1-i]
	}
	return reversed
}
