package main

import (
	"strconv"
)

func reverse(x int) int {
	if x < 0 {
		r := -reverse(-x)
		if r < -2147483648 {
			return 0
		}
		return r
	}
	s := strconv.Itoa(x)
	var r int
	o := 1
	for i := 0; i < len(s); i++ {
		r += int(s[i]-48) * o
		o *= 10
	}

	if 2147483647 < r {
		return 0
	}
	return r
}
