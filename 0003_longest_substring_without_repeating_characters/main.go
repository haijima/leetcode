package main

func lengthOfLongestSubstring(s string) int {
	cs := make(map[int32]int)
	head := 0
	max := 0
	for i, c := range s {
		if j, ok := cs[c]; ok && j >= head {
			head = j + 1
		} else {
			if i-head+1 > max {
				max = i - head + 1
			}
		}
		cs[c] = i
	}
	return max
}
