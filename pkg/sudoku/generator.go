package sudoku

func (s *Sudoku) fillBoard() {
	// TODO: Consider to fill cells with out-of-order => cells := s.randPerm(0, totalCells-1)
	cells := make([]int, totalCells)
	for i := 0; i < len(cells); i++ {
		cells[i] = i
	}

	var fill func(idx int) bool
	fill = func(idx int) bool {
		if idx == len(cells) {
			return true
		}

		nums := s.randPerm(int(minNum), int(maxNum))
		position := cells[idx]
		row := position / boardSize
		col := position % boardSize
		if s.board[row][col] != 0 {
			return fill(idx + 1)
		}

		for _, num := range nums {
			if s.isValidCell(row, col, uint8(num)) {
				s.board[row][col] = uint8(num)
				if fill(idx + 1) {
					return true
				}
				s.board[row][col] = 0
			}
		}
		return false
	}

	fill(0)
}

func (s *Sudoku) GenerateFullBoard() {
	// TODO: merge with s.Solve()
	s.fillBoard()
}
