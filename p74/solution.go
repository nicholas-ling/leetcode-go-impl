package p74

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	for i, j, mid := 0, m*n-1, 0; i <= j; {
		mid = (i + j) / 2
		if target > matrix[mid/n][mid%n] {
			i = mid + 1
		} else if target < matrix[mid/n][mid%n] {
			j = mid - 1
		} else {
			return true
		}
	}
	return false
}
