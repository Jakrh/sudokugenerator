package sudoku

func (s *Sudoku) isUniqueSolution() bool {
	clonedSudoku := s.Clone()
	count := 0

	var solve func(s *Sudoku, row, col int) bool
	solve = func(s *Sudoku, row, col int) bool {
		if row == boardSize {
			count++
			return count > 1
		}
		if col == boardSize {
			return solve(s, row+1, 0)
		}
		if s.board[row][col] != 0 {
			return solve(s, row, col+1)
		}
		for num := minNum; num <= maxNum; num++ {
			if s.isValidCell(row, col, num) {
				s.board[row][col] = num
				if solve(s, row, col+1) {
					return true
				}
				s.board[row][col] = 0
			}
		}
		return false
	}

	solve(clonedSudoku, 0, 0)

	return count == 1
}

func (s *Sudoku) removeNumbers(target uint) {
	if target > maxRemovedNumbers {
		panic("number of cells to remove exceeds the maximum allowed")
	}

	// Generate a random permutation of cell positions to determine which cells to remove
	cells := s.randPerm(0, totalCells-1)
	idx := 0

	var removed uint
	for removed < target && idx < totalCells {
		position := cells[idx]
		row := position / boardSize
		col := position % boardSize
		if s.board[row][col] == 0 {
			idx++
			continue
		}
		cellBackup := s.board[row][col]
		s.board[row][col] = 0
		if !s.isUniqueSolution() {
			s.board[row][col] = cellBackup
			idx++
			continue
		}
		idx++
		removed++
	}
}

// Maximum removed amount of cells is 64
func (s *Sudoku) CreatePuzzle(removedPercentage uint) {
	if removedPercentage > 100 {
		panic("difficulty must be in 0-100")
	}

	// Get numbers from percentage
	numbers := removedPercentage * maxRemovedNumbers / 100
	s.removeNumbers(numbers)
}
