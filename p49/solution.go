package p49

import "sort"

func groupAnagrams(strs []string) [][]string {

	ret := [][]string{}
	tmap := map[string][]int{}

	for index, str := range strs {
		key := sorts(str)
		value, ok := tmap[key]
		if ok {
			tmap[key] = append(value, index)
		} else {
			tmap[key] = []int{index}
		}
	}

	for _, indices := range tmap {
		anagram := make([]string, len(indices))
		for index, i := range indices {
			anagram[index] = strs[i]
		}
		ret = append(ret, anagram)
	}

	return ret

}

func sorts(s string) string {
	arr := []byte(s)
	sort.Slice(arr, func(i int, j int) bool { return arr[i] < arr[j] })
	return string(arr)
}
