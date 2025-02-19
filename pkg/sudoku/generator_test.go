package sudoku

import (
	"testing"
)

func TestGenerateFullBoard(t *testing.T) {
	s := NewSudoku()
	s.GenerateFullBoard()
	if s.hasAnyEmptyCell() {
		t.Errorf("the Sudoku board is not fullfilled:\n%s", s.PrintToString())
	}
	if !s.isValidSudoku() {
		t.Errorf("invalid Sudoku board:\n%s", s.PrintToString())
	}
	if !s.isUniqueSolution() {
		t.Errorf("this Sudoku board is not unique-solution:\n%s", s.PrintToString())
	}
}

func BenchmarkGenerateFullBoard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := NewSudoku()
		s.GenerateFullBoard()
	}
}
