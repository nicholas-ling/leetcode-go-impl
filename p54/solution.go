package p54

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	length := len(matrix) * len(matrix[0])
	ret := make([]int, 0, length)
	recur(ret, matrix)
	return ret[0:length]

}

func recur(ret []int, matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	for _, val := range matrix[0] {
		ret = append(ret, val)
	}
	recur(ret, anticlockwise(matrix[1:]))
}

func anticlockwise(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return [][]int{}
	}
	m, n := len(matrix), len(matrix[0])
	ret := make([][]int, n, n)
	for index := range ret {
		ret[index] = make([]int, m, m)
	}

	for j := n - 1; j >= 0; j-- {
		for i := 0; i < m; i++ {
			ret[n-1-j][i] = matrix[i][j]
		}
	}
	return ret
}
