package sol

func setZeroes(matrix [][]int) {
	firstColumnZero := false
	ROW, COL := len(matrix), len(matrix[0])
	// mark matrix[0][col] = 0 , matrix[row][0] = 0 if matrix[row][col] = 0
	for row := 0; row < ROW; row++ {
		for col := 0; col < COL; col++ {
			if matrix[row][col] == 0 {
				// mark row
				matrix[row][0] = 0
				// mark col
				if col == 0 {
					firstColumnZero = true
				} else {
					matrix[0][col] = 0
				}
			}
		}
	}
	// 1..ROW-1, 1..COL-1 , if matrix[row][0] = 0 || matrix[0][col] = 0, matrix[row][col] = 0
	for row := 1; row < ROW; row++ {
		for col := 1; col < COL; col++ {
			if matrix[row][0] == 0 || matrix[0][col] == 0 {
				if matrix[row][col] != 0 {
					matrix[row][col] = 0
				}
			}
		}
	}
	// check first row
	if matrix[0][0] == 0 {
		for col := 0; col < COL; col++ {
			if matrix[0][col] != 0 {
				matrix[0][col] = 0
			}
		}
	}
	if firstColumnZero {
		for row := 0; row < ROW; row++ {
			if matrix[row][0] != 0 {
				matrix[row][0] = 0
			}
		}
	}
}
