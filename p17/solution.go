package p17

var lmap = [11]string{
	"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
}

func letterCombinations(digits string) []string {
	//assume digits contains only 1-9
	ret := []string{}
	if len(digits) == 0 {
		return ret
	}

	comb := lmap[digits[0]-'0']

	for i := 0; i < len(comb); i++ {
		current := string(comb[i])
		if len(digits) > 1 {
			for _, element := range letterCombinations(digits[1:]) {
				ret = append(ret, current+element)
			}
		} else {
			ret = append(ret, current)
		}

	}
	return ret
}
