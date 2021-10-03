package leetcode

// rotate2DArray 初解
func rotate2DArray(matrix [][]int) {
	l := len(matrix)
	tmpArr := make([][]int, l)
	for k := 0; k < l; k++ {
		tmpArr[k] = make([]int, l)
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			tmpArr[j][l-1-i] = matrix[i][j]
		}
	}
	copy(matrix, tmpArr)
}

// rotate2DArrayV2 无额外内存消耗
func rotate2DArrayV2(matrix [][]int) {
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
