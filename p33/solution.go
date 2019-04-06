package p33

func search(nums []int, target int) int {
	for m, i, j := 0, 0, len(nums)-1; i <= j; {
		if target == nums[i] {
			return i
		}
		if target == nums[j] {
			return j
		}
		m = (j-i)/2 + i
		if m == i || m == j {
			return -1
		}
		if nums[m] < nums[j] {
			if target > nums[m] && target < nums[j] {
				i = m
			} else {
				j = m
			}
		} else {
			if target > nums[i] && target < nums[m] {
				j = m
			} else {
				i = m
			}
		}
	}
	return -1
}
