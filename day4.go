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

	word := "XMAS"
	count := countWordOccurrences(grid, word)
	fmt.Printf("The word %s occurs %d times.\n", word, count)
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

func countWordOccurrences(grid [][]rune, word string) int {
	directions := [][2]int{
		{-1, -1}, // Up-Left
		{-1, 0},  // Up
		{-1, 1},  // Up-Right
		{0, -1},  // Left
		{0, 1},   // Right
		{1, -1},  // Down-Left
		{1, 0},   // Down
		{1, 1},   // Down-Right
	}

	count := 0
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	wordLen := len(word)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, dir := range directions {
				found := true
				for k := 0; k < wordLen; k++ {
					ni := i + k*dir[0]
					nj := j + k*dir[1]
					if ni < 0 || ni >= rows || nj < 0 || nj >= cols || grid[ni][nj] != rune(word[k]) {
						found = false
						break
					}
				}
				if found {
					count++
				}
			}
		}
	}
	return count
}
