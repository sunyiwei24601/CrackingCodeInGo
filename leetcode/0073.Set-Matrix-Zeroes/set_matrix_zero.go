package _073_Set_Matrix_Zeroes

func SetMatrixZeroes(matrix [][]int) {
	// find first zero row
	setFirstRow := 0
	setFirstCol := 0
	for _, i := range matrix[0]{
		if matrix[0][i] == 0{
			setFirstRow = 1
		}
	}
	for i,_ := range matrix{
		if matrix[i][0] == 0{
			setFirstCol = 1
		}
	}

	for _, row := range matrix {
		for j, value := range row {
			if value == 0 {
				row[0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := range matrix[0] {
		if i == 0 {
			continue
		}
		if matrix[0][i] == 0 {
			setColumnZero(matrix, i)
		}
	}

	for i := range matrix {
		if i == 0 {
			continue
		}
		if matrix[i][0] == 0 {
			setZero(matrix[i])
		}
	}

	if setFirstRow == 1 {
		setZero(matrix[0])
	}
	if setFirstCol == 1 {
		setColumnZero(matrix, 0)
	}

}

func setZero(s []int) {
	for i := range s {
		s[i] = 0
	}
}

func setColumnZero(matrix [][]int, colNum int) {
	for _, row := range matrix {
		row[colNum] = 0
	}
}
