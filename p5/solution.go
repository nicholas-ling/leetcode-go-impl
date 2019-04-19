package p5

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	matrix := make([][]bool, n)
	for i := range matrix {
		matrix[i] = make([]bool, n)
	}
	ret := s[0:1]

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= i; j-- {
			if j == i {
				matrix[i][j] = true
			} else if j == i+1 {
				if s[i] == s[j] {
					matrix[i][j] = true
					if j-i+1 > len(ret) {
						ret = s[i : j+1]
					}
				}
			} else {
				if matrix[i+1][j-1] {
					if s[i] == s[j] {
						matrix[i][j] = true
						if j-i+1 > len(ret) {
							ret = s[i : j+1]
						}
					}
				}
			}
		}
	}

	return ret

}
