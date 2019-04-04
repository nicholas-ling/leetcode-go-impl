package p22

func generateParenthesis(n int) []string {
	var ret = []string{}
	var bt func(int, int, string, int)
	bt = func(o int, c int, solution string, length int) {
		if o < 0 || c < 0 || o > c {
			return
		}
		if o == 0 && c == 0 && len(solution) == length {
			ret = append(ret, solution)
			return
		}
		bt(o-1, c, solution+"(", length)
		bt(o, c-1, solution+")", length)
	}
	bt(n, n, "", n*2)
	return ret
}
