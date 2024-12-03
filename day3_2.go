package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
)

type Instruction struct {
	Position int
	Type     string // 'mul', 'do', 'don't'
	X        int    // Only for 'mul' instructions
	Y        int    // Only for 'mul' instructions
}

func main() {
	// Read the entire contents of day3.txt
	data, err := ioutil.ReadFile("day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(data)

	var instructions []Instruction

	// Define regex patterns for mul, do, and don't instructions
	mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	// Find all mul instruction matches
	mulMatches := mulRe.FindAllStringSubmatchIndex(text, -1)
	for _, match := range mulMatches {
		start := match[0]
		xStart, xEnd := match[2], match[3]
		yStart, yEnd := match[4], match[5]
		xStr := text[xStart:xEnd]
		yStr := text[yStart:yEnd]
		x, err1 := strconv.Atoi(xStr)
		y, err2 := strconv.Atoi(yStr)
		if err1 != nil || err2 != nil {
			continue // Skip if conversion fails
		}
		instr := Instruction{
			Position: start,
			Type:     "mul",
			X:        x,
			Y:        y,
		}
		instructions = append(instructions, instr)
	}

	// Find all do() instruction matches
	doMatches := doRe.FindAllStringIndex(text, -1)
	for _, match := range doMatches {
		start := match[0]
		instr := Instruction{
			Position: start,
			Type:     "do",
		}
		instructions = append(instructions, instr)
	}

	// Find all don't() instruction matches
	dontMatches := dontRe.FindAllStringIndex(text, -1)
	for _, match := range dontMatches {
		start := match[0]
		instr := Instruction{
			Position: start,
			Type:     "don't",
		}
		instructions = append(instructions, instr)
	}

	// Sort instructions by their position in the text
	sort.Slice(instructions, func(i, j int) bool {
		return instructions[i].Position < instructions[j].Position
	})

	// Process the instructions
	totalSum := 0
	enabled := true // At the beginning, mul instructions are enabled

	for _, instr := range instructions {
		switch instr.Type {
		case "do":
			enabled = true
		case "don't":
			enabled = false
		case "mul":
			if enabled {
				product := instr.X * instr.Y
				totalSum += product
			}
		}
	}

	// Output the total sum
	fmt.Println(totalSum)
}
