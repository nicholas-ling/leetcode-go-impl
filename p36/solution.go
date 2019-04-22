package p36

func isValidSudoku(board [][]byte) bool {
	return !isRowDup(board) && !isColDup(board) && !isBoxDup(board)
}

func isRowDup(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		row := [9]byte{}
		for j := 0; j < 9; j++ {
			row[j] = board[i][j]
		}
		if isDup(row) {
			return true
		}
	}
	return false
}

func isColDup(board [][]byte) bool {
	for j := 0; j < 9; j++ {
		col := [9]byte{}
		for i := 0; i < 9; i++ {
			col[i] = board[i][j]
		}
		if isDup(col) {
			return true
		}
	}
	return false
}

func isBoxDup(board [][]byte) bool {
	for m := 0; m < 9; m += 3 {
		for n := 0; n < 9; n += 3 {
			box := [9]byte{}
			k := 0
			for i := m; i <= m+2; i++ {
				for j := n; j <= n+2; j++ {
					box[k] = board[i][j]
					k++
				}
			}
			if isDup(box) {
				return true
			}
		}
	}
	return false

}

func isDup(nums [9]byte) bool {
	set := [10]bool{}
	for _, val := range nums {
		if val != '.' {
			if set[val-'1'] {
				return true
			}
			set[val-'1'] = true
		}
	}
	return false
}
