package sudoku

import (
	"fmt"
	"io"
	"math/rand/v2"
	"os"
	"strings"
)

const (
	blockSize               = 3
	boardSize               = blockSize * blockSize
	totalCells              = boardSize * boardSize
	minNum            uint8 = 1
	maxNum            uint8 = blockSize * blockSize
	maxRemovedNumbers uint  = 64 // Known 17 is minimal amount of cells of Sudoku board
)

type Sudoku struct {
	board [boardSize][boardSize]uint8
}

// Return Empty sudoku board
func NewSudoku() *Sudoku {
	return &Sudoku{}
}

func NewSudokuWithBoard(board [boardSize][boardSize]uint8) *Sudoku {
	return &Sudoku{board: board}
}

func (s *Sudoku) Clone() *Sudoku {
	ss := &Sudoku{}
	ss.board = s.board

	return ss
}

func (s *Sudoku) Clean() {
	s.board = [boardSize][boardSize]uint8{}
}

func (s *Sudoku) hasAnyEmptyCell() bool {
	for i := 0; i < len(s.board); i++ {
		for j := 0; j < len(s.board[i]); j++ {
			if s.board[i][j] == 0 {
				return true
			}
		}
	}

	return false
}

func (s *Sudoku) countZeros() int {
	count := 0
	for i := 0; i < len(s.board); i++ {
		for j := 0; j < len(s.board[i]); j++ {
			if s.board[i][j] == 0 {
				count++
			}
		}
	}

	return count
}

func (s *Sudoku) isValidSudoku() bool {
	usedRow := [boardSize][boardSize]bool{}
	usedCol := [boardSize][boardSize]bool{}
	usedBlock := [boardSize][boardSize]bool{}

	for i := 0; i < len(s.board); i++ {
		for j := 0; j < len(s.board[i]); j++ {
			if s.board[i][j] != 0 {
				num := s.board[i][j] - 1
				block := i/blockSize*blockSize + j/blockSize
				if usedRow[i][num] || usedCol[j][num] || usedBlock[block][num] {
					return false
				}
				usedRow[i][num], usedCol[j][num], usedBlock[block][num] = true, true, true
			}
		}
	}

	return true
}

func (s *Sudoku) isValidCell(row, col int, num uint8) bool {
	if num < minNum || num > maxNum {
		panic(fmt.Sprintf("our-of-range number: %d", num))
	}

	blockRow, blockCol := blockSize*(row/blockSize), blockSize*(col/blockSize) // get start cell of this block
	for i := 0; i < boardSize; i++ {
		if s.board[i][col] == num {
			return false
		}
		if s.board[row][i] == num {
			return false
		}
		if s.board[blockRow+i/blockSize][blockCol+i%blockSize] == num {
			return false
		}
	}

	return true
}

// Random permutation returns shuffled numbers from start to end (included)
// It panics if end number is less than start number.
// Usage: randPerm(1, 9) => [7 1 9 2 3 8 5 6 4]
func (s *Sudoku) randPerm(start, end int) []int {
	if start > end {
		panic(fmt.Sprintf("start number %d must not greater than end number %d", start, end))
	}

	// Generate sequential numbers
	nums := make([]int, end-start+1)
	num := start
	for i := 0; i < len(nums); i++ {
		nums[i] = num
		num++
	}

	// Fisher-Yates shuffle
	// ref: https://github.com/golang/go/blob/master/src/math/rand/v2/rand.go#L238-L247
	for i := len(nums) - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		nums[i], nums[j] = nums[j], nums[i]
	}

	return nums
}

func (s *Sudoku) PrintTo(out io.Writer) {
	fmt.Fprintln(out, "┌───────┬───────┬───────┐")
	for i := 0; i < len(s.board); i++ {
		fmt.Fprint(out, "│ ")
		for j := 0; j < len(s.board[i]); j++ {
			if s.board[i][j] != 0 {
				fmt.Fprintf(out, "%d ", s.board[i][j])
			} else {
				fmt.Fprint(out, "  ")
			}
			if (j+1)%blockSize == 0 && (j+1) != boardSize {
				fmt.Fprint(out, "│ ")
			}
		}
		fmt.Fprintln(out, "│")
		if (i+1)%blockSize == 0 && i+1 != boardSize {
			fmt.Fprintln(out, "├───────┼───────┼───────┤")
		}

	}
	fmt.Fprintln(out, "└───────┴───────┴───────┘")
}

func (s *Sudoku) PrintToString() string {
	var sb strings.Builder
	s.PrintTo(&sb)
	return sb.String()
}

func (s *Sudoku) Print() {
	s.PrintTo(os.Stdout)
}
