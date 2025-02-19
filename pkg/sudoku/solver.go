package sudoku

func (s *Sudoku) Solve() bool {
	// TODO: Consider to panic if this board does not have unique solution

	var solve func(row, col int) bool
	solve = func(row, col int) bool {
		if row == boardSize {
			return true
		}
		if col == boardSize {
			return solve(row+1, 0)
		}
		if s.board[row][col] != 0 {
			return solve(row, col+1)
		}
		for _, num := range s.randPerm(int(minNum), int(maxNum)) {
			num := uint8(num)
			if s.isValidCell(row, col, num) {
				s.board[row][col] = num
				if solve(row, col+1) {
					return true
				}
				s.board[row][col] = 0
			}
		}
		return false
	}

	return solve(0, 0)
}
