package p34

func searchRange(nums []int, target int) []int {
	min, max := -1, -1
	for i, j := 0, len(nums)-1; i <= j; {
		m := i + (j-i)/2
		if target > nums[m] {
			i = m + 1
		} else if target < nums[m] {
			j = m - 1
		} else {
			if min != -1 {
				if m < min {
					min = m
				}
			} else {
				min = m
			}
			j = m - 1
		}
	}
	for i, j := 0, len(nums)-1; i <= j; {
		m := i + (j-i)/2
		if target > nums[m] {
			i = m + 1
		} else if target < nums[m] {
			j = m - 1
		} else {
			if max != -1 {
				if m > max {
					max = m
				}
			} else {
				max = m
			}
			i = m + 1
		}
	}
	return []int{min, max}
}
