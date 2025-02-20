package main

import (
	"fmt"
	"os"
	"strconv"
	"sudokugenerator/pkg/sudoku"
)

const (
	minDifficultyLevel     uint = 1
	maxDifficultyLevel     uint = 5
	defaultDifficultyLevel uint = 2
	maxRemovedPercentage   uint = 90
)

func createSudokuPuzzle(difficulty uint) {
	s := sudoku.NewSudoku()
	s.GenerateFullBoard()
	s.CreatePuzzle(difficulty)
	s.Print()
}

func getDifficultyLevel() (uint, error) {
	var difficultyLevel uint = defaultDifficultyLevel

	if len(os.Args) == 2 {
		arg := os.Args[1]
		level, err := strconv.Atoi(arg)
		if err != nil {
			return 0, fmt.Errorf("invalid input: %v, please enter a number between 1 and 5", arg)
		}
		if level < int(minDifficultyLevel) || level > int(maxDifficultyLevel) {
			return 0, fmt.Errorf("difficulty level must be between 1 and 5")
		}
		difficultyLevel = uint(level)
	}

	return difficultyLevel, nil
}

func getRemovedPercentage(difficultyLevel uint) uint {
	difficultyLevels := map[uint]uint{
		1: 50,
		2: 60,
		3: 70,
		4: 80,
		5: 90,
	}

	var removedPercentage uint
	var ok bool
	if removedPercentage, ok = difficultyLevels[difficultyLevel]; !ok {
		removedPercentage = maxRemovedPercentage
	}

	return removedPercentage
}

func main() {
	difficultyLevel, err := getDifficultyLevel()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	removedPercentage := getRemovedPercentage(difficultyLevel)
	createSudokuPuzzle(removedPercentage)
}
