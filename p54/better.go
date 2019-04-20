// func spiralOrder(matrix [][]int) []int {
// 	n := len(matrix)
// 	if n < 1 {
// 		return nil
// 	}

// 	if n == 1 {
// 		return matrix[0]
// 	}

// 	return append(matrix[0], spiralOrder(transpose(matrix[1:]))...)

// }

// func transpose(mat [][]int) [][]int {
// 	n := len(mat)
// 	if n < 1 {
// 		return mat
// 	}
// 	m := len(mat[0])
// 	res := make([][]int, m)
// 	for i := 0; i < m; i++ {
// 		res[i] = make([]int, n)
// 		for j := 0; j < n; j++ {
// 			res[i][j] = mat[j][m-1-i]
// 		}
// 	}
// 	return res

// }