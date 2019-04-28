package p56

import "sort"

func merge(intervals [][]int) [][]int {
	ret := [][]int{}
	if len(intervals) == 0 {
		return ret
	}

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals); i++ {
		if len(ret) == 0 {
			ret = append(ret, intervals[i])
		} else {
			if intervals[i][0] > ret[len(ret)-1][1] {
				ret = append(ret, intervals[i])
			} else {
				if ret[len(ret)-1][1] < intervals[i][1] {
					ret[len(ret)-1][1] = intervals[i][1]
				}
			}
		}
	}
	return ret
}
