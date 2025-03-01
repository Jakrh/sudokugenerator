package sudoku

import "testing"

func TestIsUniqueSolution(t *testing.T) {
	tests := []struct {
		description string
		board       [boardSize][boardSize]uint8
		expected    bool
	}{
		{
			description: "Unique solution",
			board: [boardSize][boardSize]uint8{
				{5, 3, 0, 0, 7, 0, 0, 0, 0},
				{6, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 0, 2, 8, 0},
				{0, 0, 0, 4, 1, 9, 0, 0, 5},
				{0, 0, 0, 0, 8, 0, 0, 7, 9},
			},
			expected: true,
		},
		{
			description: "Multiple solutions",
			board: [boardSize][boardSize]uint8{
				{5, 3, 0, 0, 7, 0, 0, 0, 0},
				{6, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 0, 2, 8, 0},
				{0, 0, 0, 4, 1, 9, 0, 0, 5},
				{0, 0, 0, 0, 8, 0, 0, 0, 0},
			},
			expected: false,
		},
		{
			description: "No solution",
			board: [boardSize][boardSize]uint8{
				{5, 3, 0, 0, 7, 0, 0, 0, 0},
				{6, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 0, 2, 8, 0},
				{0, 0, 0, 4, 1, 9, 0, 0, 5},
				{0, 0, 0, 0, 8, 0, 0, 0, 0},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			s := NewSudokuWithBoard(tt.board)
			result := s.isUniqueSolution()
			if result != tt.expected {
				t.Errorf("expected %v, but got %v", tt.expected, result)
			}
		})
	}
}

func TestRemoveNumbers(t *testing.T) {
	tests := []struct {
		description string
		board       [boardSize][boardSize]uint8
		target      uint
	}{
		{
			description: "Remove 5 numbers",
			board: [boardSize][boardSize]uint8{
				{6, 7, 4, 5, 2, 1, 9, 3, 8},
				{9, 8, 1, 7, 3, 4, 2, 5, 6},
				{3, 2, 5, 8, 9, 6, 1, 7, 4},
				{7, 3, 8, 1, 4, 2, 6, 9, 5},
				{2, 4, 6, 3, 5, 9, 7, 8, 1},
				{1, 5, 9, 6, 8, 7, 3, 4, 2},
				{8, 1, 2, 4, 7, 3, 5, 6, 9},
				{4, 9, 7, 2, 6, 5, 8, 1, 3},
				{5, 6, 3, 9, 1, 8, 4, 2, 7},
			},
			target: 5,
		},
		{
			description: "Remove 10 numbers",
			board: [boardSize][boardSize]uint8{
				{5, 8, 7, 3, 9, 2, 4, 6, 1},
				{6, 1, 2, 5, 7, 4, 8, 9, 3},
				{9, 4, 3, 1, 6, 8, 2, 7, 5},
				{8, 6, 5, 9, 2, 7, 3, 1, 4},
				{3, 7, 9, 8, 4, 1, 6, 5, 2},
				{1, 2, 4, 6, 5, 3, 9, 8, 7},
				{4, 9, 8, 7, 3, 5, 1, 2, 6},
				{7, 3, 1, 2, 8, 6, 5, 4, 9},
				{2, 5, 6, 4, 1, 9, 7, 3, 8},
			},
			target: 10,
		},
		{
			description: "Remove 20 numbers",
			board: [boardSize][boardSize]uint8{
				{3, 2, 9, 6, 7, 5, 1, 8, 4},
				{1, 8, 6, 9, 3, 4, 5, 7, 2},
				{7, 5, 4, 2, 1, 8, 6, 3, 9},
				{6, 1, 2, 4, 8, 3, 9, 5, 7},
				{4, 7, 5, 1, 2, 9, 3, 6, 8},
				{9, 3, 8, 5, 6, 7, 4, 2, 1},
				{8, 4, 3, 7, 9, 6, 2, 1, 5},
				{2, 9, 7, 3, 5, 1, 8, 4, 6},
				{5, 6, 1, 8, 4, 2, 7, 9, 3},
			},
			target: 20,
		},
		{
			description: "Remove 50 numbers",
			board: [boardSize][boardSize]uint8{
				{8, 2, 9, 3, 6, 1, 4, 7, 5},
				{1, 5, 7, 8, 9, 4, 3, 6, 2},
				{4, 6, 3, 2, 5, 7, 9, 8, 1},
				{2, 9, 4, 7, 8, 6, 5, 1, 3},
				{6, 3, 8, 1, 4, 5, 2, 9, 7},
				{5, 7, 1, 9, 2, 3, 8, 4, 6},
				{7, 8, 5, 4, 1, 2, 6, 3, 9},
				{3, 4, 6, 5, 7, 9, 1, 2, 8},
				{9, 1, 2, 6, 3, 8, 7, 5, 4},
			},
			target: 50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			s := NewSudokuWithBoard(tt.board)
			s.removeNumbers(tt.target)
			removedCount := s.countZeros()
			if removedCount != int(tt.target) {
				t.Errorf("expected %d numbers to be removed, but got %d", tt.target, removedCount)
			}
			if !s.isValidSudoku() {
				t.Errorf("puzzle is invalid after removing numbers: \n%s", s.PrintToString())
			}
		})
	}
}

func TestCreatePuzzle(t *testing.T) {
	tests := []struct {
		description       string
		board             [boardSize][boardSize]uint8
		removedPercentage uint
	}{
		{
			description: "Very Easy difficulty",
			board: [boardSize][boardSize]uint8{
				{6, 7, 4, 5, 2, 1, 9, 3, 8},
				{9, 8, 1, 7, 3, 4, 2, 5, 6},
				{3, 2, 5, 8, 9, 6, 1, 7, 4},
				{7, 3, 8, 1, 4, 2, 6, 9, 5},
				{2, 4, 6, 3, 5, 9, 7, 8, 1},
				{1, 5, 9, 6, 8, 7, 3, 4, 2},
				{8, 1, 2, 4, 7, 3, 5, 6, 9},
				{4, 9, 7, 2, 6, 5, 8, 1, 3},
				{5, 6, 3, 9, 1, 8, 4, 2, 7},
			},
			removedPercentage: 10,
		},
		{
			description: "Easy difficulty",
			board: [boardSize][boardSize]uint8{
				{5, 8, 7, 3, 9, 2, 4, 6, 1},
				{6, 1, 2, 5, 7, 4, 8, 9, 3},
				{9, 4, 3, 1, 6, 8, 2, 7, 5},
				{8, 6, 5, 9, 2, 7, 3, 1, 4},
				{3, 7, 9, 8, 4, 1, 6, 5, 2},
				{1, 2, 4, 6, 5, 3, 9, 8, 7},
				{4, 9, 8, 7, 3, 5, 1, 2, 6},
				{7, 3, 1, 2, 8, 6, 5, 4, 9},
				{2, 5, 6, 4, 1, 9, 7, 3, 8},
			},
			removedPercentage: 30,
		},
		{
			description: "Medium difficulty",
			board: [boardSize][boardSize]uint8{
				{3, 2, 9, 6, 7, 5, 1, 8, 4},
				{1, 8, 6, 9, 3, 4, 5, 7, 2},
				{7, 5, 4, 2, 1, 8, 6, 3, 9},
				{6, 1, 2, 4, 8, 3, 9, 5, 7},
				{4, 7, 5, 1, 2, 9, 3, 6, 8},
				{9, 3, 8, 5, 6, 7, 4, 2, 1},
				{8, 4, 3, 7, 9, 6, 2, 1, 5},
				{2, 9, 7, 3, 5, 1, 8, 4, 6},
				{5, 6, 1, 8, 4, 2, 7, 9, 3},
			},
			removedPercentage: 50,
		},
		{
			description: "Hard difficulty",
			board: [boardSize][boardSize]uint8{
				{7, 3, 9, 8, 1, 4, 6, 2, 5},
				{4, 2, 6, 3, 9, 5, 7, 1, 8},
				{1, 8, 5, 2, 7, 6, 3, 9, 4},
				{5, 4, 3, 9, 6, 1, 2, 8, 7},
				{6, 7, 1, 5, 8, 2, 4, 3, 9},
				{2, 9, 8, 4, 3, 7, 5, 6, 1},
				{3, 6, 7, 1, 5, 8, 9, 4, 2},
				{8, 5, 2, 6, 4, 9, 1, 7, 3},
				{9, 1, 4, 7, 2, 3, 8, 5, 6},
			},
			removedPercentage: 70,
		},
		{
			description: "Very Hard difficulty",
			board: [boardSize][boardSize]uint8{
				{1, 5, 2, 3, 9, 8, 6, 4, 7},
				{7, 3, 4, 6, 1, 2, 5, 8, 9},
				{6, 9, 8, 5, 4, 7, 2, 3, 1},
				{9, 4, 6, 2, 8, 3, 7, 1, 5},
				{3, 1, 5, 4, 7, 6, 9, 2, 8},
				{2, 8, 7, 9, 5, 1, 3, 6, 4},
				{8, 6, 9, 7, 3, 4, 1, 5, 2},
				{5, 2, 1, 8, 6, 9, 4, 7, 3},
				{4, 7, 3, 1, 2, 5, 8, 9, 6},
			},
			removedPercentage: 90,
		},
		{
			// Not really remove all 64 numbers for now
			description: "evil Sudoku with 17 initial values",
			board: [boardSize][boardSize]uint8{
				{8, 2, 9, 3, 6, 1, 4, 7, 5},
				{1, 5, 7, 8, 9, 4, 3, 6, 2},
				{4, 6, 3, 2, 5, 7, 9, 8, 1},
				{2, 9, 4, 7, 8, 6, 5, 1, 3},
				{6, 3, 8, 1, 4, 5, 2, 9, 7},
				{5, 7, 1, 9, 2, 3, 8, 4, 6},
				{7, 8, 5, 4, 1, 2, 6, 3, 9},
				{3, 4, 6, 5, 7, 9, 1, 2, 8},
				{9, 1, 2, 6, 3, 8, 7, 5, 4},
			},
			removedPercentage: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			s := NewSudokuWithBoard(tt.board)
			s.CreatePuzzle(tt.removedPercentage)
			if !s.isValidSudoku() {
				t.Errorf("puzzle is invalid with percentage of removed numbers %d%%: \n%s", tt.removedPercentage, s.PrintToString())
			}
		})
	}
}
