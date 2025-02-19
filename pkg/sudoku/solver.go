package sudoku

func (s *Sudoku) Solve() bool {
	// TODO: Consider to panic if this board does not have unique solution

	// Generate all cell positions
	cells := make([]int, totalCells)
	for i := 0; i < totalCells; i++ {
		cells[i] = i
	}

	var solve func(idx int) bool
	solve = func(idx int) bool {
		if idx == totalCells {
			return true
		}

		position := cells[idx]
		row := position / boardSize
		col := position % boardSize
		if s.board[row][col] != 0 {
			return solve(idx + 1)
		}

		for _, num := range s.randPerm(int(minNum), int(maxNum)) {
			num := uint8(num)
			if s.isValidCell(row, col, num) {
				s.board[row][col] = num
				if solve(idx + 1) {
					return true
				}
				s.board[row][col] = 0
			}
		}
		return false
	}

	return solve(0)
}
