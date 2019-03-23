package p3

func lengthOfLongestSubstring(s string) int {
	start, end, max := 0, 0, 0
	uniq := map[rune](int){}

	for i, c := range s {
		dup_i, existed := uniq[c]
		if existed && dup_i >= start {
			start = dup_i + 1
		}
		uniq[c] = i
		end = i
		if end-start+1 > max {
			max = end - start + 1
		}
	}
	return max
}
