package sudoku

import (
	"reflect"
	"strings"
	"testing"
)

func TestClone(t *testing.T) {
	s := NewSudoku()
	s.GenerateFullBoard()
	ss := s.Clone()
	if s.board != ss.board {
		t.Error("cloned board does not the same as the original board")
	}
}

func TestIsValidSudoku(t *testing.T) {
	tests := []struct {
		description string
		board       [boardSize][boardSize]uint8
		expected    bool
	}{
		{
			description: "Valid full board with sequential numbers",
			board: [boardSize][boardSize]uint8{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{4, 5, 6, 7, 8, 9, 1, 2, 3},
				{7, 8, 9, 1, 2, 3, 4, 5, 6},
				{2, 3, 4, 5, 6, 7, 8, 9, 1},
				{5, 6, 7, 8, 9, 1, 2, 3, 4},
				{8, 9, 1, 2, 3, 4, 5, 6, 7},
				{3, 4, 5, 6, 7, 8, 9, 1, 2},
				{6, 7, 8, 9, 1, 2, 3, 4, 5},
				{9, 1, 2, 3, 4, 5, 6, 7, 8},
			},
			expected: true,
		},
		{
			description: "Valid full board with mixed numbers",
			board: [boardSize][boardSize]uint8{
				{2, 4, 5, 9, 8, 1, 7, 3, 6},
				{8, 9, 3, 6, 2, 7, 1, 5, 4},
				{6, 1, 7, 3, 5, 4, 2, 9, 8},
				{5, 6, 4, 2, 9, 3, 8, 1, 7},
				{3, 2, 1, 4, 7, 8, 5, 6, 9},
				{9, 7, 8, 5, 1, 6, 4, 2, 3},
				{4, 5, 2, 8, 6, 9, 3, 7, 1},
				{1, 8, 6, 7, 3, 5, 9, 4, 2},
				{7, 3, 9, 1, 4, 2, 6, 8, 5},
			},
			expected: true,
		},
		{
			description: "invalid cell in board[6][5]",
			board: [boardSize][boardSize]uint8{
				{2, 4, 5, 9, 8, 1, 7, 3, 6},
				{8, 9, 3, 6, 2, 7, 1, 5, 4},
				{6, 1, 7, 3, 5, 4, 2, 9, 8},
				{5, 6, 4, 2, 9, 3, 8, 1, 7},
				{3, 2, 1, 4, 7, 8, 5, 6, 9},
				{9, 7, 8, 5, 1, 6, 4, 2, 3},
				{4, 5, 2, 8, 6, 3, 3, 7, 1},
				{1, 8, 6, 7, 3, 5, 9, 4, 2},
				{7, 3, 9, 1, 4, 2, 6, 8, 5},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			s := NewSudokuWithBoard(tt.board)
			if s.isValidSudoku() != tt.expected {
				t.Errorf("invalid Sudoku board: \n%s", s.PrintToString())
			}
		})
	}
}

func TestIsValidCell(t *testing.T) {
	tests := []struct {
		description string
		board       [boardSize][boardSize]uint8
		row         int
		col         int
		num         uint8
		expected    bool
	}{
		{
			description: "Valid number in empty board",
			board:       [boardSize][boardSize]uint8{},
			row:         0,
			col:         0,
			num:         1,
			expected:    true,
		},
		{
			description: "Invalid number (same number in the same row)",
			board:       [boardSize][boardSize]uint8{{0, 1}},
			row:         0,
			col:         0,
			num:         1,
			expected:    false,
		},
		{
			description: "Invalid number (same number in the same column)",
			board:       [boardSize][boardSize]uint8{{0}, {2}},
			row:         0,
			col:         0,
			num:         2,
			expected:    false,
		},
		{
			description: "Valid number in a filled row",
			board:       [boardSize][boardSize]uint8{{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			row:         0,
			col:         0,
			num:         7,
			expected:    false,
		},
		{
			description: "Valid number in a filled column",
			board: [boardSize][boardSize]uint8{
				{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9},
			},
			row:      0,
			col:      0,
			num:      4,
			expected: false,
		},
		{
			description: "Valid number in the 5-th filled 3x3 subgrid",
			board: [boardSize][boardSize]uint8{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 2, 3, 0, 0, 0},
				{0, 0, 0, 4, 5, 6, 0, 0, 0},
				{0, 0, 0, 7, 8, 9, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			row:      4,
			col:      5,
			num:      9,
			expected: false,
		},
		{
			description: "Valid number in the 9-th filled 3x3 subgrid",
			board: [boardSize][boardSize]uint8{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 2, 3},
				{0, 0, 0, 0, 0, 0, 4, 5, 6},
				{0, 0, 0, 0, 0, 0, 7, 8, 9},
			},
			row:      8,
			col:      6,
			num:      9,
			expected: false,
		},
		{
			description: "Valid number in a partially filled board",
			board: [boardSize][boardSize]uint8{
				{0, 2, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 1}, // the last cell is 1
			},
			row:      0,
			col:      0,
			num:      1,
			expected: true,
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			s := NewSudokuWithBoard(tt.board)
			actual := s.isValidCell(tt.row, tt.col, tt.num)
			if actual != tt.expected {
				t.Errorf("isValid(%d, %d, %d) returned %t, but expected %t", tt.row, tt.col, tt.num, actual, tt.expected)
			}
		})
	}
}

func TestRandPerm(t *testing.T) {
	tests := []struct {
		description string
		start       int
		end         int
	}{
		{
			description: "Range from 1 to 9",
			start:       1,
			end:         9,
		},
		{
			description: "Range from -12 to 121",
			start:       -12,
			end:         121,
		},
		{
			description: "Range with single element 0",
			start:       0,
			end:         0,
		},
		{
			description: "Range with single element 5",
			start:       5,
			end:         5,
		},
		{
			description: "Range from -5 to 5",
			start:       -5,
			end:         5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			// Check amount of numbers and duplications
			s := NewSudoku()
			numSize := tt.end - tt.start + 1
			nums := s.randPerm(tt.start, tt.end)
			if len(nums) != numSize {
				t.Errorf("amount of numbers must be %d, but %d", numSize, len(nums))
			}
			used := make([]bool, numSize)
			for _, num := range nums {
				if used[num-tt.start] {
					t.Errorf("number %d is duplicate", num)
				}
				used[num-tt.start] = true
			}
			for num, isUsed := range used {
				if !isUsed {
					t.Errorf("missing number %d", num+tt.start)
				}
			}

			// Check are numbers shuffled
			if numSize == 1 {
				// No need to check if only one number
				return
			}
			testTimes := 5
			prev := s.randPerm(tt.start, tt.end)
			for i := 0; i < testTimes; i++ {
				nums := s.randPerm(tt.start, tt.end)
				if reflect.DeepEqual(prev, nums) {
					t.Errorf("numbers are not shuffled: %v", nums)
				}
				prev = nums
			}
		})
	}
}

func TestPrintTo(t *testing.T) {
	tests := []struct {
		description string
		board       [boardSize][boardSize]uint8
		expected    string
	}{
		{
			description: "Empty board",
			board:       [boardSize][boardSize]uint8{},
			expected: `┌───────┬───────┬───────┐
│       │       │       │
│       │       │       │
│       │       │       │
├───────┼───────┼───────┤
│       │       │       │
│       │       │       │
│       │       │       │
├───────┼───────┼───────┤
│       │       │       │
│       │       │       │
│       │       │       │
└───────┴───────┴───────┘
`,
		},
		{
			description: "Filled board",
			board: [boardSize][boardSize]uint8{
				{2, 4, 5, 9, 8, 1, 7, 0, 6},
				{8, 9, 3, 6, 2, 0, 1, 0, 4},
				{6, 1, 0, 3, 5, 4, 0, 9, 0},
				{5, 6, 0, 0, 9, 3, 8, 1, 7},
				{3, 2, 0, 4, 7, 0, 5, 6, 9},
				{0, 7, 8, 0, 1, 6, 4, 2, 0},
				{4, 0, 2, 0, 6, 9, 3, 7, 1},
				{1, 8, 0, 7, 0, 5, 9, 0, 2},
				{7, 0, 9, 0, 4, 2, 6, 8, 5},
			},
			expected: `┌───────┬───────┬───────┐
│ 2 4 5 │ 9 8 1 │ 7   6 │
│ 8 9 3 │ 6 2   │ 1   4 │
│ 6 1   │ 3 5 4 │   9   │
├───────┼───────┼───────┤
│ 5 6   │   9 3 │ 8 1 7 │
│ 3 2   │ 4 7   │ 5 6 9 │
│   7 8 │   1 6 │ 4 2   │
├───────┼───────┼───────┤
│ 4   2 │   6 9 │ 3 7 1 │
│ 1 8   │ 7   5 │ 9   2 │
│ 7   9 │   4 2 │ 6 8 5 │
└───────┴───────┴───────┘
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			s := NewSudokuWithBoard(tt.board)
			var sb strings.Builder
			s.PrintTo(&sb)
			actual := sb.String()
			if actual != tt.expected {
				t.Errorf("PrintTo() output mismatch:\nExpected:\n%s\nActual:\n%s", tt.expected, actual)
			}
		})
	}
}

func BenchmarkCreateSudoku(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := NewSudoku()
		s.GenerateFullBoard()
		s.CreatePuzzle(50)
	}
}
