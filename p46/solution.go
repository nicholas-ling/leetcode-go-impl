package p46

var ret [][]int

func permute(nums []int) [][]int {
	ret = [][]int{}
	if len(nums) == 0 {
		return ret
	}
	rec(nums, []int{})
	return ret
}

func rec(left, ans []int) {
	if len(left) == 0 {
		ret = append(ret, ans)
		return
	}
	for i, v := range left {
		cleft := append([]int{}, left...)
		cans := append([]int{}, ans...)
		rec(append(cleft[:i], cleft[i+1:]...), append(cans, v))
	}
}
