package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, n := range nums {
		numMap[n] = i
	}
	for i, n := range nums {
		j, ok := numMap[target-n]
		if ok && i != j {
			return []int{i, j}
		}
	}
	return []int{}
}
