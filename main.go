package main

import (
	"sudokugenerator/pkg/sudoku"
)

func createSudokuPuzzle(difficulty uint) {
	s := sudoku.NewSudoku()
	s.GenerateFullBoard()
	s.CreatePuzzle(difficulty)
	s.Print()
}

func main() {
	createSudokuPuzzle(50)
}
